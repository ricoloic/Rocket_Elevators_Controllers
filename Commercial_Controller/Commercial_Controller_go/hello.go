package main

import (
	"fmt"
	"sort"
)

func main() {
	scenarioPreset(1)
}

//--------------------------------------------------------//

// Battery is ...
type Battery struct {
	Printer
	columnBattery  int
	floorAmount    int
	basementAmount int
	elevatorColumn int
	floorPerColumn int
	columnList     []Column
	p              *Printer
}

func (b *Battery) startBattery(_columnBattery int, _floorAmount int, _basementAmount int, _elevatorColumn int) {
	b.p = &Printer{}
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
				b.p.columnSelectedLine(b.columnList[i])
				b.columnList[i].request(_floor, _stop, _direction)
			}
		}

	} else {
		for i := 0; i < len(b.columnList); i++ {
			if _stop >= b.columnList[i].minRange && _stop <= b.columnList[i].maxRange {
				b.p.columnSelectedLine(b.columnList[i])
				b.columnList[i].request(_floor, _stop, _direction)
			}
		}
	}
}

func (b *Battery) changeValueB(_column int, _elevator int, _status string, _currentFloor int, _stopList []int, _currentDirection string) {
	b.columnList[_column].changeValueC(_elevator, _status, _currentFloor, _stopList, _currentDirection)
}

//--------------------------------------------------------//

// Column is ...
type Column struct {
	Printer
	id             int
	basementAmount int
	maxRange       int
	minRange       int
	status         string
	elevatorList   []Elevator
	p              *Printer
}

func (c *Column) startColumn(_floorAmount int, _basementAmount int, _elevatorColumn int, _floorColumn int, _iteration int, _previousMax int) {
	c.p = &Printer{}
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

		} else {
			c.elevatorList[i].pointsUpdateFloor(_floor, _direction, c.maxRange, c.minRange)
		}
	}

	c.elevatorToSend(_floor, _stop, _direction)
}

