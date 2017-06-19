package main

import (
	"fmt"

	"github.com/karina-thompson/toy-robot-go/robot"
)

func main() {
	r := robot.ToyRobot{}
	_, err := r.Move()
	fmt.Println(err)
	r.Place(0, 0, "NORTH")
	r.Move()
	r.Report()

	// var command string
	// fmt.Println("Toy Robot - please enter a command:")
	// fmt.Scan(&command)
	// fmt.Println(command)
}
