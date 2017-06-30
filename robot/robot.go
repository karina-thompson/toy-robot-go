package robot

import (
	"errors"
	"fmt"

	"github.com/karina-thompson/toy-robot-go/compass"
	"github.com/karina-thompson/toy-robot-go/position"
)

type ToyRobot struct {
	xPos, yPos int
	facing     string
	onTable    bool
}

var ErrInvalidPosition = errors.New("Invalid position, robot could not be placed")
var errNotOnTable = errors.New("Robot is not currently on table")

func (r *ToyRobot) Place(xPos, yPos int, direction string) error {
	if position.InvalidPosition(xPos, yPos) || compass.InvalidDirection(direction) {
		return ErrInvalidPosition
	}
	r.xPos, r.yPos, r.facing, r.onTable = xPos, yPos, direction, true
	return nil
}

func (r ToyRobot) checkOnTable() error {
	if r.onTable {
		return nil
	}
	return errNotOnTable
}

func (r ToyRobot) Report() error {
	err := r.checkOnTable()
	fmt.Println(r.xPos, r.yPos, r.facing)
	return err
}

func (r *ToyRobot) Move() error {
	err := r.checkOnTable()
	if err != nil {
		return err
	}
	xPos, yPos, err := position.Move(r.xPos, r.yPos, r.facing)
	r.xPos, r.yPos = xPos, yPos
	return err
}

func (r *ToyRobot) Turn(direction string) error {
	err := r.checkOnTable()
	r.facing = compass.Turn(r.facing, direction)
	return err
}
