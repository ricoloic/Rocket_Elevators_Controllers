package elevator

import (
	"fmt"
	basic "go_controller/controller"
	"go_controller/controller/prints"
	"sort"
)

// Elevator is ...
type Elevator struct {
	ID                string
	floorAmount       int
	basementAmount    int
	Points            int
	door              string
	Status            string
	currentDirection  string
	previousDirection string
	CurrentFloor      int
	previousFloor     int
	StopList          []int
	upBuffer          []int
	downBuffer        []int
}

// StartElevator is used for setting some initial value to the elevator at its creation
func (e *Elevator) StartElevator(id string, floorAmount int, basementAmount int, minRange int) {
	e.Status = "IDLE"
	e.door = "Closed"
	e.CurrentFloor = minRange
	e.previousFloor = e.CurrentFloor
	e.floorAmount = floorAmount
	e.ID = id
	e.basementAmount = basementAmount
	e.currentDirection = "Stop"
	e.previousDirection = e.currentDirection
}

// All0Remove will delete all items that are equal to zero in the elevator stop lists
func (e *Elevator) All0Remove() {
	e.StopList = e.Remove0fromList(e.StopList)
	e.upBuffer = e.Remove0fromList(e.upBuffer)
	e.downBuffer = e.Remove0fromList(e.downBuffer)
}

// Remove0fromList will used for removing all element that are equal to zero in a given list
func (e *Elevator) Remove0fromList(nb []int) []int {
	awaitList := []int{}
	for i := 0; i < len(nb); i++ {
		if nb[i] != 0 {
			awaitList = append(awaitList, nb[i])
		}
	}
	return awaitList
}

// allCheck will return the value zero if the number provided is in one of the list else return the number provided
func (e *Elevator) allCheck(num int) int {
	if e.CheckIn(num) {
		num = 0
	}
	return num
}

// CheckIn will return the value zero if the provided value is in one of the the elevator stop lists else it will return the provided value
func (e *Elevator) CheckIn(n int) bool {
	inList := false

	for i := 0; i < len(e.StopList); i++ {
		if n == e.StopList[i] {
			inList = true
		}
	}

	for i := 0; i < len(e.upBuffer); i++ {
		if n == e.upBuffer[i] {
			inList = true
		}
	}

	for i := 0; i < len(e.downBuffer); i++ {
		if n == e.downBuffer[i] {
			inList = true
		}
	}

	if n == e.CurrentFloor {
		inList = true
	}

	return inList
}

// DoorState will change the state of the doors to open then close and also call the printing of the state
func (e *Elevator) DoorState() {
	prints.CreateArrival(e.CurrentFloor)

	prints.DoorOpen("1")
	e.door = "Open"
	prints.DoorClose("1")
	e.door = "Closed"
}

// ListSort will sort the stopList of the elevator based on the current direction of the elevator
func (e *Elevator) ListSort() {
	if e.currentDirection == "Down" {
		sort.Slice(e.StopList, func(i, j int) bool {
			return e.StopList[i] > e.StopList[j]
		})
	} else {
		sort.Slice(e.StopList, func(i, j int) bool {
			return e.StopList[i] < e.StopList[j]
		})
	}
}

// PointsUpdateFloor will give the elevator some points | this method is only used if the request was made from anywhere in the building apart from the ground floor.
// The less point the better!
// It will first set a variable for the difference between the current floor of the elevator and the "_floor".
// It will check if the elevator is IDLE or not and if not it is gonna set a new variable for the difference between the last index of the list of request and the "_floor".
// If elevator is going in the same direction and the "_floor" is in the path of the elevator / set point with the length of the stop list + the difference floor.
// if IDLE / set points to min range + the difference floor.
// if same direction not in the path / set point to max range + difference last stop + length of stop list.
// if not same direction / set point to max range * 2 + difference last stop + length of stop list.
func (e *Elevator) PointsUpdateFloor(_floor int, _direction string, _maxRange int, _minRange int) {
	differenceLastStop := 0
	differenceFloor := basic.Positive(e.CurrentFloor - _floor)
	e.Points = 0

	if e.Status != "IDLE" {
		differenceLastStop = basic.Positive(e.StopList[len(e.StopList)-1] - _floor)
	}

	conditionInPath := (_floor >= e.CurrentFloor && _direction == "Up") || (_floor <= e.CurrentFloor && _direction == "Down")
	conditionNotInPath := (_floor < e.CurrentFloor && _direction == "Up") || (_floor > e.CurrentFloor && _direction == "Down")

	if e.Status == "IDLE" {
		if _maxRange < 0 {
			e.Points = basic.Positive(_minRange) + 1 + differenceFloor
		} else {
			e.Points = basic.Positive(_maxRange) + 1 + differenceFloor
		}

	} else if e.currentDirection == _direction {
		if conditionInPath {
			e.Points = differenceFloor + len(e.StopList)

		} else if conditionNotInPath {
			e.Points = basic.Positive(_maxRange) + differenceLastStop + len(e.StopList)

		}

	} else if e.currentDirection != _direction {
		e.Points = basic.Positive(_maxRange)*2 + differenceLastStop + len(e.StopList)
	}
}

