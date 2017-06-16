package table

const tableSize = 5

func InvalidPosition(xPos, yPos int) bool {
	return xPos > tableSize || xPos < 0 || yPos > tableSize || yPos < 0
}

func Move(xPos, yPos int, facing string) (int, int) {
	var x, y = xPos, yPos
	switch facing {
	case "NORTH":
		y++
	case "SOUTH":
		y--
	case "EAST":
		x++
	case "WEST":
		x--
	}
	if InvalidPosition(x, y) {
		return xPos, yPos
	}
	return x, y
}
