package main

import (
	"github.com/karina-thompson/toy-robot-go/compass"
	"github.com/karina-thompson/toy-robot-go/robot"
)

func main() {
	robot := robot.ToyRobot{}
	robot, _ = robot.Place(2, 2, "NORTH")
	robot.Direction = compass.Turn(robot.Direction, "LEFT")
	robot.Report()

	// var command string
	// fmt.Println("Toy Robot - please enter a command:")
	// fmt.Scan(&command)
	// fmt.Println(command)
}
