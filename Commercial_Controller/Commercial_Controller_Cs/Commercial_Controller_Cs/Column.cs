using System;
using System.Collections.Generic;

namespace Elevator_Controller_CSharp
{
    class Column
    {
        public int ID; // The id of the column
        public string status = "ACTIVE"; // The status of the column, set to "ACTIVE" at the creation of the object
        public int basementAmount; // The amount of basement in the battery
        public int maxRange; // The maximum served floor of the column
        public int minRange; // The minimum served floor of the column
        public List<Elevator> elevatorList = new List<Elevator>(); // The list that will countain all the elevator of the column

        // The constructor method for the Column that will be used when creating a new Column object 
        public Column(int _floorAmount, int _basementAmount, int _elevatorColumn, int _floorColumn, int _iteration, int _previousMax, int _remainder) // The parameter used for creating a column
        {
            ID = _iteration + 1; // Setting the "ID" variable of the column to the iteration which start a 0 so we are adding 1 to the iteration so that we don't get a column id of 0
            basementAmount = _basementAmount; // Setting the number of basement to the number of basement in the battery past has a parameter of the constructor

            // If there is no basement in the battery do the following
            if (basementAmount == 0) 
            {
                // If we are creating the the first column do ...
                if (_iteration == 0)
                {
                    maxRange = _floorColumn; // Set the maximun floor that the column serve to the number of floor that the column will served
                    minRange = 1; // set the minimum floor that the column serve to the ground floor "1"
                }

                // For all the column(s) created after the first one do ...
                else
                {
                    maxRange = _iteration * _floorColumn + _remainder; // Set the maximun floor that the column serve to the iteration time the number of floor that the column will serve and adding the remainder for the last column
                    minRange = _previousMax + 1; // set the minimum floor that the column will serve to the maximun range of the previous column plus 1, past as parameter in the constructor
                }
            }

            // If there is basement in the batterry do the following
            else
            {
                if (_iteration == 0) // If we are creating the first column of the battery do ...
                {
                    maxRange = -1; // Set the maximum floor that the column will serve to minus 1 (the floor before the ground floor), the first column will serve all the basement in the battery
                    minRange = -basementAmount; // Set the minimun floor that the column will serve to the opposite number of basement in the battery
                }

                else if (_previousMax == -1) // If we are creating the second column of the battery do ...
                {
                    maxRange = _iteration * _floorColumn; // Set the maximum floor that the column will serve to the iteration time the number of floor serve by 1 column
                    minRange = _previousMax + 2; // Set the minimum range that the column will serve to "1" and skipping the 0
                }

                else // For all the column(s) created after the second one do ...
                {
                    maxRange = _iteration * _floorColumn + _remainder; // Set to maximum floor that the column will serve to the iteration time the number of floor serve by one column
                    minRange = _previousMax + 1; // set the minimum floor that the column will serve to the maximun range of the previous column plus 1, past as parameter in the constructor
                }
            }

            // use a for loop to create all the elevator that will be added to the list of elevator using the number of elevator per column past has parameter in the constructor
            for (int i = 0; i < _elevatorColumn; i++)
            {
                Elevator elevator = new Elevator(i + 1, _floorAmount, basementAmount, minRange); // creating an elevator at each iteration
                elevatorList.Add(elevator); // adding the elevator created to the list
            }
        }

        // this method is called in the method "columnSelection" method of the battery, its passing the value of the request usefull for the selection of the best elevator
        public void request(int _floor, int _stop, string _direction)
        {
            int n = 1;

            // foreach loop that iterate through all of the elevator in the list of elevator
            foreach (Elevator elevator in elevatorList)
            {
                if (_floor == 1) // if the request was made from the ground floor call the method of the elevator to update its point base on the lobby method
                {
                    elevator.pointsUpdateLobby(_floor, _direction, maxRange, minRange);
                    n = 1;
                }

                else // if the request was made from any other floor than the lobby update the pointing of the elevator with a different method
                {
                    elevator.pointsUpdateFloor(_floor, _direction, maxRange, minRange);
                }
            }
            
            elevatorToSend(_floor, _stop, _direction, n); // when all the elevator(s) will have there points updated, pass the request to the method "elevatorToSend"
        }

        // a method used for sending the information of the request to the best elevator to send for that particular request
        public void elevatorToSend(int _floor, int _stop, string _direction, int n)
        {
            elevatorList.Sort((x, y) => x.points.CompareTo(y.points)); // sort the list of elevator in incressing of points

            Elevator bestOption = elevatorList[0]; // set a new variable the the first index of the list of elevator (sorted by points)

            // call one of the two method to add a stop to the stop list of the best option determined by where the request was made from (floor)
            if (n == 1)
            {
                bestOption.addStopFloor(_floor, _stop, _direction);
            }

            else
            {
                bestOption.addStopLobby(_floor, _stop, _direction);
            }

            Console.WriteLine("Request from Floor {0} and going to Floor {1}", _floor, _stop);
            Console.WriteLine("the elevator chosen is : {0}", bestOption.ID);
            bestOption.all0Remove(); // method called for removing the element(s) that where all ready in the list
            bestOption.listSort(); // method of the elevator to sort the list before moving the elevator
            bestOption.run(); // use the run method of the elevator to move it to the floor(s) in its queue list
            runAll(); // call the run all method from the column
        }

        // method used to run all elevator in the list of elevator
        public void runAll()
        {
            foreach (Elevator elevator in elevatorList)
            {
                elevator.run();
            }
        }

        public void changeValue(int _elevator, List<int> _stopList, string _status, int _currentFloor, string _currentDirection)
        {
            elevatorList[_elevator].changeValue(_stopList, _status, _currentFloor, _currentDirection);
        }
    }
}
