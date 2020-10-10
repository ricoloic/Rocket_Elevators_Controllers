package main

import (
	"go_controller/controller/preset"
)

// for no animation/wait time look in controller/controller.go

// if the terminal doesn't support colored text,
// could you change it to one that does because you'll have something like this "\033[31m]"
// where there's supposed to be colored text

func main() {
	preset.CodeBoxx(3) // calling a scenario to take place (1 to 4)
}
