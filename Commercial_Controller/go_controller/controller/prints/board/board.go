package board

import (
	"fmt"
	"go_controller/controller/battery"
	"go_controller/controller/prints"
	"strconv"
)

// Board ...
type Board struct {
	bat *BatteryPrint
}

// BatteryPrint ...
type BatteryPrint struct {
	ID      string
	ColList []ColumnPrint
}

// ColumnPrint ...
type ColumnPrint struct {
	ID       string
	Status   string
	MaxRange string
	MinRange string
	ElevList []ElevatorPrint
}

// ElevatorPrint ...
type ElevatorPrint struct {
	ID           string
	Status       string
	CurrentFloor string
}

// PrintBattery ...
func (b *Board) PrintBattery(bat *battery.Battery) {
	b.bat = &BatteryPrint{}
	b.bat.initBatPrint(bat)
}

func (b *BatteryPrint) initBatPrint(bat *battery.Battery) {
	b.ID = strconv.Itoa(bat.ID)
	fmt.Println("		| - Battery : " + b.ID)
	prints.LineLine(2)
	fmt.Println("		|   | - Column List")
	prints.LineLine(3)

	for i := 0; i < len(bat.ColumnList); i++ {
		col := &ColumnPrint{}
		col.intiColumnPrint(i, bat)
		b.ColList = append(b.ColList, *col)
		if i < len(bat.ColumnList)-1 {
			prints.LineLine(1)
		}
	}

	fmt.Println()
}

func (c *ColumnPrint) intiColumnPrint(i int, bat *battery.Battery) {
	c.ID = bat.ColumnList[i].ID
	c.Status = bat.ColumnList[i].Status
	c.MinRange = strconv.Itoa(bat.ColumnList[i].MinRange)
	c.MaxRange = strconv.Itoa(bat.ColumnList[i].MaxRange)

	fmt.Println("		|       | - Column : " + c.ID)
	fmt.Println("		|           | - Status : " + c.Status)
	fmt.Println("		|           | - Minimun Range - " + c.MinRange)
	fmt.Println("		|           | - Maximum Range - " + c.MaxRange)
	fmt.Println("		|           | - Elevator List")
	prints.LineLine(5)

	for j := 0; j < len(bat.ColumnList[i].ElevatorList); j++ {
		elev := &ElevatorPrint{}
		elev.intiElevatorPrint(i, j, bat)
		c.ElevList = append(c.ElevList, *elev)
		if j < len(bat.ColumnList[i].ElevatorList)-1 {
			prints.LineLine(1)
		}
	}
}

func (e *ElevatorPrint) intiElevatorPrint(i int, j int, bat *battery.Battery) {
	e.ID = bat.ColumnList[i].ElevatorList[j].ID
	e.Status = bat.ColumnList[i].ElevatorList[j].Status
	e.CurrentFloor = strconv.Itoa(bat.ColumnList[i].ElevatorList[j].CurrentFloor)

	fmt.Println("		|               | - Elevator : " + e.ID)
	fmt.Println("		|                   | - Status : " + e.Status)
	fmt.Println("		|                   | - Current Floor : " + e.CurrentFloor)
	prints.LineLine(1)
}