// PointsUpdateLobby will give the elevator some point | this method is only used if the request was made from the ground floor.
// The less point the better!
// it will first set a variable for the difference between the current floor of the elevator and the "_floor".
// it will check if the elevator is IDLE or not and if not it is gonna set a new variable for the difference between the last index of the list of request and the "_floor".
// if elevator is not going in the same direction as the user direction / set point with the difference floor + the difference last stop.
// if IDLE / set points to min range + the difference floor + 1.
// if Elevator is in the same direction as the user direction / set point to max range * 2 + difference last stop + length of stop list.
// if the current floor of the elevator is equal to the "_floor" / set point to the length of stop list.
func (e *Elevator) PointsUpdateLobby(_floor int, _direction string, _maxRange int, _minRange int) {
	differenceLastStop := 0
	differenceFloor := basic.Positive(e.CurrentFloor - _floor)
	e.Points = 0

	if e.Status != "IDLE" {
		differenceLastStop = basic.Positive(e.StopList[len(e.StopList)-1] - _floor)
	}

	if e.Status == "IDLE" {
		if _maxRange < 0 {
			e.Points = basic.Positive(_minRange) + 1 + differenceFloor
		} else {
			e.Points = basic.Positive(_maxRange) + 1 + differenceFloor
		}

	} else if _direction != e.currentDirection {
		e.Points = differenceLastStop + differenceFloor

	} else if e.currentDirection == _direction {
		e.Points = basic.Positive(_maxRange)*2 + len(e.StopList) + differenceLastStop
	}

	if e.CurrentFloor == _floor {
		e.Points = len(e.StopList)
	}
}

// AddStopLobby | the big idea here is to add both the stop of the user and the current floor of the user to the good stop list of the elevator "stopList, upBuffer, downBuffer"
func (e *Elevator) AddStopLobby(_floor int, _stop int, _direction string) {
	floor := e.allCheck(_floor)
	stop := e.allCheck(_stop)

	if _direction != e.currentDirection && _floor <= e.CurrentFloor {
		e.StopList = append(e.StopList, floor)
		e.upBuffer = append(e.upBuffer, stop)

	} else if _direction != e.currentDirection && _floor >= e.CurrentFloor {
		e.StopList = append(e.StopList, floor)
		e.downBuffer = append(e.downBuffer, stop)

	} else if e.Status == "IDLE" {
		e.StopList = append(e.StopList, floor)

		if _direction == "Up" {
			e.upBuffer = append(e.upBuffer, stop)

		} else if _direction == "Down" {
			e.downBuffer = append(e.downBuffer, stop)
		}

	} else if _direction == e.currentDirection {
		if _floor == e.CurrentFloor {
			e.StopList = append(e.StopList, floor)
			e.StopList = append(e.StopList, stop)

		} else if _floor != e.CurrentFloor {
			e.StopList = append(e.StopList, floor)

			if _direction == "Up" {
				e.upBuffer = append(e.upBuffer, stop)

			} else if _direction == "Down" {
				e.downBuffer = append(e.downBuffer, stop)
			}

		} else if _floor < e.CurrentFloor {
			e.downBuffer = append(e.downBuffer, floor)
			e.upBuffer = append(e.upBuffer, stop)

		} else if _floor > e.CurrentFloor {
			e.upBuffer = append(e.upBuffer, floor)
			e.downBuffer = append(e.downBuffer, stop)
		}
	}
}

