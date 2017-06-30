package position

import "errors"

const tableSize = 5

func Invalid(xPos, yPos int) bool {
	return xPos >= tableSize || xPos < 0 || yPos >= tableSize || yPos < 0
}

func Move(xPos, yPos int, facing string) (int, int, error) {
	x, y := xPos, yPos
	switch facing {
	case "north":
		y++
	case "south":
		y--
	case "east":
		x++
	case "west":
		x--
	}
	if Invalid(x, y) {
		return xPos, yPos, errors.New("This move would make the robot fall off the table, robot could not be moved")
	}
	return x, y, nil
}
