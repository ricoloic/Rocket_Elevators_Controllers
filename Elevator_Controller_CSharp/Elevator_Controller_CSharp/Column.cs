using System;
using System.Collections;
using System.Collections.Generic;

namespace Elevator_Controller_CSharp
{
    class Column
    {
        public int ID;
        public string status = "ACTIVE";
        public int basementAmount = 0;
        public int maxRange;
        public int minRange;
        public List<Elevator> elevatorList = new List<Elevator>();

        public Column(int _floorAmount, int _basementAmount, int elevatorColumn, int _floorColumn, int _iteration, int _previousMax)
        {
            ID = _iteration + 1;
            basementAmount = _basementAmount;

            if (basementAmount != 0)
            {
                if (_iteration == 0)
                {
                    maxRange = basementAmount;
                    minRange = 1;
                }

                else
                {
                    maxRange = _iteration * _floorColumn + basementAmount + 1;
                    minRange = _previousMax + 1;
                }
            }
            
            else
            {
                if (_iteration == 0)
                {
                    maxRange = _floorColumn;
                    minRange = 1;
                }

                else
                {
                    maxRange = _iteration * _floorColumn + 1;
                    minRange = _previousMax + 1;
                }
            }

            for (int i = 0; i < elevatorColumn; i++)
            {
                Elevator elevator = new Elevator(i + 1, _floorAmount, basementAmount, minRange);
                elevatorList.Add(elevator);
            }
        }

        public void request(int _floor, int _stop, string _direction)
        {
            foreach (Elevator elevator in elevatorList)
            {
                if (_floor == basementAmount + 1) { elevator.pointsUpdateLobby(_floor, _direction, maxRange); }

                else { elevator.pointsUpdateFloor(_floor, _direction, maxRange); }
            }

            elevatorToSend(_floor, _stop, _direction);
        }

        public void elevatorToSend(int _floor, int _stop, string _direction)
        {
            elevatorList.Sort((x, y) => x.points.CompareTo(y.points));

            Elevator bestOption = elevatorList[0];

            bestOption.addStop(_floor, _stop, _direction);
            Console.WriteLine("Request from Floor {0} and goint to Floor {1}", _floor, _stop);
            Console.WriteLine("the elevator chosen is : {0}", bestOption.ID);
            runAll();
        }

        public void runAll()
        {
            foreach (Elevator elevator in elevatorList)
            {
                elevator.run();
            }
        }

        public void changeValue(int _elevator, List<int> _stopList, string _status, int _currentFloor, string _currentDirection)
        {
            elevatorList[_elevator].changeValue(_elevator, _stopList, _status, _currentFloor, _currentDirection);
        }
    }
}
