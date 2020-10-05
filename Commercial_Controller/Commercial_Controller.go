package main

import (
	"fmt"
	"sort"
	// "sort"
)

//--------------------------------------------------------//

type Battery struct {
	columnBattery  int
	floorAmount    int
	basementAmount int
	elevatorColumn int
	floorPerColumn int
	columnList     []Column
}

func (b Battery) startBattery(_columnBattery int, _floorAmount int, _basementAmount int, _elevatorColumn int) {
	b.columnBattery = _columnBattery
	b.floorAmount = _floorAmount
	b.basementAmount = _basementAmount
	b.elevatorColumn = _elevatorColumn

	previousMax := 0

	if _basementAmount == 0 {
		b.floorPerColumn = b.floorAmount / b.columnBattery

	} else {
		b.floorPerColumn = b.floorAmount / (b.columnBattery - 1)
	}

	for i := 0; i < b.columnBattery; i++ {
		if i != 0 {
			previousMax = b.columnList[i-1].maxRange
		}

		col := &Column{}
		col.startColumn(b.floorAmount, b.basementAmount, b.elevatorColumn, b.floorPerColumn, i, previousMax)
		b.columnList = append(b.columnList, *col)
	}
}

func (b Battery) columnSelection(_floor int, _stop int, _direction string) {
	if _stop == b.basementAmount+1 {
		for i := 0; i < len(b.columnList); i++ {
			if _floor >= b.columnList[i].minRange && _floor <= b.columnList[i].maxRange {
				b.columnList[i].request(_floor, _stop, _direction)
			}
		}

	} else {
		for i := 0; i < len(b.columnList); i++ {
			if _stop >= b.columnList[i].minRange && _stop <= b.columnList[i].maxRange {
				b.columnList[i].request(_floor, _stop, _direction)
			}
		}
	}
}

func (b Battery) changeValue(column int, _elevator int, _stopList []int, _status string, _currentFloor int, _currentDirection string) {
	b.columnList[column].changeValue(_elevator, _stopList, _status, _currentFloor, _currentDirection)
}

//--------------------------------------------------------//

type Column struct {
	id             int
	basementAmount int
	maxRange       int
	minRange       int
	status         string
	elevatorList   []Elevator
}

func (c Column) startColumn(_floorAmount int, _basementAmount int, _elevatorColumn int, _floorColumn int, _iteration int, _previousMax int) {
	c.id = _iteration + 1
	c.basementAmount = _basementAmount

	if c.basementAmount != 0 {
		if _iteration == 0 {
			c.maxRange = c.basementAmount
			c.minRange = 1

		} else {
			c.maxRange = _iteration*_floorColumn + c.basementAmount + 1
			c.minRange = _previousMax + 1
		}

	} else {
		if _iteration == 0 {
			c.maxRange = _floorColumn
			c.minRange = 1

		} else {
			c.maxRange = _iteration*_floorColumn + 1
			c.minRange = _previousMax + 1
		}
	}

	for i := 0; i < _elevatorColumn; i++ {
		elev := &Elevator{}
		elev.startElevator(i+1, _floorAmount, c.basementAmount, c.minRange)
		c.elevatorList = append(c.elevatorList, *elev)
	}
}

func (c Column) request(_floor int, _stop int, _direction string) {
	for i := 0; i < len(c.elevatorList); i++ {
		if _floor == c.basementAmount+1 {
			c.elevatorList[i].pointsUpdateLobby(_floor, _direction, c.maxRange)

		} else {
			c.elevatorList[i].pointsUpdateFloor(_floor, _direction, c.maxRange)
		}
	}

	c.elevatorToSend(_floor, _stop, _direction)
}

func (c Column) elevatorToSend(_floor int, _stop int, _direction string) {
	sort.Slice(c.elevatorList, func(i, j int) bool {
		return c.elevatorList[i].points < c.elevatorList[j].points
	})

	bestOption := c.elevatorList[0]
	bestOption.addStop(_floor, _stop, _direction)
	c.runAll()
}

