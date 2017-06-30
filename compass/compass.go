package compass

import "container/ring"

var directions = []string{"north", "east", "south", "west"}

func InvalidDirection(direction string) bool {
	for _, d := range directions {
		if direction == d {
			return true
		}
	}
	return false
}

func compass() (compass *ring.Ring) {
	compass = ring.New(len(directions))
	for i := 0; i < compass.Len(); i++ {
		compass.Value = directions[i]
		compass = compass.Next()
	}
	return
}

func Turn(facing, direction string) string {
	var result *ring.Ring
	compass := compass()

	for i := 0; i < compass.Len(); i++ {
		if compass.Value == facing {
			if direction == "right" {
				result = compass.Next()
			}
			if direction == "left" {
				result = compass.Prev()
			}
		}
		compass = compass.Next()
	}
	return (result.Value).(string)
}
