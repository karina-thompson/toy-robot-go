package main

import "github.com/karina-thompson/toy-robot-go/robot"

func main() {
	robot := robot.ToyRobot{}
	robot.Place(0, 0, "NORTH")
	robot.Move()
	robot.Move()
	robot.Turn("LEFT")
	robot.Move()
	robot.Report()

	// var command string
	// fmt.Println("Toy Robot - please enter a command:")
	// fmt.Scan(&command)
	// fmt.Println(command)
}
