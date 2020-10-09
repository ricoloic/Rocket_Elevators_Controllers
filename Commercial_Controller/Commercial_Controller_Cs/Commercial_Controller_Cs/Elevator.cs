using System;
using System.Collections.Generic;

namespace Elevator_Controller_CSharp
{
    class Elevator
    {
        public int ID; // The id of the elevator
        public int floorAmount; // the number of floor in the battery
        public int basementAmount; // the number of basement in the battery
        public int points; // The points of the elevator
        public List<int> stopList = new List<int>(); // the list that will countain the stops of the elevator
        public List<int> downBuffer = new List<int>(); // the list that will coutain the stops of the elevator !Buffer!
        public List<int> upBuffer = new List<int>(); // the list that will coutain the stops of the elevator !Buffer!
        public string currentDirection = "stop"; // the current direction of the elevator set to stop at creation
        public string previousDirection = "stop"; // the previous direction of the elevator set to stop at creation
        public int currentFloor; // the current floor of the elevator
        public int previousFloor; // the previous floor where the elevator was
        public string door = "closed"; // the state of the elevator door set to closed at creation
        public string status = "IDLE"; // the status of the elevator set to IDLE at creation ! IDLE MOVING MAINTENANCE

        // the constructor for the elevator
        public Elevator(int _id, int _floorAmount, int _basementAmount, int _minRange)
        {
            ID = _id; // set the id of the elevator to the iteration
            floorAmount = _floorAmount; // set the number of floor to the number of floor in the battery
            basementAmount = _basementAmount; // set the number of basement to the number of basement in the battery
            currentFloor = _minRange; // set the current floor the minimun floor of the column past as parameter in the constructor
            previousFloor = currentFloor; // set the previous floor to the current floor
        }

        // return 0 if the number provided is in one of the list else return the number provided
        public int allCheck(int num)
        {
            if (checkIn(num))
            {
                num = 0;
            }
            return num;
        }

        // removing all the element 0 from all the lists
        public void all0Remove()
        {
            stopList = remove0fromList(stopList);
            upBuffer = remove0fromList(upBuffer);
            downBuffer = remove0fromList(downBuffer);
        }

        // method used for removing all element (int) 0 in a given list
        public List<int> remove0fromList(List<int> _nb)
        {
            List<int> awaitList = new List<int>(); // set a new empty list to be used to make the new list whit out the zeros'

            foreach (int nb in _nb) // use a for each loop to iterate through all element in the provided list 
            {
                if (nb != 0) // if the element is not zero do ...
                {
                    awaitList.Add(nb); // add the non-zero element to the await list
                }
            }

            return awaitList; // return the new list / the list with the elements zero removed
        }

        // this method check if there is a given element in some list
        public bool checkIn(int n)
        {
            bool inList = false; // creating a new variable of type bool and set its value to false

            // the next comments apply for all 3 for loop
            foreach (int stop in stopList) // iterate through a list of number (int)
            {
                // If the element is equal to one of the element in the list do ...
                if (n == stop)
                {
                    inList = true; // set the value of the bool variable to true
                }
            }

            foreach (int stop in upBuffer)
            {
                if (n == stop)
                {
                    inList = true;
                }
            }

            foreach (int stop in downBuffer)
            {
                if (n == stop)
                {
                    inList = true;
                }
            }

            // if the element is equal to the current floor of the elevator do ...
            if (n == currentFloor)
            {
                inList = true; // set the value of the bool variable to true
            }

            return inList; // return the bool variable
        }

        // method that return the absolute value of a given number
        public int positive(int n)
        {
            // if the number is smaller than zero => multiply that number by -1
            if (n < 0) { n *= -1; }
            return n; // return the number
        }

        // method to change and print the door state of the elevator
        public void doorState()
        {
            Console.WriteLine("The Elevator {0} has arrived at Floor {1}", ID, currentFloor);

            door = "open";
            Console.WriteLine("The door are {0}", door);

            door = "close";
            Console.WriteLine("The door are {0}\n", door);
        }

        // this method sort the list based on the current direction of the elevator
        public void listSort()
        {
            if (currentDirection == "down") { stopList.Sort((x, y) => y.CompareTo(x)); }

            else { stopList.Sort((x, y) => x.CompareTo(y)); }
        }


