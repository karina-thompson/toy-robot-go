package simulator

import (
	"errors"
	"strconv"
	"strings"
)

const (
	place  = "place"
	report = "report"
	move   = "move"
)

var errInvalidCommand = errors.New("Invalid command, please try again")

// ProcessCommand executes a command if it is valid
func ProcessCommand(r *Robot, command string) error {
	command = strings.TrimSpace(strings.ToLower(command))
	switch {
	case validPlaceCommand(strings.Split(command, " ")):
		xPos, yPos, direction, err := placeParameters(strings.Split(command, " "))
		if err != nil {
			return err
		}
		return r.Place(xPos, yPos, direction)
	case command == report:
		return r.Report()
	case command == left || command == right:
		return r.Turn(command)
	case command == move:
		return r.Move()
	default:
		return errInvalidCommand
	}
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
		err = errInvalidPosition
	}
	return xPos, yPos, position[2], err
}
