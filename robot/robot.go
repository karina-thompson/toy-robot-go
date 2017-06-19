package robot

import (
	"errors"
	"fmt"

	"github.com/karina-thompson/toy-robot-go/compass"
	"github.com/karina-thompson/toy-robot-go/table"
)

type ToyRobot struct {
	xPos, yPos int
	facing     string
	onTable    bool
}

func (r *ToyRobot) Place(xPos, yPos int, direction string) (*ToyRobot, error) {
	if table.InvalidPosition(xPos, yPos) {
		return r, errors.New("Invalid position, robot could not be placed")
	}
	r.xPos, r.yPos, r.facing, r.onTable = xPos, yPos, direction, true
	return r, nil
}

func (r ToyRobot) Report() {
	fmt.Println(r.xPos, r.yPos, r.facing)
}

func (r *ToyRobot) Move() (*ToyRobot, error) {
	if !r.onTable {
		return r, errors.New("Robot not on table, could not be moved")
	}
	var err error
	r.xPos, r.yPos, err = table.Move(r.xPos, r.yPos, r.facing)
	return r, err
}

func (r *ToyRobot) Turn(direction string) (*ToyRobot, error) {
	if !r.onTable {
		return r, errors.New("Robot not on table, could not be turned")
	}
	r.facing = compass.Turn(r.facing, direction)
	return r, nil
}
