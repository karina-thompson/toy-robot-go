package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/karina-thompson/toy-robot-go/robot"
)

func main() {
	fmt.Println("Toy Robot - please enter a command:")
	var exit bool
	r := robot.ToyRobot{}

	for !exit {
		var command, parameters string
		var err error
		fmt.Scanln(&command, &parameters)

		switch command {
		case "PLACE":
			position := strings.Split(parameters, ",")
			xPos, _ := strconv.Atoi(position[0])
			yPos, _ := strconv.Atoi(position[1])
			_, err = r.Place(xPos, yPos, position[2])
		case "REPORT":
			r.Report()
		case "LEFT":
			_, err = r.Turn("LEFT")
		case "RIGHT":
			_, err = r.Turn("RIGHT")
		case "MOVE":
			_, err = r.Move()
		case "EXIT":
			exit = true
		default:
			fmt.Println("Invalid command,please try again")
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
