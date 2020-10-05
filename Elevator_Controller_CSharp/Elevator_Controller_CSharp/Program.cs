using System;
using System.Collections.Generic;

namespace Elevator_Controller_CSharp
{
    class Programa
    {


        static void Main(string[] args)
        {
            Battery battery = new Battery(4, 60, 6, 5);

            List<int> _stopList1 = new List<int>();
            _stopList1.Add(21);

            List<int> _stopList2 = new List<int>();
            _stopList2.Add(28);

            List<int> _stopList3 = new List<int>();
            _stopList3.Add(7);

            List<int> _stopList4 = new List<int>();
            _stopList4.Add(24);

            List<int> _stopList5 = new List<int>();
            _stopList5.Add(7);

            battery.changeValue(2, 0, _stopList1, "MOVING", 7, "up");
            battery.changeValue(2, 1, _stopList2, "MOVING", 23, "up");
            battery.changeValue(2, 2, _stopList3, "MOVING", 33, "down");
            battery.changeValue(2, 3, _stopList4, "MOVING", 40, "down");
            battery.changeValue(2, 4, _stopList5, "MOVING", 39, "down");

            battery.columnList[2].elevatorList[2].downBuffer.Add(30);

            battery.columnSelection(36, 7, "up");
        }
    }
}
