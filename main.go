package main

import "github.com/karina-thompson/toy-robot-go/robot"

func main() {
	robot := robot.ToyRobot{}
	robot, _ = robot.Place(1, 2, "EAST")
	robot, _ = robot.Move()
	robot, _ = robot.Move()
	robot, _ = robot.Turn("LEFT")
	robot, _ = robot.Move()
	robot.Report()

	// var command string
	// fmt.Println("Toy Robot - please enter a command:")
	// fmt.Scan(&command)
	// fmt.Println(command)
}
