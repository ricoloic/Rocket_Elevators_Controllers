package main

import (
	"fmt"
	"sort"
)

func main() {
	scenario(2)
}

//--------------------------------------------------------//

// Battery is ...
type Battery struct {
	columnBattery  int
	floorAmount    int
	basementAmount int
	elevatorColumn int
	floorPerColumn int
	columnList     []Column
}

func (b *Battery) startBattery(_columnBattery int, _floorAmount int, _basementAmount int, _elevatorColumn int) {
	b.columnBattery = _columnBattery
	b.floorAmount = _floorAmount
	b.basementAmount = _basementAmount
	b.elevatorColumn = _elevatorColumn

	previousMax := -b.basementAmount

	if _basementAmount == 0 {
		b.floorPerColumn = b.floorAmount / b.columnBattery

	} else {
		b.floorPerColumn = b.floorAmount / (b.columnBattery - 1)
	}

	for i := 0; i < b.columnBattery; i++ {
		if i > 0 {
			previousMax = b.columnList[i-1].maxRange
		}

		col := &Column{}
		col.startColumn(b.floorAmount, b.basementAmount, b.elevatorColumn, b.floorPerColumn, i, previousMax)
		b.columnList = append(b.columnList, *col)
	}
}

func (b *Battery) columnSelection(_floor int, _stop int, _direction string) {
	if _stop == 1 {
		for i := 0; i < len(b.columnList); i++ {
			if _floor >= b.columnList[i].minRange && _floor <= b.columnList[i].maxRange {
				fmt.Println("The column selected is :", b.columnList[i].id)
				b.columnList[i].request(_floor, _stop, _direction)
			}
		}

	} else {
		for i := 0; i < len(b.columnList); i++ {
			if _stop >= b.columnList[i].minRange && _stop <= b.columnList[i].maxRange {
				fmt.Println("The column selected is :", b.columnList[i].id)
				b.columnList[i].request(_floor, _stop, _direction)
			}
		}
	}
}

func (b *Battery) changeValueB(_column int, _elevator int, _stopList []int, _status string, _currentFloor int, _currentDirection string) {
	b.columnList[_column].changeValueC(_elevator, _stopList, _status, _currentFloor, _currentDirection)
}

//--------------------------------------------------------//

// Column is ...
type Column struct {
	id             int
	basementAmount int
	maxRange       int
	minRange       int
	status         string
	elevatorList   []Elevator
}

func (c *Column) startColumn(_floorAmount int, _basementAmount int, _elevatorColumn int, _floorColumn int, _iteration int, _previousMax int) {
	c.id = _iteration + 1
	c.basementAmount = _basementAmount

	if c.basementAmount == 0 {
		if _iteration == 0 {
			c.maxRange = _floorColumn
			c.minRange = 1

		} else {
			c.maxRange = _iteration * _floorColumn
			c.minRange = _previousMax + 1
		}

	} else {
		if _iteration == 0 {
			c.maxRange = -1
			c.minRange = -c.basementAmount

		} else if _previousMax == -1 {
			c.maxRange = _iteration * _floorColumn
			c.minRange = _previousMax + 2
		} else {
			c.maxRange = _iteration * _floorColumn
			c.minRange = _previousMax + 1
		}
	}

	for i := 0; i < _elevatorColumn; i++ {
		elev := &Elevator{}
		elev.startElevator(i+1, _floorAmount, c.basementAmount, c.minRange)
		c.elevatorList = append(c.elevatorList, *elev)
	}
}

func (c *Column) request(_floor int, _stop int, _direction string) {
	for i := 0; i < len(c.elevatorList); i++ {
		if _floor == 1 {
			c.elevatorList[i].pointsUpdateLobby(_floor, _direction, c.maxRange, c.minRange)
			//fmt.Println("Elevator", c.elevatorList[i].id, "has", c.elevatorList[i].points, "Points")

		} else {
			c.elevatorList[i].pointsUpdateFloor(_floor, _direction, c.maxRange, c.minRange)
			//fmt.Println("Elevator", c.elevatorList[i].id, "has", c.elevatorList[i].points, "Points")
		}
	}

	c.elevatorToSend(_floor, _stop, _direction)
}

