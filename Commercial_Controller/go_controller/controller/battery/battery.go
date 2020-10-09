package battery

import (
	"go_controller/controller"
	"go_controller/controller/battery/column"
	"strconv"
)

// Battery ...
type Battery struct {
	ID             int
	columnBattery  int
	floorAmount    int
	basementAmount int
	elevatorColumn int
	floorPerColumn int
	Letters        []string
	ColumnList     []column.Column
}

// StartBattery will set some initial value to the battery and call the method to create all the column that the battery will countain
func (b *Battery) StartBattery(id int, _columnBattery int, _floorAmount int, _basementAmount int, _elevatorColumn int) {
	b.ID = id
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
	var columnID string
	var t int

	for i := 0; i < b.columnBattery; i++ {
		lastRemain = b.calcLastRemain(i, remainder)
		prevMax := b.calcPrevMax(i)
		columnID, t = b.calcColumnID(i, t)

		col := &column.Column{}
		col.StartColumn(b.floorAmount, b.basementAmount, b.elevatorColumn, b.floorPerColumn, i, prevMax, lastRemain, columnID)
		b.ColumnList = append(b.ColumnList, *col)
	}
}

// calcRemainder will set and return the remainder to add to the last column
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

// calcLastRemain will return the remainder if the loop for creating the columns is at the last iteration
func (b *Battery) calcLastRemain(i int, remainder int) int {
	if i == b.columnBattery-1 {
		return remainder
	}

	return 0
}

// calcPrevMax will return the maxRange from the column at the previous iteration
func (b *Battery) calcPrevMax(i int) int {
	var previousMax int

	if i > 0 {
		previousMax = b.ColumnList[i-1].MaxRange
	}

	return previousMax
}

// calcColumnID will return the the id of the column at its creation -- A - Z , 1A - 1Z , 2A - 2Z , ...
func (b *Battery) calcColumnID(i int, t int) (string, int) {
	if i > 25 {
		iInt := i % 26
		if iInt == 0 {
			t++
		}
		bString := strconv.Itoa(t)
		return bString + b.Letters[i%26], t
	}

	return b.Letters[i], t
}

// ColumnSelection is used to for sending the request to the good column
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
