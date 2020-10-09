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

// StartElevator ...
func (e *Elevator) StartElevator(_id string, _floorAmount int, _basementAmount int, _minRange int) {
	e.Status = "IDLE"
	e.door = "Closed"
	e.CurrentFloor = _minRange
	e.previousFloor = e.CurrentFloor
	e.floorAmount = _floorAmount
	e.ID = _id
	e.basementAmount = _basementAmount
	e.currentDirection = "Stop"
	e.previousDirection = e.currentDirection
}

// Remove0fromList ...
func (e *Elevator) Remove0fromList(_nb []int) []int {
	awaitList := []int{}
	for i := 0; i < len(_nb); i++ {
		if _nb[i] != 0 {
			awaitList = append(awaitList, _nb[i])
		}
	}
	return awaitList
}

// All0Remove ...
func (e *Elevator) All0Remove() {
	e.StopList = e.Remove0fromList(e.StopList)
	e.upBuffer = e.Remove0fromList(e.upBuffer)
	e.downBuffer = e.Remove0fromList(e.downBuffer)
}

// CheckIn ...
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

// Remove ...
func (e *Elevator) Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// DoorState ...
func (e *Elevator) DoorState() {
	prints.CreateArrival(e.CurrentFloor)

	e.door = "Open"
	basic.Wait(2)
	prints.DoorOpen("1")

	basic.Wait(5)
	e.door = "Closed"
	prints.DoorClose("1")
}

// ListSort ...
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

// PointsUpdateFloor ...
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

// PointsUpdateLobby ...
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

// AllCheck ...
func (e *Elevator) AllCheck(num int) int {
	if e.CheckIn(num) {
		num = 0
	}
	return num
}

// AddStopLobby ...
func (e *Elevator) AddStopLobby(_floor int, _stop int, _direction string) {
	floor := e.AllCheck(_floor)
	stop := e.AllCheck(_stop)

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

// AddStopFloor ...
func (e *Elevator) AddStopFloor(_floor int, _stop int, _direction string) {
	floor := e.AllCheck(_floor)
	stop := e.AllCheck(_stop)

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

// StopSwitch ...
func (e *Elevator) StopSwitch() {
	if len(e.upBuffer) != 0 && len(e.downBuffer) != 0 {
		if e.previousDirection == "Up" {
			e.StopList = e.downBuffer
			for i := 0; i < len(e.downBuffer); i++ {
				e.downBuffer = e.Remove(e.downBuffer, 0)
			}

		} else if e.previousDirection == "Down" {
			e.StopList = e.upBuffer
			for i := 0; i < len(e.upBuffer); i++ {
				e.upBuffer = e.Remove(e.upBuffer, 0)
			}
		}

	} else if len(e.upBuffer) == 0 && len(e.downBuffer) != 0 {
		e.StopList = e.downBuffer
		for i := 0; i < len(e.downBuffer); i++ {
			e.downBuffer = e.Remove(e.downBuffer, 0)
		}

	} else if len(e.upBuffer) != 0 && len(e.downBuffer) == 0 {
		e.StopList = e.upBuffer
		for i := 0; i < len(e.upBuffer); i++ {
			e.upBuffer = e.Remove(e.upBuffer, 0)
		}

	} else if len(e.upBuffer) == 0 && len(e.downBuffer) == 0 {
		e.Status = "IDLE"
		e.currentDirection = "Stop"
	}

	if len(e.StopList) != 0 {
		e.Run()
	}
}

// Run ...
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
			if len(e.StopList) > 1 {
				fmt.Println("")
				prints.CreateState(e.ID, e.previousFloor, e.Status, e.StopList[0])
				e.DoorState()
				e.previousFloor = e.StopList[0]
				e.StopList = e.Remove(e.StopList, 0)

			} else {
				fmt.Println("")
				prints.CreateState(e.ID, e.previousFloor, e.Status, e.StopList[0])
				e.DoorState()
				e.previousFloor = e.StopList[0]
				e.StopList = e.Remove(e.StopList, 0)
			}

		} else if len(e.StopList) != 0 && e.previousFloor == e.CurrentFloor {
			e.StopList = e.Remove(e.StopList, 0)
		}
	}

	if len(e.StopList) == 0 {
		e.StopSwitch()
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