func (c *Column) elevatorToSend(_floor int, _stop int, _direction string) {
	sort.Slice(c.elevatorList, func(i, j int) bool {
		return c.elevatorList[i].points < c.elevatorList[j].points
	})

	fmt.Println("The elevator selected is :", c.elevatorList[0].id, "with", c.elevatorList[0].points, "Points")
	c.elevatorList[0].addStop(_floor, _stop, _direction)
	c.elevatorList[0].run()
	c.runAll()
}

func (c *Column) runAll() {
	for i := 0; i < len(c.elevatorList); i++ {
		c.elevatorList[i].run()
		fmt.Println()
	}
}

func (c *Column) changeValueC(_elevator int, _stopList []int, _status string, _currentFloor int, _currentDirection string) {
	c.elevatorList[_elevator].changeValueE(_stopList, _status, _currentFloor, _currentDirection)
}

//--------------------------------------------------------//

// Elevator is ...
type Elevator struct {
	id                int
	floorAmount       int
	basementAmount    int
	points            int
	door              string
	status            string
	currentDirection  string
	previousDirection string
	currentFloor      int
	previousFloor     int
	stopList          []int
	upBuffer          []int
	downBuffer        []int
}

func (e *Elevator) startElevator(_id int, _floorAmount int, _basementAmount int, _minRange int) {
	e.status = "IDLE"
	e.door = "Closed"
	e.currentFloor = _minRange
	e.previousFloor = e.currentFloor
	e.floorAmount = _floorAmount
	e.id = _id
	e.points = 0
	e.basementAmount = _basementAmount
	e.currentDirection = "Stop"
	e.previousDirection = e.currentDirection
}

