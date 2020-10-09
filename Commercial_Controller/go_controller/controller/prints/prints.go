package prints

import (
	"fmt"
	"go_controller/controller"
	"strconv"
	"time"
)

// don't look for comments in there, you won't be seeing any

// Colors ...
func Colors(color string, value string) string {
	if color == "red" {
		return string("\033[31m") + value + string("\033[0m")
	} else if color == "green" {
		return string("\033[32m") + value + string("\033[0m")
	} else if color == "yellow" {
		return string("\033[33m") + value + string("\033[0m")
	} else if color == "blue" {
		return string("\033[34m") + value + string("\033[0m")
	} else if color == "purple" {
		return string("\033[35m") + value + string("\033[0m")
	} else if color == "cyan" {
		return string("\033[36m") + value + string("\033[0m")
	} else if color == "white" {
		return string("\033[37m") + value + string("\033[0m")
	}
	return ""
}

// CreateState ...
func CreateState(_id string, _floor int, _status string, _stop int) {
	controller.Wait(2)
	TopBottomLine(_status)
	InnerArrowElevatorLine(_status)
	EmptyDoubleLine(_status)
	IDLine(_id, _status)
	FloorLine(_floor, _status)
	StatusLine(_status)
	StopLine(_stop, _status)
	EmptyDoubleLine(_status)
	InnerArrowLine(_status)
	TopBottomLine(_status)
}

// CreateArrival ...
func CreateArrival(_floor int) {
	size := CountStr(_floor)

	controller.Wait(2)
	TopBottomLine(size)
	FloorArivedLine(_floor)
	TopBottomLine(size)
}

// CreateRequest ...
func CreateRequest(_floor int, _stop int, _direction string) {
	controller.Wait(2)
	fmt.Println()
	TopBottomLine("2")
	InnerArrowAndRequestLine()
	EmptyDoubleLine("2")
	AtFloorLine(_floor)
	FloorRequestLine(_stop)
	DirectionLine(_direction)
	EmptyDoubleLine("2")
	InnerArrowLine("2")
	TopBottomLine("2")
}

// CreatePointing ...
func CreatePointing(_columnID string, _id []string, _points []string) {
	fmt.Println("")
	controller.Wait(2)

	ColumnSelectedLine(_columnID)
	fmt.Println("")

	for i := 0; i < len(_points); i++ {
		fmt.Println("		ELEVATOR " + Colors("red", _id[i]) + " - HAS " + Colors("yellow", _points[i]) + " pts")
		fmt.Println("")
	}

	BestElevatorLine(_id[0], _points[0])
	fmt.Println("")

	controller.Wait(4)
}

// BestElevatorLine ...
func BestElevatorLine(_id string, _points string) {
	fmt.Println("		THE BEST ELEVATOR IS ELEVATOR " + Colors("red", _id) + " WITH " + Colors("yellow", _points) + " pts")
}

// ColumnSelectedLine ...
func ColumnSelectedLine(_columnID string) {
	fmt.Println("		THE COLUMN SELECTED IS COLUMN " + Colors("red", _columnID))
}

// CountStr ...
func CountStr(n int) string {
	count := "0"

	if n >= 0 && n < 10 {
		count = "1"
	} else if n < 100 && n > 9 || n < 0 && n > -10 {
		count = "2"
	} else if n < 1000 && n > 99 || n < -10 && n > -100 {
		count = "3"
	} else if n < 10000 && n > 999 || n < -100 && n > -1000 {
		count = "4"
	}

	return count
}

// CountInt ...
func CountInt(n int) int {
	count := 0

	if n >= 0 && n < 10 {
		count = 1
	} else if n < 100 && n > 9 || n < 0 && n > -10 {
		count = 2
	} else if n < 1000 && n > 99 || n < -10 && n > -100 {
		count = 3
	} else if n < 10000 && n > 999 || n < -100 && n > -1000 {
		count = 4
	}

	return count
}

