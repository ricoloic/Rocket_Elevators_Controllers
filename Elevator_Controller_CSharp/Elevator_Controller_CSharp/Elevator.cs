using System;
using System.Collections.Generic;

namespace Elevator_Controller_CSharp
{
    class Elevator
    {
        public int ID;
        public int floorAmount;
        public int basementAmount;
        public int points;
        public List<int> stopList = new List<int>();
        public List<int> downBuffer = new List<int>();
        public List<int> upBuffer = new List<int>();
        public string currentDirection = "stop";
        public string previousDirection = "stop";
        public int currentFloor;
        public int previousFloor;
        public string door = "closed";
        public string status = "IDLE";

        public Elevator(int _id, int _floorAmount, int _basementAmount, int minRange)
        {
            ID = _id;
            floorAmount = _floorAmount;
            basementAmount = _basementAmount;
            currentFloor = minRange;
            previousFloor = currentFloor;
        }

        public int positive(int n)
        {
            if (n < 0) { n *= -1; }
            return n;
        }

        public void DoorState()
        {
            door = "open";
            Console.WriteLine("The door are {0}", door);

            door = "close";
            Console.WriteLine("The door are {0}", door);
        }

        public void listSort()
        {
            if (currentDirection == "down") { stopList.Sort((x, y) => y.CompareTo(x)); }

            else { stopList.Sort((x, y) => x.CompareTo(y)); }
        }

        public void sendRequest(int requestedFloor)
        {
            stopList.Add(requestedFloor);
            status = "MOVING";

            if (currentFloor > requestedFloor) { currentDirection = "down"; }

            else if (currentFloor < requestedFloor) { currentDirection = "up"; }

            listSort();
        }

        public void pointsUpdateFloor(int _floor, string _direction, int maxRange)
        {
            int differenceLastStop = 0;

            if (status != "IDLE")
            {
                differenceLastStop = positive(stopList[stopList.Count - 1] - _floor);
            }

            int differenceFloor = positive(currentFloor - _floor);
            points = 0;

            if (status == "IDLE") { points = maxRange + differenceFloor + 1; }

            else if (currentDirection == _direction)
            {
                if (_floor >= currentFloor && _direction == "up") { points = differenceFloor + stopList.Count; }

                else if (_floor <= currentFloor && _direction == "down") { points = differenceFloor + stopList.Count; }

                else if (_floor < currentFloor && _direction == "up" || _floor > currentFloor && _direction == "down")
                {
                    points = maxRange + differenceLastStop + stopList.Count;
                }
            }

            else if (currentDirection != _direction) { points = maxRange * 2 + differenceLastStop + stopList.Count; }
        }

        public void pointsUpdateLobby(int _floor, string _direction, int maxRange)
        {
            int differenceLastStop = 0;

            if (status != "IDLE")
            {
                differenceLastStop = positive(stopList[stopList.Count - 1] - _floor);
            }

            int differenceFloor = positive(currentFloor - _floor);
            points = 0;

            if (status == "IDLE") { points = maxRange + differenceLastStop + 1; }

            else if (_direction != currentDirection) { points = differenceLastStop + differenceFloor; }

            else if (currentDirection == _direction) { points = maxRange * 2 + stopList.Count + differenceLastStop; }

            if (currentFloor == _floor) { points = stopList.Count; }
        }

        public void addStop(int _floor, int _stop, string _direction)
        {
            if (_floor == basementAmount + 1)
            {
                if (_direction != currentDirection && _floor <= currentFloor)
                {
                    stopList.Add(_floor);
                    upBuffer.Add(_stop);
                }

                else if (_direction != currentDirection && _floor >= currentFloor)
                {
                    stopList.Add(_floor);
                    downBuffer.Add(_stop);
                }

                else if (status == "IDLE")
                {
                    stopList.Add(_floor);

                    if (_direction == "up") { upBuffer.Add(_stop); }

                    else if (_direction == "down") { downBuffer.Add(_stop); }
                }

                else if (_direction == currentDirection)
                {
                    if (_floor != currentFloor)
                    {
                        if (_direction == "up")
                        {
                            stopList.Add(_floor);
                            upBuffer.Add(_stop);
                        }

                        else if (_direction == "up")
                        {
                            stopList.Add(_floor);
                            downBuffer.Add(_stop);
                        }
                    }

                    else if (_floor < currentFloor)
                    {
                        downBuffer.Add(_floor);
                        upBuffer.Add(_stop);
                    }

                    else if (_floor > currentFloor)
                    {
                        upBuffer.Add(_floor);
                        downBuffer.Add(_stop);
                    }
                }
            }

            else
            {
                if (status == "IDLE")
                {
                    stopList.Add(_floor);

                    if (_direction == "up") { upBuffer.Add(_stop); }

                    else if (_direction == "down") { downBuffer.Add(_stop); }
                }

                else if (_direction == currentDirection)
                {
                    if (_direction == "up" && _floor >= currentFloor)
                    {
                        stopList.Add(_stop);
                        stopList.Add(_stop);
                    }

                    else if (_direction == "down" && _floor <= currentFloor)
                    {
                        stopList.Add(_stop);
                        stopList.Add(_stop);
                    }

                    else if (_direction == "up" && _floor < currentFloor)
                    {
                        downBuffer.Add(_floor);
                        upBuffer.Add(_stop);
                    }

                    else if (_direction == "down" && _floor > currentFloor)
                    {
                        upBuffer.Add(_floor);
                        downBuffer.Add(_stop);
                    }
                }

                else if (_direction != currentDirection)
                {
                    if (_direction == "up")
                    {
                        upBuffer.Add(_floor);
                        upBuffer.Add(_stop);
                    }

                    else if (_direction == "down")
                    {
                        downBuffer.Add(_floor);
                        downBuffer.Add(_stop);
                    }
                }
            }

            listSort();
        }

        public void stopSwitch()
        {
            if (upBuffer.Count != 0 && downBuffer.Count != 0)
            {
                if (previousDirection == "up") { foreach (int i in downBuffer) { stopList.Add(i); } }

                else if (previousDirection == "down") { foreach (int i in upBuffer) { stopList.Add(i); } }
            }

            else if (downBuffer.Count != 0 && upBuffer.Count == 0) { foreach (int i in downBuffer) { stopList.Add(i); } }

            else if (downBuffer.Count == 0 && upBuffer.Count != 0) { foreach (int i in upBuffer) { stopList.Add(i); } }

            else if (upBuffer.Count == 0 && upBuffer.Count == 0)
            {
                status = "IDLE";
                currentDirection = "stop";
            }
        }

        public void run()
        {
            if (stopList.Count != 0)
            {
                while (stopList.Count != 0)
                {
                    while (currentFloor != stopList[0])
                    {
                        status = "MOVING";

                        if (stopList[0] < currentFloor)
                        {
                            currentDirection = "down";
                            previousDirection = currentDirection;
                            currentFloor -= 1;
                        }

                        else if (stopList[0] > currentFloor)
                        {
                            currentDirection = "up";
                            previousDirection = currentDirection;
                            currentFloor += 1;
                        }
                    }

                    if (stopList[0] == currentFloor)
                    {
                        DoorState();
                        stopList.RemoveAt(0);
                    }
                }
            }

            if (stopList.Count == 0) { stopSwitch(); }
        }

        public void changeValue(int _elevator, List<int> _stopList, string _status, int _currentFloor, string _currentDirection)
        {
            currentFloor = _currentFloor;
            stopList = _stopList;
            currentDirection = _currentDirection;
            status = _status;
            listSort();
        }
    }
}