func (c *Column) elevatorToSend(_floor int, _stop int, _direction string) {
	sort.Slice(c.elevatorList, func(i, j int) bool {
		return c.elevatorList[i].points < c.elevatorList[j].points
	})

	points := []int{}
	IDs := []int{}

	for i := 0; i < len(c.elevatorList); i++ {
		IDs = append(IDs, c.elevatorList[i].id)
		points = append(points, c.elevatorList[i].points)
	}

	c.p.createPointing(IDs, points)
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

func (c *Column) changeValueC(_elevator int, _status string, _currentFloor int, _stopList []int, _currentDirection string) {
	c.elevatorList[_elevator].changeValueE(_status, _currentFloor, _stopList, _currentDirection)
}

//--------------------------------------------------------//

// Elevator is ...
type Elevator struct {
	Printer
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
	p                 *Printer
}

func (e *Elevator) startElevator(_id int, _floorAmount int, _basementAmount int, _minRange int) {
	e.p = &Printer{}
	e.status = "IDLE"
	e.door = "Closed"
	e.currentFloor = _minRange
	e.previousFloor = e.currentFloor
	e.floorAmount = _floorAmount
	e.id = _id
	e.basementAmount = _basementAmount
	e.currentDirection = "Stop"
	e.previousDirection = e.currentDirection
}

func (e *Elevator) remove0fromList(_nb []int) []int {
	nb := _nb
	awaitList := []int{}

	for i := 0; i < len(nb); i++ {
		if nb[i] != 0 {
			awaitList = append(awaitList, nb[i])
		}
	}
	nb = awaitList
	return nb
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
	e.p.createArrival(e.currentFloor)

	e.door = "Open"
	e.p.doorOpen("1")

	e.door = "Closed"
	e.p.doorClose("1")
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
	e.stopList = e.remove0fromList(e.stopList)
	fmt.Println(e.stopList)
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
					e.previousFloor = e.currentFloor

				} else if e.stopList[0] > e.currentFloor {
					e.currentDirection = "Up"
					e.previousDirection = e.currentDirection
					e.currentFloor++
					e.previousFloor = e.currentFloor

				}

				if e.previousFloor != e.currentFloor && e.stopList[0] != e.currentFloor && e.currentFloor != 0 {
					switch e.previousDirection {
					case "Up":
						e.p.upArrow("1")
					case "Down":
						e.p.upArrow("1")
					}
				} else if e.currentFloor
			}

			if e.stopList[0] == e.currentFloor && e.currentFloor != 0 {
				if e.previousFloor != 0 {
					e.doorState()
					e.previousFloor = 0
				}
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

func (e *Elevator) changeValueE(_status string, _currentFloor int, _stopList []int, _currentDirection string) {
	e.status = _status
	e.currentFloor = _currentFloor
	e.previousFloor = _currentFloor
	e.currentDirection = _currentDirection
	e.stopList = _stopList
	e.listSort()
}

//-----------------------------------------------------------------------------//

func scenarioPreset(n int) {
	battery := &Battery{}
	p := &Printer{}
	battery.startBattery(4, 60, 6, 5)

	// ||=========> BASEMENT COLUMN #1 <=========||
	battery.changeValueB(0, 0, "IDLE", -4, []int{}, "Stop")
	battery.changeValueB(0, 1, "IDLE", 1, []int{}, "Stop")
	battery.changeValueB(0, 2, "MOVING", -3, []int{-5}, "Down")
	battery.changeValueB(0, 3, "MOVING", -6, []int{1}, "Up")
	battery.changeValueB(0, 4, "MOVING", -1, []int{-6}, "Down")

	// ||=========> FLOOR COLUMN #2 <=========||
	battery.changeValueB(1, 0, "MOVING", 20, []int{5}, "Down")
	battery.changeValueB(1, 1, "MOVING", 3, []int{15}, "Up")
	battery.changeValueB(1, 2, "MOVING", 13, []int{1}, "Down")
	battery.changeValueB(1, 3, "MOVING", 15, []int{2}, "Down")
	battery.changeValueB(1, 4, "MOVING", 6, []int{1}, "Down")

	// ||=========> FLOOR COLUMN #3 <=========||
	battery.changeValueB(2, 0, "MOVING", 1, []int{21}, "Up")
	battery.changeValueB(2, 1, "MOVING", 23, []int{28}, "Up")
	battery.changeValueB(2, 2, "MOVING", 33, []int{1}, "Down")
	battery.changeValueB(2, 3, "MOVING", 40, []int{24}, "Down")
	battery.changeValueB(2, 4, "MOVING", 39, []int{1}, "Down")

	// ||=========> FLOOR COLUMN #4 <=========||
	battery.changeValueB(3, 0, "MOVING", 58, []int{1}, "Down")
	battery.changeValueB(3, 1, "MOVING", 50, []int{60}, "Up")
	battery.changeValueB(3, 2, "MOVING", 46, []int{58}, "Up")
	battery.changeValueB(3, 3, "MOVING", 1, []int{54}, "Up")
	battery.changeValueB(3, 4, "MOVING", 60, []int{1}, "Down")

	f := []int{}
	d := ""

	if n == 1 {
		f = []int{1, 20}
	} else if n == 2 {
		f = []int{1, 36}
	} else if n == 3 {
		f = []int{54, 1}
	} else if n == 4 {
		f = []int{-3, 1}
	}

	if f[0] < f[1] {
		d = "Up"
	} else if f[0] > f[1] {
		d = "Down"
	}

	p.createRequest(f[0], f[1], d)
	battery.columnSelection(f[0], f[1], d)
}

//-----------------------------------------------------------------------------//

// Printer ...
type Printer struct {
	id             int
	status         string
	floor          int
	nextStop       int
	atFloor        int
	floorRequested int
	direction      string
}

func (p *Printer) positive(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func (p *Printer) createState(_id int, _floor int, _status string, _stop int) {
	p.topBottomLine(_status)
	p.innerArrowLine(_status)
	p.emptyDoubleLine(_status)
	p.idLine(_id, _status)
	p.floorLine(_floor, _status)
	p.statusLine(_status)
	p.stopLine(_stop, _status)
	p.emptyDoubleLine(_status)
	p.innerArrowLine(_status)
	p.topBottomLine(_status)
}

func (p *Printer) createArrival(_floor int) {
	size := p.countStr(_floor)

	//p.elevatorLine(_id, _floor)
	p.topBottomLine(size)
	p.floorArivedLine(_floor)
	p.topBottomLine(size)
}

func (p *Printer) createRequest(_floor int, _stop int, _direction string) {
	p.topBottomLine("2")
	p.innerArrowAndRequestLine()
	p.emptyDoubleLine("2")
	p.atFloorLine(_floor)
	p.floorRequestLine(_stop)
	p.directionLine(_direction)
	p.emptyDoubleLine("2")
	p.innerArrowLine("2")
	p.topBottomLine("2")
}

func (p *Printer) createPointing(_id []int, _points []int) {
	fmt.Println("")

	// count := p.countInt(_id[0]) + p.countInt(_points[0])

	for i := 0; i < len(_points); i++ {
		fmt.Println("		      ELEVATOR", _id[i], "- HAS", _points[i], "pts")
		fmt.Println("")
	}

	p.bestElevatorLine(_id[0], _points[0])
}

func (p *Printer) bestElevatorLine(_id int, _points int) {
	fmt.Println("		THE BEST ELEVATOR IS ELEVATOR", _id, "- WITH", _points, "pts")
}

func (p *Printer) columnSelectedLine(c Column) {
	fmt.Println("		THE SELECTED COLUMN IS COLUMN", c.id)
}

func (p *Printer) countStr(n int) string {
	count := "0"

	if n >= 0 && n < 10 {
		count = "1"
	} else if n < 100 && n > 9 || n < 0 && n > -10 {
		count = "2"
	} else if n < 1000 && n > 99 || n < -10 && n > -100 {
		count = "3"
	} else if n < 10000 && n > 999 || n < -100 && n > -1000 {
		count = "4"
	}

	return count
}

func (p *Printer) countInt(n int) int {
	count := 0

	if n >= 0 && n < 10 {
		count = 1
	} else if n < 100 && n > 9 || n < 0 && n > -10 {
		count = 2
	} else if n < 1000 && n > 99 || n < -10 && n > -100 {
		count = 3
	} else if n < 10000 && n > 999 || n < -100 && n > -1000 {
		count = 4
	}

	return count
}

func (p *Printer) topBottomLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		+------------------------------+")
	} else if _size == "MOVING" {
		fmt.Println("		+--------------------------------+")
	} else if _size == "1" {
		fmt.Println("		+-----------------------------------+")
	} else if _size == "2" {
		fmt.Println("		+------------------------------------+")
	} else if _size == "3" || _size == "MAINTENANCE" {
		fmt.Println("		+-------------------------------------+")
	} else if _size == "4" {
		fmt.Println("		+--------------------------------------+")
	}
}

