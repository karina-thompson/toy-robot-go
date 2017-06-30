package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/karina-thompson/toy-robot-go/robot"
)

func main() {
	fmt.Println("Toy Robot - please enter a command:")
	acceptCommands := true
	r := robot.ToyRobot{}

	for acceptCommands {
		var command, parameters string
		var err error
		fmt.Scanln(&command, &parameters)

		switch strings.ToLower(command) {
		case "place":
			if len(parameters) == 0 {
				err = errors.New("The PLACE command requires a position, please try again")
				break
			}
			position := strings.Split(parameters, ",")
			if len(position) != 3 {
				err = errors.New("Invalid position, robot could not be placed")
				break
			}
			xPos, errX := strconv.Atoi(position[0])
			yPos, errY := strconv.Atoi(position[1])
			if errX != nil || errY != nil {
				err = errors.New("Invalid position, robot could not be placed")
				break
			}
			err = r.Place(xPos, yPos, position[2])
		case "report":
			err = r.Report()
		case "left":
			err = r.Turn("left")
		case "right":
			err = r.Turn("right")
		case "move":
			err = r.Move()
		case "exit":
			acceptCommands = false
		default:
			err = errors.New("Invalid command, please try again")
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
