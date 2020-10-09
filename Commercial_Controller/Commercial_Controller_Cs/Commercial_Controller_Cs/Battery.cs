using System.Collections.Generic;

namespace Elevator_Controller_CSharp
{
    class Battery
    {
        public int columnBattery; // Number of column in the battery
        public int floorAmount; // Number of floor in the battery
        public int basementAmount; // Number of basement in the battery
        public int elevatorColumn; // Number of elevator per column
        public int floorPerColumn; // Number of floor serverd by one column (apart the column assign for the basement(s))
        public List<Column> columnList = new List<Column>(); // The list that will countain the columns in the battery

        // Constructor method of the battery class that will be called when ever you create a new Battery object
        public Battery(int _columnBattery, int _floorAmount, int _basementAmount, int _elevatorColumn)
        {
            // Assigning the value past at the creation of a new Battery object
            columnBattery = _columnBattery;
            floorAmount = _floorAmount;
            basementAmount = _basementAmount;
            elevatorColumn = _elevatorColumn;
            int remainder = 0;
            int waitLast = 0;

            int previousMax = -basementAmount; // Making the new variable "previousMax" and set it to the opposite value of the number of basement

            if (_basementAmount == 0)
            {
                remainder = floorAmount % columnBattery; // finding the remainder if there is one, the number of floor and the number of column
                floorPerColumn = (floorAmount - remainder) / columnBattery; // If there is no basement set the number of floor per column to the number of floor divided by the number of collumn in the battery
            }

            else
            {
                remainder = floorAmount % (columnBattery - 1); // finding the remainder if there is one, the number of floor and the number of column minus the basement column
                floorPerColumn = (floorAmount - remainder) / (columnBattery - 1); // If there is some basement set the number of floor per column (apart the column assign for the basement(s)) to the number of floor divided by the number of column in the battery minus the column that will be assign for the basement(s)
            }

            // Loop used for creating all the column(s) in the battery using the number of column in the battery provided when creating the battery
            for (int i = 0; i < columnBattery; i++)
            {
                if (i > 0)
                {
                    previousMax = columnList[columnList.Count - 1].maxRange; // If the loop is not at its first iteration "i" change the variable "previousMax" to the "maxRange" of the previous column created using the column list of the battery at the previous index of iteration
                }

                if (i == columnBattery - 1)
                {
                    waitLast = remainder; // applying the remainder only if we are creating the last column
                }

                Column column = new Column(floorAmount, basementAmount, elevatorColumn, floorPerColumn, i, previousMax, waitLast); // Making a new variable "column" and setting it to a new object of type Column and passing the number of floor in the battery "floorAmount", the number of basement in the battery "basementAmount", the number of elevator in a column "elevatorColumn", the number of floor that the column will serve "floorPerColumn", the iteration at which the loop is "i" and the variable "previousMax" to its constructor

                columnList.Add(column); // Adding the object countain in the variable "column" to the list of column(s) in the battery "columnList"
            }
        }

        // Method that will be called when ever there is a user request using the current floor of the user "_floor", the floor requested "_stop" and the direction of the request "_direction"
        public void columnSelection(int _floor, int _stop, string _direction)
        {
            foreach (Column column in columnList) // Iterate through the list of column using a for each loop (where the variable will be the value of the element in the list and changing at every iteration)
            {
                if (_stop == 1) // If the floor requested is the ground floor "1"
                {
                    if (_floor >= column.minRange && _floor <= column.maxRange) // At every iteration check if the current floor of the user is in the range of that column "minRange maxRange"
                    {
                        column.request(_floor, _stop, _direction); // when the above "if statement" is true call the method of the column to find the best elevator an adding the current floor and the requested floor of the user using the current floor of the user, the requested floor and the direction of the request
                    }
                }

                else // If the user is at the ground floor do almost the same has above but insted of using the current floor of the user, use the floor requested to find the column to send the request to
                {
                    if (_stop >= column.minRange && _stop <= column.maxRange)
                    {
                        column.request(_floor, _stop, _direction);
                    }
                }
            }
        }

        public void changeValue(int _column, int _elevator, string _status, int _currentFloor, List<int> _stopList, string _currentDirection)
        {
            columnList[_column].changeValue(_elevator, _stopList, _status, _currentFloor, _currentDirection);
        }
    }
}
