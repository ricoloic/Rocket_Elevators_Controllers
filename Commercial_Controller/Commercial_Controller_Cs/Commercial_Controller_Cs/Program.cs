using System;
using System.Collections.Generic;

namespace Elevator_Controller_CSharp
{
    class Program
    {
        static void Main(string[] args)
        {
            codeboxx(4);

            void codeboxx(int n)
            {
                Battery battery = new Battery(4, 60, 6, 5);

                // ||=========> BASEMENT COLUMN #1 <=========||
                battery.changeValue(0, 0, "IDLE", -4, new List<int>(), "stop");
                battery.changeValue(0, 1, "IDLE", 1, new List<int>(), "stop");
                battery.changeValue(0, 2, "MOVING", -3, new List<int>() { -5 }, "down");
                battery.changeValue(0, 3, "MOVING", -6, new List<int>() { 1 }, "up");
                battery.changeValue(0, 4, "MOVING", -1, new List<int>() { -6 }, "down");

                // ||=========> FLOOR COLUMN #2 <=========||
                battery.changeValue(1, 0, "MOVING", 20, new List<int>() { 5 }, "down");
                battery.changeValue(1, 1, "MOVING", 3, new List<int>() { 15 }, "up");
                battery.changeValue(1, 2, "MOVING", 13, new List<int>() { 1 }, "down");
                battery.changeValue(1, 3, "MOVING", 15, new List<int>() { 2 }, "down");
                battery.changeValue(1, 4, "MOVING", 6, new List<int>() { 1 }, "down");

                // ||=========> FLOOR COLUMN #3 <=========||
                battery.changeValue(2, 0, "MOVING", 1, new List<int>() { 21 }, "up");
                battery.changeValue(2, 1, "MOVING", 23, new List<int>() { 28 }, "up");
                battery.changeValue(2, 2, "MOVING", 33, new List<int>() { 1 }, "down");
                battery.changeValue(2, 3, "MOVING", 40, new List<int>() { 24 }, "down");
                battery.changeValue(2, 4, "MOVING", 39, new List<int>() { 1 }, "down");

                // ||=========> FLOOR COLUMN #4 <=========||
                battery.changeValue(3, 0, "MOVING", 58, new List<int>() { 1 }, "down");
                battery.changeValue(3, 1, "MOVING", 50, new List<int>() { 60 }, "up");
                battery.changeValue(3, 2, "MOVING", 46, new List<int>() { 58 }, "up");
                battery.changeValue(3, 3, "MOVING", 1, new List<int>() { 54 }, "up");
                battery.changeValue(3, 4, "MOVING", 60, new List<int>() { 1 }, "down");

                if (n == 1)
                {
                    battery.columnSelection(1, 20, "up");
                }
                
                else if (n == 2)
                {
                    battery.columnSelection(1, 36, "up");
                }

                else if (n == 3)
                {
                    battery.columnSelection(54, 1, "down");
                }

                else if (n == 4)
                {
                    battery.columnSelection(-3, 1, "up");
                }
            }
        }
    }
}