// AddStopFloor | the big idea here is also to add both the stop of the user and the current floor of the user to the good stop list of the elevator "stopList, upBuffer, downBuffer"
func (e *Elevator) AddStopFloor(_floor int, _stop int, _direction string) {
	floor := e.allCheck(_floor)
	stop := e.allCheck(_stop)

	if e.Status == "IDLE" {
		e.StopList = append(e.StopList, floor)

		if _direction == "Up" {
			e.upBuffer = append(e.upBuffer, stop)

		} else if _direction == "Down" {
			e.downBuffer = append(e.downBuffer, stop)
		}

	} else if _direction == e.currentDirection {
		if (_direction == "Up" && _floor >= e.CurrentFloor) || (_direction == "Down" && _floor <= e.CurrentFloor) {
			e.StopList = append(e.StopList, floor)
			e.StopList = append(e.StopList, stop)

		} else if _direction == "Up" && _floor < e.CurrentFloor {
			e.downBuffer = append(e.downBuffer, floor)
			e.upBuffer = append(e.upBuffer, stop)

		} else if _direction == "Down" && _floor > e.CurrentFloor {
			e.upBuffer = append(e.upBuffer, floor)
			e.downBuffer = append(e.downBuffer, stop)
		}

	} else if _direction != e.currentDirection {
		if _direction == "Up" {
			e.upBuffer = append(e.upBuffer, floor)
			e.upBuffer = append(e.upBuffer, stop)

		} else if _direction == "Down" {
			e.downBuffer = append(e.downBuffer, floor)
			e.downBuffer = append(e.downBuffer, stop)
		}
	}
}

// stopSwitch will replace the stopList of the elevator with one of the buffer lists
func (e *Elevator) stopSwitch() {
	if len(e.upBuffer) != 0 && len(e.downBuffer) != 0 {
		if e.previousDirection == "Up" {
			e.StopList = e.downBuffer
			e.downBuffer = []int{}

		} else if e.previousDirection == "Down" {
			e.StopList = e.upBuffer
			e.upBuffer = []int{}
		}

	} else if len(e.upBuffer) == 0 && len(e.downBuffer) != 0 {
		e.StopList = e.downBuffer
		e.downBuffer = []int{}

	} else if len(e.upBuffer) != 0 && len(e.downBuffer) == 0 {
		e.StopList = e.upBuffer
		e.upBuffer = []int{}

	} else if len(e.upBuffer) == 0 && len(e.downBuffer) == 0 {
		e.Status = "IDLE"
		e.currentDirection = "Stop"
	}

	if len(e.StopList) != 0 {
		e.Run()
	}
}

// Run will move the elevator based on its currentFloor and the next stop in the stopList
func (e *Elevator) Run() {
	if e.currentDirection != e.previousDirection {
		e.previousDirection = e.currentDirection
	}

	for len(e.StopList) != 0 {
		for e.CurrentFloor != e.StopList[0] {
			e.Status = "MOVING"

			if e.StopList[0] < e.CurrentFloor {
				e.currentDirection = "Down"
				e.CurrentFloor--

			} else if e.StopList[0] > e.CurrentFloor {
				e.currentDirection = "Up"
				e.CurrentFloor++
			}
		}

		if e.CurrentFloor == e.StopList[0] && e.previousFloor != e.CurrentFloor {
			fmt.Println("")
			prints.CreateState(e.ID, e.previousFloor, e.Status, e.StopList[0])
			e.DoorState()
			e.previousFloor = e.StopList[0]
			e.StopList = basic.Remove(e.StopList, 0)

		} else if len(e.StopList) != 0 && e.previousFloor == e.CurrentFloor {
			e.StopList = basic.Remove(e.StopList, 0)
		}
	}

	if len(e.StopList) == 0 {
		e.stopSwitch()
	}
}

// ChangeValueE ...
func (e *Elevator) ChangeValueE(_status string, _currentFloor int, _stopList []int, _currentDirection string) {
	e.Status = _status
	e.CurrentFloor = _currentFloor
	e.previousFloor = e.CurrentFloor
	e.currentDirection = _currentDirection
	e.StopList = _stopList
	e.ListSort()
}