        /// <summary>
        /// This method is used to give the elevator some point | this method is only used if the request was made from anywhere in the building apart from the ground floor
        /// It will first set a variable for the difference between the current floor of the elevator and the "_floor"
        /// It will check if the elevator is IDLE or not and if not it is gonna set a new variable for the difference between the last index of the list of request and the "_floor"
        /// The less point the better!
        /// If elevator is going in the same direction and the "_floor" is in the path of the elevator / set point with the length of the stop list + the differance floor
        /// if IDLE / set points to min range + the differance floor
        /// if same direction not in the path / set point to max range + differance last stop + length of stop list
        /// if not same direction / set point to max range * 2 + differance last stop + length of stop list
        /// </summary>
        /// <param name="_floor">The floor at which the request was made</param>
        /// <param name="_direction">The direction of the request</param>
        /// <param name="_maxRange">The maximum floor served by the column</param>
        /// <param name="_minRange">The minimum floor served by the column</param>
        public void pointsUpdateFloor(int _floor, string _direction, int _maxRange, int _minRange)
        {
            int differenceLastStop = 0;
            int differenceFloor = positive(currentFloor - _floor);
            points = 0;

            if (status != "IDLE")
            {
                differenceLastStop = positive(stopList[stopList.Count - 1] - _floor);
            }

            bool conditionInPath = (_floor >= currentFloor && _direction == "up") || (_floor <= currentFloor && _direction == "down");
            bool conditionNotInPath = (_floor < currentFloor && _direction == "up") || (_floor > currentFloor && _direction == "down");

            if (status == "IDLE")
            {
                if (_maxRange < 0)
                {
                    points = positive(_minRange) + differenceFloor + 1;
                }

                else
                {
                    points = positive(_maxRange) + differenceFloor + 1;
                }
                
            }

            else if (currentDirection == _direction)
            {
                if (conditionInPath)
                {
                    points = differenceFloor + stopList.Count;
                }

                else if (conditionNotInPath)
                {
                    points = positive(_maxRange) + differenceLastStop + stopList.Count;
                }
            }

            else if (currentDirection != _direction) {
                points = positive(_maxRange) * 2 + differenceLastStop + stopList.Count;
            }

            Console.WriteLine("{0}, {1}", ID, points);
        }

        /// <summary>
        /// This method is used to give the elevator some point | this method is only used if the request was made from the ground floor
        /// It will first set a variable for the difference between the current floor of the elevator and the "_floor"
        /// It will check if the elevator is IDLE or not and if not it is gonna set a new variable for the difference between the last index of the list of request and the "_floor"
        /// The less point the better!
        /// If elevator is not going in the same direction as the user direction / set point with the differance floor + the difference last stop
        /// If IDLE / set points to min range + the differance floor + 1
        /// If Elevator is in the same direction as the user direction / set point to max range * 2 + differance last stop + length of stop list
        /// If the current floor of the elevator is equal to the "_floor" / set point to the length of stop list
        /// </summary>
        /// <param name="_floor">The floor at which the request was made</param>
        /// <param name="_direction">The direction of the request</param>
        /// <param name="_maxRange">The maximum floor served by the column</param>
        /// <param name="_minRange">The minimum floor served by the column</param>
        public void pointsUpdateLobby(int _floor, string _direction, int _maxRange, int _minRange)
        {
            int differenceLastStop = 0;
            int differenceFloor = positive(currentFloor - _floor);
            points = 0;

            if (status != "IDLE")
            {
                differenceLastStop = positive(stopList[stopList.Count - 1] - _floor);
            }

            if (status == "IDLE")
            {
                if (_maxRange < 0)
                {
                    points = positive(_minRange) + differenceFloor + 1;
                }

                else
                {
                    points = positive(_maxRange) + differenceFloor + 1;
                }
            }

            else if (_direction != currentDirection)
            {
                points = differenceLastStop + differenceFloor;
            }

            else if (currentDirection == _direction)
            {
                points = positive(_maxRange) * 2 + stopList.Count + differenceLastStop;
            }

            if (currentFloor == _floor)
            {
                points = stopList.Count;
            }

            Console.WriteLine("{0}, {1}", ID, points);
        }

        /// <summary>
        /// To long to describe and I'm tired of writing comments
        /// but the big idea here is to add both the stop of the user and the current floor of the user to the good stop list of the elevator "stopList, upbuffer, downBuffer"
        /// </summary>
        /// <param name="_floor"></param>
        /// <param name="_stop"></param>
        /// <param name="_direction"></param>