// TopBottomLine ...
func TopBottomLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		+------------------------------+")
	} else if _size == "MOVING" {
		fmt.Println("		+--------------------------------+")
	} else if _size == "1" {
		fmt.Println("		+-----------------------------------+")
	} else if _size == "2" {
		fmt.Println("		+------------------------------------+")
	} else if _size == "3" || _size == "MAINTENANCE" {
		fmt.Println("		+-------------------------------------+")
	} else if _size == "4" {
		fmt.Println("		+--------------------------------------+")
	}
}

// InnerArrowLine ...
func InnerArrowLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		| +--->                  <---+ |")
	} else if _size == "MOVING" {
		fmt.Println("		| +--->                    <---+ |")
	} else if _size == "1" {
		fmt.Println("		| +--->                       <---+ |")
	} else if _size == "2" {
		fmt.Println("		| +--->                        <---+ |")
	} else if _size == "3" || _size == "MAINTENANCE" {
		fmt.Println("		| +--->                         <---+ |")
	} else if _size == "4" {
		fmt.Println("		| +--->                          <---+ |")
	}
}

// InnerArrowElevatorLine ...
func InnerArrowElevatorLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		| +--->     Elevator     <---+ |")
	} else if _size == "MOVING" {
		fmt.Println("		| +--->      Elevator      <---+ |")
	} else if _size == "1" {
		fmt.Println("		| +--->        Elevator       <---+ |")
	} else if _size == "2" {
		fmt.Println("		| +--->        Elevator        <---+ |")
	} else if _size == "3" || _size == "MAINTENANCE" {
		fmt.Println("		| +--->         Elevator        <---+ |")
	} else if _size == "4" {
		fmt.Println("		| +--->         Elevator         <---+ |")
	}
}

// ElevatorLine ...
func ElevatorLine(_id int) {
	count := CountStr(_id)

	if count == "1" {
		fmt.Println("		  +--->      ELEVATOR", _id, "      <---+  ")
	} else if count == "2" {
		fmt.Println("		  +--->      ELEVATOR", _id, "     <---+  ")
	} else if count == "3" {
		fmt.Println("		  +--->     ELEVATOR", _id, "     <---+  ")
	} else if count == "4" {
		fmt.Println("		  +--->     ELEVATOR", _id, "    <---+  ")
	}
}

// InnerArrowAndRequestLine ...
func InnerArrowAndRequestLine() {
	fmt.Println("		| +--->         REQUEST        <---+ |") // "2"
}

// EmptyDoubleLine ...
func EmptyDoubleLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		| |                          | |")
	} else if _size == "MOVING" {
		fmt.Println("		| |                            | |")
	} else if _size == "2" {
		fmt.Println("		| |                                | |")
	} else if _size == "MAINTENANCE" || _size == "3" {
		fmt.Println("		| |                                 | |")
	} else if _size == "4" {
		fmt.Println("		| |                                  | |")
	}
}

// UpArrow ...
func UpArrow(_size string) {
	if _size == "1" {
		fmt.Println("		+---------------------+")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		| |        -        | |")
		fmt.Println("		| ▼      -/-\\-      ▼ |")
		fmt.Println("		|       /-/-\\-\\       |")
		fmt.Println("		|          |          |")
		fmt.Println("		| ▲       | |       ▲ |")
		fmt.Println("		| |                 | |")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		+---------------------+")
	} else if _size == "2" {
		fmt.Println("		+-------------------------+")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| ▼          -          ▼ |")
		fmt.Println("		|          -/-\\-          |")
		fmt.Println("		|        -/-/-\\-\\-        |")
		fmt.Println("		|           | |           |")
		fmt.Println("		|          | - |          |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		+-------------------------+")
	} else if _size == "3" {
		fmt.Println("		+----------------------------+")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| ▼            -           ▼ |")
		fmt.Println("		|            -/-\\-           |")
		fmt.Println("		|          -/-/-\\-\\-         |")
		fmt.Println("		|         /-/-/-\\-\\-\\        |")
		fmt.Println("		|            |   |           |")
		fmt.Println("		|             | |            |")
		fmt.Println("		| ▲          | - |         ▲ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		+----------------------------+")
	}
}

