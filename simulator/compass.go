package simulator

import "container/ring"

const (
	Left  = "left"
	Right = "right"
	north = "north"
	east  = "east"
	south = "south"
	west  = "west"
)

var directions = []string{north, east, south, west}

func compass() *ring.Ring {
	compass := ring.New(len(directions))
	for i := 0; i < compass.Len(); i++ {
		compass.Value = directions[i]
		compass = compass.Next()
	}
	return compass
}

func isValidDirection(direction string) bool {
	for _, d := range directions {
		if direction == d {
			return true
		}
	}
	return false
}

func turn(facing, direction string) string {
	var result *ring.Ring
	compass := compass()
	for i := 0; i < compass.Len(); i++ {
		if compass.Value == facing {
			if direction == Right {
				result = compass.Next()
			}
			if direction == Left {
				result = compass.Prev()
			}
		}
		compass = compass.Next()
	}
	return (result.Value).(string)
}