        public void addStopLobby(int _floor, int _stop, string _direction)
        {
            int floor = allCheck(_floor);
            int stop = allCheck(_stop);

            if (_direction != currentDirection && _floor <= currentFloor)
            {
                stopList.Add(floor);
                upBuffer.Add(stop);
            }

            else if (_direction != currentDirection && _floor >= currentFloor)
            {
                stopList.Add(floor);
                downBuffer.Add(stop);
            }

            else if (status == "IDLE")
            {
                stopList.Add(floor);

                if (_direction == "up")
                {
                    upBuffer.Add(stop);
                }

                else if (_direction == "down")
                {
                    downBuffer.Add(stop);
                }
            }

            else if (_direction == currentDirection)
            {
                if (_floor == currentFloor)
                {
                    stopList.Add(floor);
                    stopList.Add(stop);
                }

                else if (_floor != currentFloor)
                {
                    stopList.Add(floor);

                    if (_direction == "up")
                    {
                        upBuffer.Add(stop);
                    }

                    else if (_direction == "down")
                    {
                        downBuffer.Add(stop);
                    }
                }

                else if (_floor < currentFloor)
                {
                    downBuffer.Add(floor);
                    upBuffer.Add(stop);
                }

                else if (_floor > currentFloor)
                {
                    upBuffer.Add(floor);
                    downBuffer.Add(stop);
                }
            }
        }

        public void addStopFloor(int _floor, int _stop, string _direction)
        {
            int floor = allCheck(_floor);
            int stop = allCheck(_stop);
            
            if (status == "IDLE")
            {
                stopList.Add(floor);

                if (_direction == "up")
                {
                    upBuffer.Add(stop);
                }

                else if (_direction == "down")
                {
                    downBuffer.Add(stop);
                }
            }

            else if (_direction == currentDirection)
            {
                if ((_direction == "up" && _floor >= currentFloor) || (_direction == "down" && _floor <= currentFloor))
                {
                    stopList.Add(floor);
                    stopList.Add(stop);
                }

                else if (_direction == "up" && _floor < currentFloor)
                {
                    downBuffer.Add(floor);
                    upBuffer.Add(stop);
                }

                else if (_direction == "down" && _floor > currentFloor)
                {
                    upBuffer.Add(floor);
                    downBuffer.Add(stop);
                }
            }

            else if (_direction != currentDirection)
            {
                if (_direction == "up")
                {
                    upBuffer.Add(floor);
                    upBuffer.Add(stop);
                }

                else if (_direction == "down")
                {
                    downBuffer.Add(floor);
                    downBuffer.Add(stop);
                }
            }
        }

        // method used for updating the list of stop with a buffer based on the previous direction and the length of the "up and down buffer"
        public void stopSwitch()
        {
            if (upBuffer.Count != 0 && downBuffer.Count != 0)
            {
                if (previousDirection == "up")
                {
                    stopList = downBuffer;
                    downBuffer = new List<int>();
                }

                else if (previousDirection == "down")
                {
                    stopList = upBuffer;
                    upBuffer = new List<int>();
                }
            }

            else if (downBuffer.Count != 0 && upBuffer.Count == 0)
            {
                stopList = downBuffer;
                downBuffer = new List<int>();
            }

            else if (downBuffer.Count == 0 && upBuffer.Count != 0)
            {
                stopList = upBuffer;
                upBuffer = new List<int>();
            }

            else if (downBuffer.Count == 0 && upBuffer.Count == 0)
            {
                status = "IDLE";
                currentDirection = "stop";
            }

            if (stopList.Count != 0)
            {
                run();
            }
        }

        // method used for moving the elevator and calling the oppening and closing of the doors when arrived at destination
        public void run()
        {
            if (currentDirection != previousDirection)
            {
                previousDirection = currentDirection;
            }

            while (stopList.Count != 0)
            {
                while(currentFloor != stopList[0])
                {
                    status = "MOVING";

                    if (stopList[0] < currentFloor)
                    {
                        currentDirection = "down";
                        currentFloor--;
                    }

                    else if (stopList[0] > currentFloor)
                    {
                        currentDirection = "up";
                        currentFloor++;
                    }
                }

                if (currentFloor == stopList[0] && previousFloor != currentFloor)
                {
                    doorState();
                    previousFloor = stopList[0];
                    stopList.RemoveAt(0);
                }

                else if (stopList.Count != 0 && previousFloor == currentFloor)
                {
                    stopList.RemoveAt(0);
                }
            }

            if (stopList.Count == 0)
            {
                stopSwitch();
            }
        }

        public void changeValue(List<int> _stopList, string _status, int _currentFloor, string _currentDirection)
        {
            currentFloor = _currentFloor;
            previousFloor = currentFloor;
            stopList = _stopList;
            currentDirection = _currentDirection;
            status = _status;
            listSort();
        }
    }
}