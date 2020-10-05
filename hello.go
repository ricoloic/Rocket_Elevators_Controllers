package main

import (
	"fmt"
)

func main() {
	var thisArray []int

	fmt.Println(thisArray)
	thisArray = append(thisArray, 3)

	fmt.Println(thisArray)
	var thatArray []int = []int{2, 3, 5}
	fmt.Println(len(thatArray))

	thisArray = append(thisArray, 1)
	fmt.Println(thisArray)
}