func (e *Elevator) checkIn(n int) bool {
	inList := false

	for i := 0; i < len(e.stopList); i++ {
		if n == e.stopList[i] {
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

	return inList
}

func (e *Elevator) remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func (e *Elevator) positive(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func (e *Elevator) doorState() {
	fmt.Println("The Elevator", e.id, "has arrived at Floor :", e.currentFloor)

	e.door = "Open"
	fmt.Println(e.door)

	e.door = "Closed"
	fmt.Println(e.door)
}

func (e *Elevator) listSort() {
	if e.currentDirection == "Down" {
		sort.Slice(e.stopList, func(i, j int) bool {
			return e.stopList[i] > e.stopList[j]
		})
	} else {
		sort.Slice(e.stopList, func(i, j int) bool {
			return e.stopList[i] < e.stopList[j]
		})
	}
}

func (e *Elevator) pointsUpdateFloor(_floor int, _direction string, _maxRange int, _minRange int) {
	differenceLastStop := 0

	if e.status != "IDLE" {
		differenceLastStop = e.positive(e.stopList[len(e.stopList)-1] - _floor)
	}

	differenceFloor := e.positive(e.currentFloor - _floor)

	e.points = 0

	if e.status == "IDLE" {
		if _maxRange < 0 {
			e.points = e.positive(_minRange) + 1 + differenceFloor
		} else {
			e.points = e.positive(_maxRange) + 1 + differenceFloor
		}

	} else if e.currentDirection == _direction {
		if _floor >= e.currentFloor && _direction == "Up" {
			e.points = differenceFloor + len(e.stopList)

		} else if _floor <= e.currentFloor && _direction == "Down" {
			e.points = differenceFloor + len(e.stopList)

		} else if _floor < e.currentFloor && _direction == "Up" {
			e.points = e.positive(_maxRange) + differenceLastStop + len(e.stopList)

		} else if _floor > e.currentFloor && _direction == "Down" {
			e.points = e.positive(_maxRange) + differenceLastStop + len(e.stopList)
		}

	} else if e.currentDirection != _direction {
		e.points = e.positive(_maxRange)*2 + differenceLastStop + len(e.stopList)
	}
}

func (e *Elevator) pointsUpdateLobby(_floor int, _direction string, _maxRange int, _minRange int) {
	differenceLastStop := 0

	if e.status != "IDLE" {
		differenceLastStop = e.positive(e.stopList[len(e.stopList)-1] - _floor)
	}

	differenceFloor := e.positive(e.currentFloor - _floor)
	e.points = 0

	if e.status == "IDLE" {
		if _maxRange < 0 {
			e.points = e.positive(_minRange) + 1 + differenceFloor
		} else {
			e.points = e.positive(_maxRange) + 1 + differenceFloor
		}

	} else if _direction != e.currentDirection {
		e.points = differenceLastStop + differenceFloor

	} else if e.currentDirection == _direction {
		e.points = e.positive(_maxRange)*2 + len(e.stopList) + differenceLastStop
	}

	if e.currentFloor == _floor {
		e.points = len(e.stopList)
	}
}

func (e *Elevator) addStop(_floor int, _stop int, _direction string) {
	floor := _floor
	stop := _stop

	if e.checkIn(_floor) {
		floor = 0
	}

	if e.checkIn(_stop) {
		stop = 0
	}

	if _floor == 1 {
		if _direction != e.currentDirection && _floor <= e.currentFloor {
			e.stopList = append(e.stopList, floor)
			e.upBuffer = append(e.upBuffer, stop)

		} else if _direction != e.currentDirection && _floor >= e.currentFloor {
			e.stopList = append(e.stopList, floor)
			e.downBuffer = append(e.downBuffer, stop)

		} else if e.status == "IDLE" {
			e.stopList = append(e.stopList, floor)

			if _direction == "Up" {
				e.upBuffer = append(e.upBuffer, stop)

			} else if _direction == "Down" {
				e.downBuffer = append(e.downBuffer, stop)
			}

		} else if _direction == e.currentDirection {
			if _floor == e.currentFloor {
				e.stopList = append(e.stopList, floor)
				e.upBuffer = append(e.stopList, stop)

			} else if _floor != e.currentFloor {
				if _direction == "Up" {
					e.stopList = append(e.stopList, floor)
					e.upBuffer = append(e.upBuffer, stop)

				} else if _direction == "Up" {
					e.stopList = append(e.stopList, floor)
					e.downBuffer = append(e.downBuffer, stop)
				}

			} else if _floor < e.currentFloor {
				e.downBuffer = append(e.downBuffer, floor)
				e.upBuffer = append(e.upBuffer, stop)

			} else if _floor > e.currentFloor {
				e.upBuffer = append(e.upBuffer, floor)
				e.downBuffer = append(e.downBuffer, stop)
			}
		}
	} else {
		if e.status == "IDLE" {
			e.stopList = append(e.stopList, floor)

			if _direction == "Up" {
				e.upBuffer = append(e.upBuffer, stop)

			} else if _direction == "Down" {
				e.downBuffer = append(e.downBuffer, stop)
			}

		} else if _direction == e.currentDirection {
			if _direction == "Up" && _floor >= e.currentFloor {
				e.stopList = append(e.stopList, floor)
				e.stopList = append(e.stopList, stop)

			} else if _direction == "Down" && _floor <= e.currentFloor {
				e.stopList = append(e.stopList, floor)
				e.stopList = append(e.stopList, stop)

			} else if _direction == "Up" && _floor < e.currentFloor {
				e.downBuffer = append(e.downBuffer, floor)
				e.upBuffer = append(e.upBuffer, stop)

			} else if _direction == "Down" && _floor > e.currentFloor {
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
	e.listSort()
}

func (e *Elevator) stopSwitch() {
	if len(e.upBuffer) != 0 && len(e.downBuffer) != 0 {
		if e.previousDirection == "Up" {
			e.stopList = e.downBuffer
			for i := 0; i < len(e.downBuffer); i++ {
				e.downBuffer = e.remove(e.downBuffer, 0)
			}

		} else if e.previousDirection == "Down" {
			e.stopList = e.upBuffer
			for i := 0; i < len(e.upBuffer); i++ {
				e.upBuffer = e.remove(e.upBuffer, 0)
			}
		}

	} else if len(e.upBuffer) == 0 && len(e.downBuffer) != 0 {
		e.stopList = e.downBuffer
		for i := 0; i < len(e.downBuffer); i++ {
			e.downBuffer = e.remove(e.downBuffer, 0)
		}

	} else if len(e.upBuffer) != 0 && len(e.downBuffer) == 0 {
		e.stopList = e.upBuffer
		for i := 0; i < len(e.upBuffer); i++ {
			e.upBuffer = e.remove(e.upBuffer, 0)
		}

	} else if len(e.upBuffer) == 0 && len(e.downBuffer) == 0 {
		e.status = "IDLE"
		e.currentDirection = "Stop"
	}
}

func (e *Elevator) run() {
	for len(e.stopList) != 0 {
		if len(e.stopList) != 0 {
			for e.currentFloor != e.stopList[0] {
				e.status = "MOVING"

				if e.stopList[0] < e.currentFloor {
					e.currentDirection = "Down"
					e.previousDirection = e.currentDirection
					e.currentFloor--

				} else if e.stopList[0] > e.currentFloor {
					e.currentDirection = "Up"
					e.previousDirection = e.currentDirection
					e.currentFloor++
				}

				if e.previousFloor != e.currentFloor && e.stopList[0] != e.currentFloor && e.currentFloor != 0 {
					fmt.Println("Elevator :", e.id, "- Floor :", e.currentFloor)
					e.previousFloor = e.currentFloor
				}
			}

			if e.stopList[0] == e.currentFloor && e.currentFloor != 0 {
				e.doorState()
				e.stopList = e.remove(e.stopList, 0)
			} else if e.stopList[0] == e.currentFloor && e.currentFloor == 0 {
				e.stopList = e.remove(e.stopList, 0)
			}
		}

		if len(e.stopList) == 0 {
			e.stopSwitch()
		}
	}
	if len(e.stopList) == 0 {
		e.stopSwitch()
	}
}

func (e *Elevator) changeValueE(_stopList []int, _status string, _currentFloor int, _currentDirection string) {
	e.status = _status
	e.currentFloor = _currentFloor
	e.currentDirection = _currentDirection
	e.stopList = _stopList
	e.listSort()
}

//-----------------------------------------------------------------------------//

func scenario(n int) {
	battery := &Battery{}
	battery.startBattery(4, 60, 6, 5)

	if n == 1 {
		battery.changeValueB(1, 0, []int{5}, "MOVING", 20, "Down")
		battery.changeValueB(1, 1, []int{15}, "MOVING", 3, "Up")
		battery.changeValueB(1, 2, []int{1}, "MOVING", 13, "Down")
		battery.changeValueB(1, 3, []int{2}, "MOVING", 15, "Down")
		battery.changeValueB(1, 4, []int{1}, "MOVING", 6, "Down")

		battery.columnSelection(1, 20, "Up")

	} else if n == 2 {
		battery.changeValueB(2, 0, []int{21}, "MOVING", 1, "Up")
		battery.changeValueB(2, 1, []int{28}, "MOVING", 23, "Up")
		battery.changeValueB(2, 2, []int{1}, "MOVING", 33, "Down")
		battery.changeValueB(2, 3, []int{24}, "MOVING", 40, "Down")
		battery.changeValueB(2, 4, []int{1}, "MOVING", 39, "Down")

		battery.columnSelection(1, 36, "Up")

	} else if n == 3 {
		battery.changeValueB(3, 0, []int{1}, "MOVING", 58, "Down")
		battery.changeValueB(3, 1, []int{60}, "MOVING", 50, "Up")
		battery.changeValueB(3, 2, []int{58}, "MOVING", 46, "Up")
		battery.changeValueB(3, 3, []int{54}, "MOVING", 1, "Up")
		battery.changeValueB(3, 4, []int{1}, "MOVING", 60, "Down")

		battery.columnSelection(54, 1, "Down")

	} else if n == 4 {
		battery.changeValueB(0, 0, []int{}, "IDLE", -4, "Stop")
		battery.changeValueB(0, 1, []int{}, "IDLE", 1, "Stop")
		battery.changeValueB(0, 2, []int{-5}, "MOVING", -3, "Down")
		battery.changeValueB(0, 3, []int{1}, "MOVING", -6, "Up")
		battery.changeValueB(0, 4, []int{-6}, "MOVING", -1, "Down")

		battery.columnSelection(-3, 1, "Up")
	}
}
