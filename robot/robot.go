package robot

import (
	"errors"
	"fmt"

	"github.com/karina-thompson/toy-robot-go/table"
)

type ToyRobot struct {
	xPos, yPos int
	Direction  string
	onTable    bool
}

func (r ToyRobot) Place(xPos, yPos int, direction string) (ToyRobot, error) {
	if table.InvalidPosition(xPos, yPos) {
		return r, errors.New("Invalid position, robot could not be placed")
	}
	r.xPos, r.yPos, r.Direction, r.onTable = xPos, yPos, direction, true
	return r, nil
}

func (r ToyRobot) Report() {
	fmt.Println(r.xPos, r.yPos, r.Direction)
}
