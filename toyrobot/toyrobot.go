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

func Run(inputFile string) {
	r := simulator.Robot{}
	if inputFile == "" {
		fmt.Println("Toy Robot - please enter a command:")
		stop := false
		for !stop {
			command := getCliCommand()
			if command == exit {
				stop = true
				return
			}
			if err := processCommand(&r, command); err != nil {
				fmt.Println(err)
			}
		}
	}
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return
	}
	commands := strings.Split(string(data), "\n")
	for _, c := range commands {
		if err = processCommand(&r, c); err != nil {
			fmt.Println(err)
		}
	}
}

func processCommand(r *simulator.Robot, command string) error {
	errInvalidCommand := errors.New("Invalid command, please try again")
	command = strings.ToLower(command)
	switch {
	case validPlaceCommand(strings.Split(command, " ")):
		xPos, yPos, direction, err := placeParameters(strings.Split(command, " "))
		if err != nil {
			return err
		}
		return r.Place(xPos, yPos, direction)
	case command == report:
		return r.Report()
	case command == simulator.Left || command == simulator.Right:
		return r.Turn(command)
	case command == move:
		return r.Move()
	default:
		return errInvalidCommand
	}
}

func getCliCommand() string {
	var command, position string
	fmt.Scanln(&command, &position)
	if position == "" {
		return command
	}
	return fmt.Sprintf("%v %v", command, position)
}

func validPlaceCommand(command []string) bool {
	if command[0] == place && len(command) == 2 {
		if len(strings.Split(command[1], ",")) == 3 {
			return true
		}
	}
	return false
}

func placeParameters(command []string) (int, int, string, error) {
	var err error
	position := strings.Split(command[1], ",")
	xPos, errX := strconv.Atoi(position[0])
	yPos, errY := strconv.Atoi(position[1])
	if errX != nil || errY != nil {
		err = simulator.ErrInvalidPosition
	}
	return xPos, yPos, position[2], err
}