// DownArrow ...
func DownArrow(_size string) {
	if _size == "1" {
		fmt.Println("		+---------------------+")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		| |                 | |")
		fmt.Println("		| ▼       | |       ▼ |")
		fmt.Println("		|          |          |")
		fmt.Println("		|       \\-\\-/-/       |")
		fmt.Println("		| ▲      -\\-/-      ▲ |")
		fmt.Println("		| |        -        | |")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		+---------------------+")
	} else if _size == "2" {
		fmt.Println("		+-------------------------+")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| ▼                     ▼ |")
		fmt.Println("		|          | - |          |")
		fmt.Println("		|           | |           |")
		fmt.Println("		|        -\\-\\-/-/-        |")
		fmt.Println("		|          -\\-/-          |")
		fmt.Println("		| ▲          -          ▲ |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		+-------------------------+")
	} else if _size == "3" {
		fmt.Println("		+----------------------------+")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| ▼          | - |         ▼ |")
		fmt.Println("		|             | |            |")
		fmt.Println("		|            |   |           |")
		fmt.Println("		|         \\-\\-\\-/-/-/        |")
		fmt.Println("		|          -\\-\\-/-/-         |")
		fmt.Println("		|            -\\-/-           |")
		fmt.Println("		| ▲            -           ▲ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		+----------------------------+")
	}
}

// DoorOpen ...
func DoorOpen(_size string, graph ...string) {
	if _size == "1" || _size == "2" {
		DoorTopBottomLine(_size)
		DoorMiddleLine(_size)
		if _size == "2" {
			DoorMiddleLine(_size)
		}
		LeftArrowLine(_size)
		if _size == "2" {
			DoorMiddleLine(_size)
		}
		DoorMiddleLine(_size)
		OpenLine(_size)
		DoorMiddleLine(_size)
		DoorMiddleLine(_size)
		LeftArrowLine(_size)
		if _size == "2" {
			DoorMiddleLine(_size)
			DoorMiddleLine(_size)
			DoorMiddleLine(_size)
			DoorMiddleLine(_size)
		}
		DoorMiddleLine(_size)
		DoorTopBottomLine(_size)
	}

	if len(graph) == 0 {
		graph = append(graph, "█")
	}
	ProgressBar(120, 3, "OPENING", graph[0])
}

// DoorClose ...
func DoorClose(_size string, graph ...string) {
	if _size == "1" || _size == "2" {
		DoorTopBottomLine(_size)
		DoorMiddleLine(_size)
		if _size == "2" {
			DoorMiddleLine(_size)
		}
		RightArrowLine(_size)
		if _size == "2" {
			DoorMiddleLine(_size)
		}
		DoorMiddleLine(_size)
		ClosingLine(_size)
		DoorMiddleLine(_size)
		DoorMiddleLine(_size)
		RightArrowLine(_size)
		if _size == "2" {
			DoorMiddleLine(_size)
			DoorMiddleLine(_size)
			DoorMiddleLine(_size)
			DoorMiddleLine(_size)
		}
		DoorMiddleLine(_size)
		DoorTopBottomLine(_size)
	}

	if len(graph) == 0 {
		graph = append(graph, "█")
	}
	ProgressBar(120, 3, "CLOSING", graph[0])
}

// ProgressBar ...
func ProgressBar(_speed int, _timeSecond int, state string, graph string) {
	speed := _speed
	timeSecond := _timeSecond
	done := true
	go func() {
		controller.Wait(timeSecond)
		done = false
	}()

	fmt.Printf("		" + state + " ")
	for done {
		time.Sleep(time.Duration(speed) * time.Millisecond)
		fmt.Printf(graph)
	}
	fmt.Println()
}