func (c Column) runAll() {
	for i := 0; i < len(c.elevatorList); i++ {
		c.elevatorList[i].run()
	}
}

func (c Column) changeValue(_elevator int, _stopList []int, _status string, _currentFloor int, _currentDirection string) {
	c.elevatorList[_elevator].changeValue(_stopList, _status, _currentFloor, _currentDirection)
}

//--------------------------------------------------------//

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

func (e Elevator) startElevator(_id int, _floorAmount int, _basementAmount int, _minRange int) {
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

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func (e Elevator) positive(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func (e Elevator) doorState() {
	e.door = "Open"
	fmt.Println(e.door)

	e.door = "Closed"
	fmt.Println(e.door)
}

func (e Elevator) listSort() {
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

func (e Elevator) pointsUpdateFloor(_floor int, _direction string, _maxRange int) {
	expretion := (_floor >= e.currentFloor && _direction == "Up") || (_floor <= e.currentFloor && _direction == "Down")
	differenceLastStop := 0

	if e.status != "IDLE" {
		differenceLastStop = e.positive(e.stopList[len(e.stopList)-1] - _floor)
	}

	differenceFloor := e.positive(e.currentFloor - _floor)

	e.points = 0

	if e.status == "IDLE" {
		e.points = _maxRange + differenceFloor + 1

	} else if e.currentDirection == _direction {
		if expretion {
			e.points = differenceFloor + len(e.stopList)

		} else if _floor < e.currentFloor && _direction == "Up" || _floor > e.currentFloor && _direction == "Down" {
			e.points = _maxRange + differenceLastStop + len(e.stopList)
		}
	} else if e.currentDirection != _direction {
		e.points = _maxRange*2 + differenceLastStop + len(e.stopList)
	}
}

func (e Elevator) pointsUpdateLobby(_floor int, _direction string, _maxRange int) {
	differenceLastStop := 0

	if e.status != "IDLE" {
		differenceLastStop = e.positive(e.stopList[len(e.stopList)-1] - _floor)
	}

	differenceFloor := e.positive(e.currentFloor - _floor)
	e.points = 0

	if e.status == "IDLE" {
		e.points = _maxRange + differenceLastStop + 1

	} else if _direction != e.currentDirection {
		e.points = differenceLastStop + differenceFloor

	} else if e.currentDirection == _direction {
		e.points = _maxRange*2 + len(e.stopList) + differenceLastStop
	}

	if e.currentFloor == _floor {
		e.points = len(e.stopList)
	}
}

func (e Elevator) addStop(_floor int, _stop int, _direction string) {
	if _floor == e.basementAmount+1 {
		if _direction != e.currentDirection && _floor <= e.currentFloor {
			e.stopList = append(e.stopList, _floor)
			e.upBuffer = append(e.upBuffer, _stop)

		} else if _direction != e.currentDirection && _floor >= e.currentFloor {
			e.stopList = append(e.stopList, _floor)
			e.upBuffer = append(e.downBuffer, _stop)

		} else if e.status == "IDLE" {
			e.stopList = append(e.stopList, _floor)

			if _direction == "up" {
				e.upBuffer = append(e.upBuffer, _stop)

			} else if _direction == "down" {
				e.downBuffer = append(e.downBuffer, _stop)
			}

		} else if _direction == e.currentDirection {
			if _floor == e.currentFloor {
				e.stopList = append(e.stopList, _floor)
				e.upBuffer = append(e.stopList, _stop)

			} else if _floor != e.currentFloor {
				if _direction == "up" {
					e.stopList = append(e.stopList, _floor)
					e.upBuffer = append(e.upBuffer, _stop)

				} else if _direction == "up" {
					e.stopList = append(e.stopList, _floor)
					e.downBuffer = append(e.downBuffer, _stop)
				}

			} else if _floor < e.currentFloor {
				e.downBuffer = append(e.downBuffer, _floor)
				e.upBuffer = append(e.upBuffer, _stop)

			} else if _floor > e.currentFloor {
				e.upBuffer = append(e.upBuffer, _floor)
				e.downBuffer = append(e.downBuffer, _stop)
			}
		}
	} else {
		if e.status == "IDLE" {
			e.stopList = append(e.stopList, _floor)

			if _direction == "up" {
				e.upBuffer = append(e.upBuffer, _stop)

			} else if _direction == "down" {
				e.downBuffer = append(e.downBuffer, _stop)
			}

		} else if _direction == e.currentDirection {
			if _direction == "up" && _floor >= e.currentFloor {
				e.stopList = append(e.stopList, _stop)
				e.stopList = append(e.stopList, _stop)

			} else if _direction == "down" && _floor <= e.currentFloor {
				e.stopList = append(e.stopList, _stop)
				e.stopList = append(e.stopList, _stop)

			} else if _direction == "up" && _floor < e.currentFloor {
				e.downBuffer = append(e.downBuffer, _floor)
				e.upBuffer = append(e.upBuffer, _stop)

			} else if _direction == "down" && _floor > e.currentFloor {
				e.upBuffer = append(e.upBuffer, _floor)
				e.downBuffer = append(e.downBuffer, _stop)
			}

		} else if _direction != e.currentDirection {
			if _direction == "up" {
				e.upBuffer = append(e.upBuffer, _floor)
				e.upBuffer = append(e.upBuffer, _stop)

			} else if _direction == "down" {
				e.downBuffer = append(e.downBuffer, _floor)
				e.downBuffer = append(e.downBuffer, _stop)
			}
		}
	}

	e.listSort()
}

func (e Elevator) stopSwitch() {
	if len(e.upBuffer) != 0 && len(e.downBuffer) != 0 {
		if e.previousDirection == "Up" {
			e.stopList = append(e.stopList, e.downBuffer...)
			for i := 0; i < len(e.downBuffer); i++ {
				e.downBuffer = remove(e.downBuffer, 0)
			}

		} else if e.previousDirection == "Down" {
			e.stopList = append(e.stopList, e.upBuffer...)
			for i := 0; i < len(e.upBuffer); i++ {
				e.upBuffer = remove(e.upBuffer, 0)
			}
		}

	} else if len(e.upBuffer) == 0 && len(e.downBuffer) != 0 {
		e.stopList = append(e.stopList, e.downBuffer...)
		for i := 0; i < len(e.downBuffer); i++ {
			e.downBuffer = remove(e.downBuffer, 0)
		}

	} else if len(e.upBuffer) != 0 && len(e.downBuffer) == 0 {
		e.stopList = append(e.stopList, e.upBuffer...)
		for i := 0; i < len(e.upBuffer); i++ {
			e.upBuffer = remove(e.upBuffer, 0)
		}

	} else if len(e.upBuffer) == 0 && len(e.downBuffer) == 0 {
		e.status = "IDLE"
		e.currentDirection = "Stop"
	}
}

func (e Elevator) run() {
	if len(e.stopList) != 0 {
		for len(e.stopList) != 0 {
			for len(e.stopList) != 0 {
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

				if e.previousFloor != e.currentFloor {
					e.previousFloor = e.currentFloor
				}
			}

			if e.stopList[0] == e.currentFloor {
				e.doorState()
				e.stopList = remove(e.stopList, 0)
			}
		}
	}

	if len(e.stopList) == 0 {
		e.stopSwitch()
	}
}

func (e Elevator) changeValue(_stopList []int, _status string, _currentFloor int, _currentDirection string) {
	e.stopList = append(e.stopList, _stopList...)
	e.status = _status
	e.currentFloor = _currentFloor
	e.currentDirection = _currentDirection
}

//-----------------------------------------------------------------------------//

//_elevator int, _stopList []int, _status string, _currentFloor int, _currentDirection string

func main() {

}
