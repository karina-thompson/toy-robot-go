package simulator

import "errors"

const tableSize = 5

var ErrInvalidPosition = errors.New("Invalid position, robot could not be placed")
var errInvalidMove = errors.New("This move would make the robot fall off the table, robot could not be moved")
var errNotOnTable = errors.New("Robot is not currently on table")

func invalidPosition(xPos, yPos int) bool {
	return xPos >= tableSize || xPos < 0 || yPos >= tableSize || yPos < 0
}

func move(xPos, yPos int, facing string) (int, int, error) {
	x, y := xPos, yPos
	switch facing {
	case north:
		y++
	case south:
		y--
	case east:
		x++
	case west:
		x--
	}
	if invalidPosition(x, y) {
		return xPos, yPos, errInvalidMove
	}
	return x, y, nil
}