// LeftArrowLine ...
func LeftArrowLine(_size string) {
	if _size == "1" {
		fmt.Println("		|  <<  <<  ||   |")
	} else if _size == "2" {
		fmt.Println("		|   <<   <<   <<   ||      |")
	}
}

// RightArrowLine ...
func RightArrowLine(_size string) {
	if _size == "1" {
		fmt.Println("		|  >>  >>  ||   |")
	} else if _size == "2" {
		fmt.Println("		|   >>   >>   >>   ||      |")
	}
}

// OpenLine ...
func OpenLine(_size string) {
	if _size == "1" {
		fmt.Println("		|   Open   ||   |")
	} else if _size == "2" {
		fmt.Println("		|     Opening      ||      |")
	}
}

// DoorMiddleLine ...
func DoorMiddleLine(_size string) {
	if _size == "1" {
		fmt.Println("		|   	   ||   |")
	} else if _size == "2" {
		fmt.Println("		|     	           ||      |")
	}
}

// DoorTopBottomLine ...
func DoorTopBottomLine(_size string) {
	if _size == "1" {
		fmt.Println("		+----------++---+")
	} else if _size == "2" {
		fmt.Println("		+------------------++------+")
	}
}

// ClosingLine ...
func ClosingLine(_size string) {
	if _size == "1" {
		fmt.Println("		|   Close  ||   |")
	} else if _size == "2" {
		fmt.Println("		|     Closing      ||      |")
	}
}

// FloorArivedLine ...
func FloorArivedLine(_floor int) {
	count := CountStr(_floor)

	if count == "1" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	} else if count == "2" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	} else if count == "3" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	} else if count == "4" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	}
}

// IDLine ...
func IDLine(_id string, _status string) {
	count := strconv.Itoa(len(_id))
	ID := Colors("red", _id)

	if _status == "IDLE" {
		if count == "1" {
			fmt.Println("		| |       ID: " + ID + "               | |")
		} else if count == "2" {
			fmt.Println("		| |       ID: " + ID + "             | |")
		} else if count == "3" {
			fmt.Println("		| |       ID: " + ID + "            | |")
		} else if count == "4" {
			fmt.Println("		| |       ID: " + ID + "           | |")
		}
	} else if _status == "MOVING" {
		if count == "1" {
			fmt.Println("		| |       ID: " + ID + "                | |")
		} else if count == "2" {
			fmt.Println("		| |       ID: " + ID + "               | |")
		} else if count == "3" {
			fmt.Println("		| |       ID: " + ID + "              | |")
		} else if count == "4" {
			fmt.Println("		| |       ID: " + ID + "             | |")
		}
	} else if _status == "MAINTENANCE" {
		if count == "1" {
			fmt.Println("		| |       ID: " + ID + "                     | |")
		} else if count == "2" {
			fmt.Println("		| |       ID: " + ID + "                    | |")
		} else if count == "3" {
			fmt.Println("		| |       ID: " + ID + "                   | |")
		} else if count == "4" {
			fmt.Println("		| |       ID: " + ID + "                  | |")
		}
	}
}

// DirectionLine ...
func DirectionLine(_direction string) {
	if _direction == "Up" {
		fmt.Println("		| |        DIRECTION:", "UP", "          | |")
	} else if _direction == "Down" {
		fmt.Println("		| |        DIRECTION:", "DOWN", "        | |")
	} else if _direction == "Stop" {
		fmt.Println("		| |        DIRECTION:", "STOP", "        | |")
	}
}

// AtFloorLine ...
func AtFloorLine(_atFloor int) {
	count := CountStr(_atFloor)

	if count == "1" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "            | |")
	} else if count == "2" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "           | |")
	} else if count == "3" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "          | |")
	} else if count == "4" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "         | |")
	}
}