func (p *Printer) innerArrowLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		| +--->                  <---+ |")
	} else if _size == "MOVING" {
		fmt.Println("		| +--->                    <---+ |")
	} else if _size == "1" {
		fmt.Println("		| +--->                       <---+ |")
	} else if _size == "2" {
		fmt.Println("		| +--->                        <---+ |")
	} else if _size == "3" || _size == "MAINTENANCE" {
		fmt.Println("		| +--->                         <---+ |")
	} else if _size == "4" {
		fmt.Println("		| +--->                          <---+ |")
	}
}

func (p *Printer) elevatorLine(_id int) {
	count := p.countStr(_id)

	if count == "1" {
		fmt.Println("		  +--->      ELEVATOR", _id, "      <---+  ")
	} else if count == "2" {
		fmt.Println("		  +--->      ELEVATOR", _id, "     <---+  ")
	} else if count == "3" {
		fmt.Println("		  +--->     ELEVATOR", _id, "     <---+  ")
	} else if count == "4" {
		fmt.Println("		  +--->     ELEVATOR", _id, "    <---+  ")
	}
}

func (p *Printer) innerArrowAndRequestLine() {
	fmt.Println("		| +--->         REQUEST        <---+ |") // "2"
}

func (p *Printer) emptyDoubleLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		| |                          | |")
	} else if _size == "MOVING" {
		fmt.Println("		| |                            | |")
	} else if _size == "2" {
		fmt.Println("		| |                                | |")
	} else if _size == "MAINTENANCE" || _size == "3" {
		fmt.Println("		| |                                 | |")
	} else if _size == "4" {
		fmt.Println("		| |                                  | |")
	}
}

