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

func (r *ToyRobot) Place(xPos, yPos int, direction string) error {
	if position.Invalid(xPos, yPos) {
		return errors.New("Invalid position, robot could not be placed")
	}
	r.xPos, r.yPos, r.facing, r.onTable = xPos, yPos, direction, true
	return nil
}

func (r ToyRobot) Report() error {
	if !r.onTable {
		return errors.New("Robot is not currently on table")
	}
	fmt.Println(r.xPos, r.yPos, r.facing)
	return nil
}

func (r *ToyRobot) Move() error {
	if !r.onTable {
		return errors.New("Robot not on table, could not be moved")
	}
	xPos, yPos, err := position.Move(r.xPos, r.yPos, r.facing)
	r.xPos, r.yPos = xPos, yPos
	return err
}

func (r *ToyRobot) Turn(direction string) error {
	if !r.onTable {
		return errors.New("Robot not on table, could not be turned")
	}
	r.facing = compass.Turn(r.facing, direction)
	return nil
}
