package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/karina-thompson/toy-robot-go/robot"
)

func Run() {
	fmt.Println("Toy Robot - please enter a command:")
	r := robot.ToyRobot{}
	acceptCommands := true
	errInvalidCommand := errors.New("Invalid command, please try again")

	for acceptCommands {
		var command, parameters string
		var err error
		fmt.Scanln(&command, &parameters)
		command = strings.ToLower(command)

		switch {
		case invalidExtraInput(command, parameters):
			err = errInvalidCommand
			break
		case validPlaceCommand(command, parameters):
			position := strings.Split(parameters, ",")
			xPos, errX := strconv.Atoi(position[0])
			yPos, errY := strconv.Atoi(position[1])
			if errX != nil || errY != nil {
				err = robot.ErrInvalidPosition
				break
			}
			err = r.Place(xPos, yPos, position[2])
		case command == "report":
			err = r.Report()
		case command == "left" || command == "right":
			err = r.Turn(command)
		case command == "move":
			err = r.Move()
		case command == "exit":
			acceptCommands = false
		default:
			err = errInvalidCommand
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}

func invalidExtraInput(command, parameters string) bool {
	return command != "place" && len(parameters) != 0
}

func validPlaceCommand(command, parameters string) bool {
	return command == "place" && len(parameters) != 0 && len(strings.Split(parameters, ",")) == 3
}