// FloorRequestLine ...
func FloorRequestLine(_requestFloor int) {
	count := CountStr(_requestFloor)

	if count == "1" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "     | |")
	} else if count == "2" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "    | |")
	} else if count == "3" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "   | |")
	} else if count == "4" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "  | |")
	}
}

// FloorLine ...
func FloorLine(_floor int, _status string) {
	count := CountStr(_floor)

	if _status == "IDLE" {
		if count == "1" {
			fmt.Println("		| |       Floor:", _floor, "          | |")
		} else if count == "2" {
			fmt.Println("		| |       Floor:", _floor, "         | |")
		} else if count == "3" {
			fmt.Println("		| |       Floor:", _floor, "        | |")
		} else if count == "4" {
			fmt.Println("		| |       Floor:", _floor, "       | |")
		}
	} else if _status == "MOVING" {
		if count == "1" {
			fmt.Println("		| |       Floor:", _floor, "            | |")
		} else if count == "2" {
			fmt.Println("		| |       Floor:", _floor, "           | |")
		} else if count == "3" {
			fmt.Println("		| |       Floor:", _floor, "          | |")
		} else if count == "4" {
			fmt.Println("		| |       Floor:", _floor, "         | |")
		}
	} else if _status == "MAINTENANCE" {
		if count == "1" {
			fmt.Println("		| |       Floor:", _floor, "                 | |")
		} else if count == "2" {
			fmt.Println("		| |       Floor:", _floor, "                | |")
		} else if count == "3" {
			fmt.Println("		| |       Floor:", _floor, "               | |")
		} else if count == "4" {
			fmt.Println("		| |       Floor:", _floor, "              | |")
		}
	}
}

// StopLine ...
func StopLine(_stop int, _status string) {
	count := CountStr(_stop)

	if _status == "IDLE" {
		if count == "1" {
			fmt.Println("		| |       Next-Stop:", _stop, "      | |")
		} else if count == "2" {
			fmt.Println("		| |       Next-Stop:", _stop, "     | |")
		} else if count == "3" {
			fmt.Println("		| |       Next-Stop:", _stop, "    | |")
		} else if count == "4" {
			fmt.Println("		| |       Next-Stop:", _stop, "   | |")
		}
	} else if _status == "MOVING" {
		if count == "1" {
			fmt.Println("		| |       Next-Stop:", _stop, "        | |")
		} else if count == "2" {
			fmt.Println("		| |       Next-Stop:", _stop, "       | |")
		} else if count == "3" {
			fmt.Println("		| |       Next-Stop:", _stop, "      | |")
		} else if count == "4" {
			fmt.Println("		| |       Next-Stop:", _stop, "     | |")
		}
	} else if _status == "MAINTENANCE" {
		if count == "1" {
			fmt.Println("		| |       Next-Stop:", _stop, "             | |")
		} else if count == "2" {
			fmt.Println("		| |       Next-Stop:", _stop, "            | |")
		} else if count == "3" {
			fmt.Println("		| |       Next-Stop:", _stop, "           | |")
		} else if count == "4" {
			fmt.Println("		| |       Next-Stop:", _stop, "          | |")
		}
	}
}

// StatusLine ...
func StatusLine(_status string) {
	if _status == "IDLE" {
		fmt.Println("		| |       Status:", "IDLE", "      | |")
	} else if _status == "MOVING" {
		fmt.Println("		| |       Status:", "MOVING", "      | |")
	} else if _status == "MAINTENANCE" {
		fmt.Println("		| |       Status:", "MAINTENANCE", "      | |")
	}
}

// LineLine ...
func LineLine(i int) {
	if i == 1 {
		fmt.Println("		|")
	} else if i == 2 {
		fmt.Println("		|   |")
	} else if i == 3 {
		fmt.Println("		|       |")
	} else if i == 4 {
		fmt.Println("		|           |")
	} else if i == 5 {
		fmt.Println("		|               |")
	} else if i == 6 {
		fmt.Println("		|                   |")
	}
}