func (p *Printer) upArrow(_size string) {
	if _size == "1" {
		fmt.Println("		+---------------------+")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		| |        -        | |")
		fmt.Println("		| ▼      -/-\\-      ▼ |")
		fmt.Println("		|       /-/-\\-\\       |")
		fmt.Println("		|          |          |")
		fmt.Println("		| ▲       | |       ▲ |")
		fmt.Println("		| |                 | |")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		+---------------------+")
	} else if _size == "2" {
		fmt.Println("		+-------------------------+")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| ▼          -          ▼ |")
		fmt.Println("		|          -/-\\-          |")
		fmt.Println("		|        -/-/-\\-\\-        |")
		fmt.Println("		|           | |           |")
		fmt.Println("		|          | - |          |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		+-------------------------+")
	} else if _size == "3" {
		fmt.Println("		+----------------------------+")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| ▼            -           ▼ |")
		fmt.Println("		|            -/-\\-           |")
		fmt.Println("		|          -/-/-\\-\\-         |")
		fmt.Println("		|         /-/-/-\\-\\-\\        |")
		fmt.Println("		|            |   |           |")
		fmt.Println("		|             | |            |")
		fmt.Println("		| ▲          | - |         ▲ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		+----------------------------+")
	}
}

func (p *Printer) downArrow(_size string) {
	if _size == "1" {
		fmt.Println("		+---------------------+")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		| |                 | |")
		fmt.Println("		| ▼       | |       ▼ |")
		fmt.Println("		|          |          |")
		fmt.Println("		|       \\-\\-/-/       |")
		fmt.Println("		| ▲      -\\-/-      ▲ |")
		fmt.Println("		| |        -        | |")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		+---------------------+")
	} else if _size == "2" {
		fmt.Println("		+-------------------------+")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| ▼                     ▼ |")
		fmt.Println("		|          | - |          |")
		fmt.Println("		|           | |           |")
		fmt.Println("		|        -\\-\\-/-/-        |")
		fmt.Println("		|          -\\-/-          |")
		fmt.Println("		| ▲          -          ▲ |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		+-------------------------+")
	} else if _size == "3" {
		fmt.Println("		+----------------------------+")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| ▼          | - |         ▼ |")
		fmt.Println("		|             | |            |")
		fmt.Println("		|            |   |           |")
		fmt.Println("		|         \\-\\-\\-/-/-/        |")
		fmt.Println("		|          -\\-\\-/-/-         |")
		fmt.Println("		|            -\\-/-           |")
		fmt.Println("		| ▲            -           ▲ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		+----------------------------+")
	}
}

func (p *Printer) doorOpen(_size string) {
	if _size == "1" || _size == "2" {
		p.doorTopBottomLine(_size)
		p.doorMiddleLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
		}
		p.leftArrowLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
		}
		p.doorMiddleLine(_size)
		p.openLine(_size)
		p.doorMiddleLine(_size)
		p.doorMiddleLine(_size)
		p.leftArrowLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
		}
		p.doorMiddleLine(_size)
		p.doorTopBottomLine(_size)
	} else {
		fmt.Println("\n		The size entered for the door opening ain't good change it and try again !")
	}
}

