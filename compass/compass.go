package compass

import "container/ring"

func compass() (compass *ring.Ring) {
	directions := []string{"NORTH", "EAST", "SOUTH", "WEST"}
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
			if direction == "RIGHT" {
				result = compass.Next()
			}
			if direction == "LEFT" {
				result = compass.Prev()
			}
		}
		compass = compass.Next()
	}
	return (result.Value).(string)
}
