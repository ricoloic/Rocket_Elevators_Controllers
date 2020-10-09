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
	ID               string
	basementAmount   int
	MaxRange         int
	MinRange         int
	floorPerColumn   int
	Status           string
	ElevatorList     []elevator.Elevator
	lobby            bool
	atFirstIteration bool
	bestOption       *elevator.Elevator
}

// StartColumn will call all the method needed for a column to be created
func (c *Column) StartColumn(floorAmount int, basementAmount int, elevatorColumn int, floorColumn int, i int, previousMax int, remainder int, id string) {
	c.initValue(id, basementAmount, floorColumn)
	c.setRange(i, previousMax, remainder)
	c.createElevatorList(i, elevatorColumn, floorAmount)
}

// initValue will be use to set the some initial value to the column (ID, basementAmount and floorPerColumn)
func (c *Column) initValue(id string, basementAmount int, floorColumn int) {
	c.ID = id
	c.basementAmount = basementAmount
	c.floorPerColumn = floorColumn
	c.Status = "ACTIVE"
}

// createElevatorList will add Elevator(s) to the column
func (c *Column) createElevatorList(i int, elevatorColumn int, floorAmount int) {
	for i := 0; i < elevatorColumn; i++ {
		elev := &elevator.Elevator{}
		elev.StartElevator(c.ID+strconv.Itoa(i+1), floorAmount, c.basementAmount, c.MinRange)
		c.ElevatorList = append(c.ElevatorList, *elev)
	}
}

// setRange will call the a method to make the min and max range of the column
func (c *Column) setRange(i int, previousMax int, remainder int) {
	if c.basementAmount == 0 {
		c.ifNoBasement(i, remainder, previousMax)
	} else {
		c.ifBasement(i, remainder, previousMax)
	}
}

// ifNoBasement set the ranges of the column for a battery without any basement
func (c *Column) ifNoBasement(i int, remainder int, previousMax int) {
	if i == 0 {
		c.MaxRange = c.floorPerColumn
		c.MinRange = 1

	} else {
		c.MaxRange = i*c.floorPerColumn + remainder
		c.MinRange = previousMax + 1
	}
}

// ifBasement set the ranges of the column for a battery with basement
func (c *Column) ifBasement(i int, remainder int, previousMax int) {
	if i == 0 {
		c.MaxRange = -1
		c.MinRange = -c.basementAmount

	} else if previousMax == -1 {
		c.MaxRange = i * c.floorPerColumn
		c.MinRange = previousMax + 2
	} else {
		c.MaxRange = i*c.floorPerColumn + remainder
		c.MinRange = previousMax + 1
	}
}

// InitializeRequest will call methods to make a request possible
func (c *Column) InitializeRequest(_floor int, _stop int, _direction string) {
	c.isAtLobby(_floor)
	c.updatePoints(_floor, _direction)
	c.sortByPoint()
	c.initPointsPrints()
	c.addStop(_floor, _stop, _direction)
	c.runAll()
}

// isAtLobby will set the boolean variable lobby to true if the request is made from the lobby
func (c *Column) isAtLobby(_floor int) {
	if _floor == 1 {
		c.lobby = true
	} else {
		c.lobby = false
	}
}

// updatePoints will call the method for updating the points of every elevator in the column
func (c *Column) updatePoints(_floor int, _direction string) {
	for i := 0; i < len(c.ElevatorList); i++ {
		if c.lobby {
			c.ElevatorList[i].PointsUpdateLobby(_floor, _direction, c.MaxRange, c.MinRange)

		} else {
			c.ElevatorList[i].PointsUpdateFloor(_floor, _direction, c.MaxRange, c.MinRange)
		}
	}
}

// sortByPoint will reorganize the list of elevators in such a way that the first index of the list is the elevator with the fewest points then set the variable bestOption to the fist index in the list of elevators
func (c *Column) sortByPoint() {
	sort.Slice(c.ElevatorList, func(i, j int) bool {
		return c.ElevatorList[i].Points < c.ElevatorList[j].Points
	})

	c.bestOption = &c.ElevatorList[0]
}

// initPointsPrints will create two list, one for the elevators id and one for there points then call the printing of the points for each elevator
func (c *Column) initPointsPrints() {
	var IDs []string
	var points []string

	for i := 0; i < len(c.ElevatorList); i++ {
		IDs = append(IDs, c.ElevatorList[i].ID)
		points = append(points, strconv.Itoa(c.ElevatorList[i].Points))
	}

	prints.CreatePointing(c.ID, IDs, points)
}

// addStop will call the method of the bestOption to add the request to its list
func (c *Column) addStop(_floor int, _stop int, _direction string) {
	if !c.lobby {
		c.bestOption.AddStopFloor(_floor, _stop, _direction)
	} else {
		c.bestOption.AddStopLobby(_floor, _stop, _direction)
	}
}

// runAll will call all the methods needed to make all the elevator to move
func (c *Column) runAll() {
	for i := 0; i < len(c.ElevatorList); i++ {
		c.ElevatorList[i].All0Remove()
		c.ElevatorList[i].ListSort()
		c.ElevatorList[i].Run()
		fmt.Println()
	}
}

// ChangeValueC ...
func (c *Column) ChangeValueC(_elevator int, _status string, _currentFloor int, _stopList []int, _currentDirection string) {
	c.ElevatorList[_elevator].ChangeValueE(_status, _currentFloor, _stopList, _currentDirection)
}
