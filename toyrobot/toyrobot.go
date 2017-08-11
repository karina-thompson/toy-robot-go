package toyrobot

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/karina-thompson/toy-robot-go/simulator"
)

const (
	place  = "place"
	report = "report"
	move   = "move"
	exit   = "exit"
)

func Run(r *simulator.Robot, inputFile string) {
	if inputFile == "" {
		fmt.Println("Toy Robot - please enter a command:")
		for {
			command, parameters := getCliCommand()
			if command == exit {
				return
			}
			err := processCommand(r, command, parameters)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	data, err := ioutil.ReadFile(inputFile)
	if err == nil {
		commands := strings.Split(string(data), "\n")
		for _, c := range commands {
			command, parameters := parseInput(c)
			err = processCommand(r, command, parameters)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func parseInput(input string) (string, string) {
	splitInput := strings.Split(input, " ")
	var parameters string
	command := splitInput[0]
	if len(splitInput) > 1 {
		parameters = splitInput[1]
	}
	return command, parameters
}

func processCommand(r *simulator.Robot, command, parameters string) error {
	var err error
	errInvalidCommand := errors.New("Invalid command, please try again")
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
			err = simulator.ErrInvalidPosition
			break
		}
		err = r.Place(xPos, yPos, strings.ToLower(position[2]))
	case command == report:
		err = r.Report()
	case command == simulator.Left || command == simulator.Right:
		err = r.Turn(command)
	case command == move:
		err = r.Move()
	default:
		err = errInvalidCommand
	}
	return err
}

func getCliCommand() (string, string) {
	var command, parameters string
	fmt.Scanln(&command, &parameters)
	return command, parameters
}

func invalidExtraInput(command, parameters string) bool {
	return command != place && len(parameters) != 0
}

func validPlaceCommand(command, parameters string) bool {
	return command == place && len(parameters) != 0 && len(strings.Split(parameters, ",")) == 3
}