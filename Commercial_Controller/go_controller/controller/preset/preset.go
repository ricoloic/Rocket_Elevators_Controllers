package preset

import (
	"go_controller/controller/battery"
	"go_controller/controller/prints"
)

// CodeBoxx ...
func CodeBoxx(n int) {
	battery := &battery.Battery{}
	battery.StartBattery(4, 60, 6, 5)

	// ||=========> BASEMENT COLUMN #1 <=========||
	battery.ChangeValueB(0, 0, "IDLE", -4, []int{}, "Stop")
	battery.ChangeValueB(0, 1, "IDLE", 1, []int{}, "Stop")
	battery.ChangeValueB(0, 2, "MOVING", -3, []int{-5}, "Down")
	battery.ChangeValueB(0, 3, "MOVING", -6, []int{1}, "Up")
	battery.ChangeValueB(0, 4, "MOVING", -1, []int{-6}, "Down")

	// ||=========> FLOOR COLUMN #2 <=========||
	battery.ChangeValueB(1, 0, "MOVING", 20, []int{5}, "Down")
	battery.ChangeValueB(1, 1, "MOVING", 3, []int{15}, "Up")
	battery.ChangeValueB(1, 2, "MOVING", 13, []int{1}, "Down")
	battery.ChangeValueB(1, 3, "MOVING", 15, []int{2}, "Down")
	battery.ChangeValueB(1, 4, "MOVING", 6, []int{1}, "Down")

	// ||=========> FLOOR COLUMN #3 <=========||
	battery.ChangeValueB(2, 0, "MOVING", 1, []int{21}, "Up")
	battery.ChangeValueB(2, 1, "MOVING", 23, []int{28}, "Up")
	battery.ChangeValueB(2, 2, "MOVING", 33, []int{1}, "Down")
	battery.ChangeValueB(2, 3, "MOVING", 40, []int{24}, "Down")
	battery.ChangeValueB(2, 4, "MOVING", 39, []int{1}, "Down")

	// ||=========> FLOOR COLUMN #4 <=========||
	battery.ChangeValueB(3, 0, "MOVING", 58, []int{1}, "Down")
	battery.ChangeValueB(3, 1, "MOVING", 50, []int{60}, "Up")
	battery.ChangeValueB(3, 2, "MOVING", 46, []int{58}, "Up")
	battery.ChangeValueB(3, 3, "MOVING", 1, []int{54}, "Up")
	battery.ChangeValueB(3, 4, "MOVING", 60, []int{1}, "Down")

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

	prints.CreateRequest(f[0], f[1], d)
	battery.ColumnSelection(f[0], f[1], d)
}
