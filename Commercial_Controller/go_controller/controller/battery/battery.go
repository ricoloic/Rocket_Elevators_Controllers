package battery

import (
	"go_controller/controller"
	"go_controller/controller/battery/column"
)

// Battery ...
type Battery struct {
	columnBattery  int
	floorAmount    int
	basementAmount int
	elevatorColumn int
	floorPerColumn int
	Letters        []string
	columnList     []column.Column
}

// StartBattery ...
func (b *Battery) StartBattery(_columnBattery int, _floorAmount int, _basementAmount int, _elevatorColumn int) {
	b.columnBattery = _columnBattery
	b.floorAmount = _floorAmount
	b.basementAmount = _basementAmount
	b.elevatorColumn = _elevatorColumn
	b.Letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	remainder := 0
	waitLast := 0

	previousMax := -b.basementAmount

	if _basementAmount == 0 {
		remainder = b.floorAmount % b.columnBattery
		b.floorPerColumn = (b.floorAmount - remainder) / b.columnBattery

	} else {
		remainder = b.floorAmount % (b.columnBattery - 1)
		b.floorPerColumn = (b.floorAmount - remainder) / (b.columnBattery - 1)
	}

	for i := 0; i < b.columnBattery; i++ {
		if i > 0 {
			previousMax = b.columnList[i-1].MaxRange
		}

		if i == b.columnBattery-1 {
			waitLast = remainder
		}

		col := &column.Column{}
		col.StartColumn(b.floorAmount, b.basementAmount, b.elevatorColumn, b.floorPerColumn, i, previousMax, waitLast, b.Letters)
		b.columnList = append(b.columnList, *col)
	}
}

// ColumnSelection ...
func (b *Battery) ColumnSelection(_floor int, _stop int, _direction string) {
	for i := 0; i < len(b.columnList); i++ {
		if _stop == 1 {
			if _floor >= b.columnList[i].MinRange && _floor <= b.columnList[i].MaxRange {
				controller.Wait(2)
				b.columnList[i].Request(_floor, _stop, _direction)
			}

		} else {
			if _stop >= b.columnList[i].MinRange && _stop <= b.columnList[i].MaxRange {
				controller.Wait(2)
				b.columnList[i].Request(_floor, _stop, _direction)
			}
		}
	}
}

// ChangeValueB ...
func (b *Battery) ChangeValueB(_column int, _elevator int, _status string, _currentFloor int, _stopList []int, _currentDirection string) {
	b.columnList[_column].ChangeValueC(_elevator, _status, _currentFloor, _stopList, _currentDirection)
}
