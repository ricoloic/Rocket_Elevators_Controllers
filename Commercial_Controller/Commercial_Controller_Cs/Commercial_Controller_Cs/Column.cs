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
        public int floorPerColumn;
        public Elevator bestOption;
        public bool lobby;
        // public bool atFirstIteration;
        public List<Elevator> elevatorList = new List<Elevator>(); // The list that will countain all the elevator of the column

        public Column(int _floorAmount, int _basementAmount, int _elevatorColumn, int _floorColumn, int _iteration, int _previousMax, int _remainder) // The parameter used for creating a column
        {
            ID = _iteration + 1; // Setting the "ID" variable of the column to the iteration which start a 0 so we are adding 1 to the iteration so that we don't get a column id of 0
            basementAmount = _basementAmount; // Setting the number of basement to the number of basement in the battery past has a parameter of the constructor
            floorPerColumn = _floorColumn;

            setRange(_iteration, _remainder, _previousMax);
            createElevatorList(_iteration, _elevatorColumn, _floorAmount);
        }

        // createElevatorList will add Elevator(s) to the column
        public void createElevatorList(int _i, int elevatorPerColumn, int floorAmount)
        {
            // use a for loop to create all the elevator that will be added to the list of elevator using the number of elevator per column past has parameter in the constructor
            for (int i = 0; i < elevatorPerColumn; i++)
            {
                Elevator elevator = new Elevator(i + 1, floorAmount, basementAmount, minRange); // creating an elevator at each iteration
                elevatorList.Add(elevator); // adding the elevator created to the list
            }
        }

        // setRange will call the a method to make the min and max range of the column
        public void setRange(int i, int remainder, int previousMax)
        {
            if (basementAmount == 0)
            {
                ifNoBasement(i, remainder, previousMax);
            }

            else
            {
                ifBasement(i, remainder, previousMax);
            }
        }

        // ifNoBasement set the ranges of the column for a battery without any basement
        public void ifNoBasement(int i, int remainder, int previousMax)
        {
            if (i == 0)
            {
                maxRange = floorPerColumn;
                minRange = 1;
            }

            else
            {
                maxRange = i * floorPerColumn + remainder;
                minRange = previousMax + 1;
            }
        }

        // ifBasement set the ranges of the column for a battery with basement
        public void ifBasement(int i, int remainder, int previousMax)
        {
            if (i == 0)
            {
                maxRange = -1;
                minRange = -basementAmount;
            }

            else if (i == 2)
            {
                maxRange = i * floorPerColumn;
                minRange = previousMax + 2;
            }

            else
            {
                maxRange = i * floorPerColumn + remainder;
                minRange = previousMax + 1;
            }
        }

        // initRequest will call methods to make a request possible
        public void initRequest(int _floor, int _stop, string _direction)
        {
            isAtLobby(_floor);
            updatePoints(_floor, _direction);
            sortByPoint();
            addStop(_floor, _stop, _direction);
            runAll();
        }

        // isAtLobby will set the boolean variable lobby to true if the request is made from the lobby
        public void isAtLobby(int _floor)
        {
            if (_floor == 1)
            {
                lobby = true;
            }

            else
            {
                lobby = false;
            }
        }

        // updatePoints will call the method for updating the points of every elevator in the column
        public void updatePoints(int _floor, string _direction)
        {
            for (int i = 0; i < elevatorList.Count; i++)
            {
                if (lobby)
                {
                    elevatorList[i].pointsUpdateLobby(_floor, _direction, maxRange, minRange);
                }

                else
                {
                    elevatorList[i].pointsUpdateFloor(_floor, _direction, maxRange, minRange);
                }
            }
            Console.WriteLine();
        }

        // sortByPoint will reorganize the list of elevators in such a way that the first index of the list is the elevator with the fewest points then set the variable bestOption to the fist index in the list of elevators
        public void sortByPoint()
        {
            elevatorList.Sort((x, y) => x.points.CompareTo(y.points)); // sort the list of elevator in incressing of points

            bestOption = elevatorList[0];
            Console.WriteLine("The elevator {0} is sent", bestOption.ID);
        }

        // addStop will call the method of the bestOption to add the request to its list
        public void addStop(int _floor, int _stop, string _direction)
        {
            if (!lobby)
            {
                bestOption.addStopFloor(_floor, _stop, _direction);
            }

            else
            {
                bestOption.addStopLobby(_floor, _stop, _direction);
            }
        }

        // a method used for sending the information of the request to the best elevator to send for that particular request
        public void elevatorToSend(int _floor, int _stop, string _direction, int n)
        {
            

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

            runAll(); // call the run all method from the column
        }

        // method used to run all elevator in the list of elevator
        public void runAll()
        {
            foreach (Elevator elevator in elevatorList)
            {
                elevator.all0Remove(); // method called for removing the element(s) that where all ready in the list
                elevator.listSort(); // method of the elevator to sort the list before moving the elevator
                elevator.run(); // use the run method of the elevator to move it to the floor(s) in its queue list
                Console.WriteLine();
            }
        }

        public void changeValue(int _elevator, List<int> _stopList, string _status, int _currentFloor, string _currentDirection)
        {
            elevatorList[_elevator].changeValue(_stopList, _status, _currentFloor, _currentDirection);
        }
    }
}
