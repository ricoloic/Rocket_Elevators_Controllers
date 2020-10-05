using System;
using System.Collections.Generic;
using System.Text;

namespace Elevator_Controller_CSharp
{
    class Battery
    {
        public int columnBattery;
        public int floorAmount;
        public int basementAmount;
        public int elevatorColumn;
        public List<Column> columnList = new List<Column>();
        public int floorPerColumn;

        public Battery(int _columnBattery, int _floorAmount, int _basementAmount, int _elevatorColumn)
        {
            columnBattery = _columnBattery;
            floorAmount = _floorAmount;
            basementAmount = _basementAmount;
            elevatorColumn = _elevatorColumn;
            int previousMax = 0;

            if (_basementAmount == 0) { floorPerColumn = floorAmount / columnBattery; }
            else { floorPerColumn = floorAmount / (columnBattery - 1); }

            for (int i = 0; i < columnBattery; i++)
            {
                if (i != 0) { previousMax = columnList[i - 1].maxRange; }

                Column column = new Column(floorAmount, basementAmount, elevatorColumn, floorPerColumn, i, previousMax);

                columnList.Add(column);
            }
        }

        public void columnSelection(int _stop, int _floor, string _direction)
        {
            if (_stop == basementAmount + 1)
            {
                foreach (Column column in columnList)
                {
                    if (_floor >= column.minRange && _floor <= column.maxRange) { column.request(_floor, _stop, _direction); }
                }
            }

            else
            {
                foreach (Column column in columnList)
                {
                    if (_stop >= column.minRange && _stop <= column.maxRange) { column.request(_floor, _stop, _direction); }
                }
            }
        }

        public void changeValue(int _column, int _elevator, List<int> _stopList, string _status, int _currentFloor, string _currentDirection)
        {
            columnList[_column].changeValue(_elevator, _stopList, _status, _currentFloor, _currentDirection);
        }
    }
}