func (p *Printer) doorClose(_size string) {
	if _size == "1" || _size == "2" {
		p.doorTopBottomLine(_size)
		p.doorMiddleLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
		}
		p.rightArrowLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
		}
		p.doorMiddleLine(_size)
		p.closingLine(_size)
		p.doorMiddleLine(_size)
		p.doorMiddleLine(_size)
		p.rightArrowLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
		}
		p.doorMiddleLine(_size)
		p.doorTopBottomLine(_size)
	} else {
		fmt.Println("\n		The size entered for the door opening ain't good change it and try again !")
	}
}

func (p *Printer) leftArrowLine(_size string) {
	if _size == "1" {
		fmt.Println("		|  <<  <<  ||   |")
	} else if _size == "2" {
		fmt.Println("		|   <<   <<   <<   ||      |")
	}
}

func (p *Printer) rightArrowLine(_size string) {
	if _size == "1" {
		fmt.Println("		|  >>  >>  ||   |")
	} else if _size == "2" {
		fmt.Println("		|   >>   >>   >>   ||      |")
	}
}

func (p *Printer) openLine(_size string) {
	if _size == "1" {
		fmt.Println("		|   Open   ||   |")
	} else if _size == "2" {
		fmt.Println("		|     Opening      ||      |")
	}
}

func (p *Printer) doorMiddleLine(_size string) {
	if _size == "1" {
		fmt.Println("		|   	   ||   |")
	} else if _size == "2" {
		fmt.Println("		|     	           ||      |")
	}
}

func (p *Printer) doorTopBottomLine(_size string) {
	if _size == "1" {
		fmt.Println("		+----------++---+")
	} else if _size == "2" {
		fmt.Println("		+------------------++------+")
	}
}

func (p *Printer) closingLine(_size string) {
	if _size == "1" {
		fmt.Println("		|   Close  ||   |")
	} else if _size == "2" {
		fmt.Println("		|     Closing      ||      |")
	}
}

func (p *Printer) floorArivedLine(_floor int) {
	count := p.countStr(_floor)

	if count == "1" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	} else if count == "2" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	} else if count == "3" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	} else if count == "4" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	}
}

func (p *Printer) idLine(_id int, _status string) {
	count := p.countStr(_id)

	if _status == "IDLE" {
		if count == "1" {
			fmt.Println("		| |       ID:", _id, "             | |")
		} else if count == "2" {
			fmt.Println("		| |       ID:", _id, "            | |")
		} else if count == "3" {
			fmt.Println("		| |       ID:", _id, "           | |")
		} else if count == "4" {
			fmt.Println("		| |       ID:", _id, "          | |")
		}
	} else if _status == "MOVING" {
		if count == "1" {
			fmt.Println("		| |       ID:", _id, "               | |")
		} else if count == "2" {
			fmt.Println("		| |       ID:", _id, "              | |")
		} else if count == "3" {
			fmt.Println("		| |       ID:", _id, "             | |")
		} else if count == "4" {
			fmt.Println("		| |       ID:", _id, "            | |")
		}
	} else if _status == "MAINTENANCE" {
		if count == "1" {
			fmt.Println("		| |       ID:", _id, "                    | |")
		} else if count == "2" {
			fmt.Println("		| |       ID:", _id, "                   | |")
		} else if count == "3" {
			fmt.Println("		| |       ID:", _id, "                  | |")
		} else if count == "4" {
			fmt.Println("		| |       ID:", _id, "                 | |")
		}
	}
}

func (p *Printer) directionLine(_direction string) {
	if _direction == "Up" {
		fmt.Println("		| |        DIRECTION:", "UP", "          | |")
	} else if _direction == "Down" {
		fmt.Println("		| |        DIRECTION:", "DOWN", "        | |")
	} else if _direction == "Stop" {
		fmt.Println("		| |        DIRECTION:", "STOP", "        | |")
	}
}

