package column

import (
	"fmt"
	"go_controller/controller/battery/column/elevator"
	"go_controller/controller/prints"
	"sort"
	"strconv"
)

// Column is ...
type Column struct {
	ID             string
	basementAmount int
	MaxRange       int
	MinRange       int
	status         string
	elevatorList   []elevator.Elevator
}

// StartColumn ...
func (c *Column) StartColumn(_floorAmount int, _basementAmount int, _elevatorColumn int, _floorColumn int, _i int, _previousMax int, _remainder int, _letters []string) {
	c.ID = _letters[_i]
	c.basementAmount = _basementAmount

	if c.basementAmount == 0 {
		if _i == 0 {
			c.MaxRange = _floorColumn
			c.MinRange = 1

		} else {
			c.MaxRange = _i*_floorColumn + _remainder
			c.MinRange = _previousMax + 1
		}

	} else {
		if _i == 0 {
			c.MaxRange = -1
			c.MinRange = -c.basementAmount

		} else if _previousMax == -1 {
			c.MaxRange = _i * _floorColumn
			c.MinRange = _previousMax + 2
		} else {
			c.MaxRange = _i*_floorColumn + _remainder
			c.MinRange = _previousMax + 1
		}
	}

	for i := 0; i < _elevatorColumn; i++ {
		elev := &elevator.Elevator{}
		elev.StartElevator(c.ID+strconv.Itoa(i+1), _floorAmount, c.basementAmount, c.MinRange)
		c.elevatorList = append(c.elevatorList, *elev)
	}
}

func (c *Column) addStop(_floor int, _stop int, _direction string, n int) {
	if n == 1 {
		c.elevatorList[0].AddStopFloor(_floor, _stop, _direction)
	} else {
		c.elevatorList[0].AddStopLobby(_floor, _stop, _direction)
	}

	c.elevatorList[0].All0Remove()
	c.elevatorList[0].ListSort()
}

func (c *Column) elevatorToSend() {
	sort.Slice(c.elevatorList, func(i, j int) bool {
		return c.elevatorList[i].Points < c.elevatorList[j].Points
	})

	points := []string{}
	IDs := []string{}

	for i := 0; i < len(c.elevatorList); i++ {
		IDs = append(IDs, c.elevatorList[i].ID)
		points = append(points, strconv.Itoa(c.elevatorList[i].Points))
	}

	prints.CreatePointing(c.ID, IDs, points)
}

func (c *Column) runAll() {
	for i := 0; i < len(c.elevatorList); i++ {
		c.elevatorList[i].Run()
		fmt.Println()
	}
}

// Request ...
func (c *Column) Request(_floor int, _stop int, _direction string) {
	n := 1

	if _floor == 1 {
		for i := 0; i < len(c.elevatorList); i++ {
			c.elevatorList[i].PointsUpdateLobby(_floor, _direction, c.MaxRange, c.MinRange)
		}
		n = 0

	} else {
		for i := 0; i < len(c.elevatorList); i++ {
			c.elevatorList[i].PointsUpdateFloor(_floor, _direction, c.MaxRange, c.MinRange)
		}
	}

	c.requestToElev(_floor, _stop, _direction, n)
}

func (c *Column) requestToElev(_floor int, _stop int, _direction string, n int) {
	c.elevatorToSend()
	c.addStop(_floor, _stop, _direction, n)
	c.runAll()
}

// ChangeValueC ...
func (c *Column) ChangeValueC(_elevator int, _status string, _currentFloor int, _stopList []int, _currentDirection string) {
	c.elevatorList[_elevator].ChangeValueE(_status, _currentFloor, _stopList, _currentDirection)
}
