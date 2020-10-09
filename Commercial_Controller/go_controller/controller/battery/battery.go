package battery

import (
	"go_controller/controller"
	"go_controller/controller/battery/column"
	"strconv"
)

// Battery ...
type Battery struct {
	columnBattery  int
	floorAmount    int
	basementAmount int
	elevatorColumn int
	floorPerColumn int
	Letters        []string
	ColumnList     []column.Column
	columnID       string
	t              int
}

// StartBattery will set some initial value to the battery and call the method to create all the column that the battery will countain
func (b *Battery) StartBattery(_columnBattery int, _floorAmount int, _basementAmount int, _elevatorColumn int) {
	b.columnBattery = _columnBattery
	b.floorAmount = _floorAmount
	b.basementAmount = _basementAmount
	b.elevatorColumn = _elevatorColumn
	b.Letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	b.createColumnList()
}

// createColumnList will create all the columns in the battery using the number of column in the battery
func (b *Battery) createColumnList() {
	remainder := b.calcRemainder()
	lastRemain := 0

	for i := 0; i < b.columnBattery; i++ {
		lastRemain = b.calcLastRemain(i, remainder)
		prevMax := b.calcPrevMax(i)
		b.calcColumnID(i)

		col := &column.Column{}
		col.StartColumn(b.floorAmount, b.basementAmount, b.elevatorColumn, b.floorPerColumn, i, prevMax, lastRemain, b.columnID)
		b.ColumnList = append(b.ColumnList, *col)
	}
}

// calcRemainder will set and return the
func (b *Battery) calcRemainder() int {
	remainder := 0

	if b.basementAmount == 0 {
		remainder = b.floorAmount % b.columnBattery
		b.floorPerColumn = (b.floorAmount - remainder) / b.columnBattery

	} else {
		remainder = b.floorAmount % (b.columnBattery - 1)
		b.floorPerColumn = (b.floorAmount - remainder) / (b.columnBattery - 1)
	}

	return remainder
}

func (b *Battery) calcLastRemain(i int, remainder int) int {
	if i == b.columnBattery-1 {
		return remainder
	} else {
		return 0
	}
}

func (b *Battery) calcPrevMax(i int) int {
	var previousMax int

	if i > 0 {
		previousMax = b.ColumnList[i-1].MaxRange
	}

	return previousMax
}

func (b *Battery) calcColumnID(i int) {
	if i > 25 {
		iInt := i % 26
		if iInt == 0 {
			b.t++
		}
		bString := strconv.Itoa(b.t)
		b.columnID = bString + b.Letters[i%26]
	} else {
		b.columnID = b.Letters[i]
	}
}

// ColumnSelection ...
func (b *Battery) ColumnSelection(_floor int, _stop int, _direction string) {
	for i := 0; i < len(b.ColumnList); i++ {
		if _stop == 1 {
			if _floor >= b.ColumnList[i].MinRange && _floor <= b.ColumnList[i].MaxRange {
				controller.Wait(2)
				b.ColumnList[i].InitializeRequest(_floor, _stop, _direction)
			}

		} else {
			if _stop >= b.ColumnList[i].MinRange && _stop <= b.ColumnList[i].MaxRange {
				controller.Wait(2)
				b.ColumnList[i].InitializeRequest(_floor, _stop, _direction)
			}
		}
	}
}

// ChangeValueB ...
func (b *Battery) ChangeValueB(_column int, _elevator int, _status string, _currentFloor int, _stopList []int, _currentDirection string) {
	b.ColumnList[_column].ChangeValueC(_elevator, _status, _currentFloor, _stopList, _currentDirection)
}