func (p *Printer) atFloorLine(_atFloor int) {
	count := p.countStr(_atFloor)

	if count == "1" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "            | |")
	} else if count == "2" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "           | |")
	} else if count == "3" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "          | |")
	} else if count == "4" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "         | |")
	}
}

func (p *Printer) floorRequestLine(_requestFloor int) {
	count := p.countStr(_requestFloor)

	if count == "1" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "     | |")
	} else if count == "2" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "    | |")
	} else if count == "3" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "   | |")
	} else if count == "4" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "  | |")
	}
}

func (p *Printer) floorLine(_floor int, _status string) {
	count := p.countStr(_floor)

	if _status == "IDLE" {
		if count == "1" {
			fmt.Println("		| |       Floor:", _floor, "          | |")
		} else if count == "2" {
			fmt.Println("		| |       Floor:", _floor, "         | |")
		} else if count == "3" {
			fmt.Println("		| |       Floor:", _floor, "        | |")
		} else if count == "4" {
			fmt.Println("		| |       Floor:", _floor, "       | |")
		}
	} else if _status == "MOVING" {
		if count == "1" {
			fmt.Println("		| |       Floor:", _floor, "            | |")
		} else if count == "2" {
			fmt.Println("		| |       Floor:", _floor, "           | |")
		} else if count == "3" {
			fmt.Println("		| |       Floor:", _floor, "          | |")
		} else if count == "4" {
			fmt.Println("		| |       Floor:", _floor, "         | |")
		}
	} else if _status == "MAINTENANCE" {
		if count == "1" {
			fmt.Println("		| |       Floor:", _floor, "                 | |")
		} else if count == "2" {
			fmt.Println("		| |       Floor:", _floor, "                | |")
		} else if count == "3" {
			fmt.Println("		| |       Floor:", _floor, "               | |")
		} else if count == "4" {
			fmt.Println("		| |       Floor:", _floor, "              | |")
		}
	}
}

func (p *Printer) stopLine(_stop int, _status string) {
	count := p.countStr(_stop)

	if _status == "IDLE" {
		if count == "1" {
			fmt.Println("		| |       Next-Stop:", _stop, "      | |")
		} else if count == "2" {
			fmt.Println("		| |       Next-Stop:", _stop, "     | |")
		} else if count == "3" {
			fmt.Println("		| |       Next-Stop:", _stop, "    | |")
		} else if count == "4" {
			fmt.Println("		| |       Next-Stop:", _stop, "   | |")
		}
	} else if _status == "MOVING" {
		if count == "1" {
			fmt.Println("		| |       Next-Stop:", _stop, "        | |")
		} else if count == "2" {
			fmt.Println("		| |       Next-Stop:", _stop, "       | |")
		} else if count == "3" {
			fmt.Println("		| |       Next-Stop:", _stop, "      | |")
		} else if count == "4" {
			fmt.Println("		| |       Next-Stop:", _stop, "     | |")
		}
	} else if _status == "MAINTENANCE" {
		if count == "1" {
			fmt.Println("		| |       Next-Stop:", _stop, "             | |")
		} else if count == "2" {
			fmt.Println("		| |       Next-Stop:", _stop, "            | |")
		} else if count == "3" {
			fmt.Println("		| |       Next-Stop:", _stop, "           | |")
		} else if count == "4" {
			fmt.Println("		| |       Next-Stop:", _stop, "          | |")
		}
	}
}

func (p *Printer) statusLine(_status string) {
	if _status == "IDLE" {
		fmt.Println("		| |       Status:", "IDLE", "      | |")
	} else if _status == "MOVING" {
		fmt.Println("		| |       Status:", "MOVING", "      | |")
	} else if _status == "MAINTENANCE" {
		fmt.Println("		| |       Status:", "MAINTENANCE", "      | |")
	}
}
