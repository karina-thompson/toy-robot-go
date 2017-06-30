package position

import "testing"

func TestInvalidPosition(t *testing.T) {
	positionTests := []struct {
		xPos, yPos int
		result     bool
	}{
		{0, 0, false},
		{3, 4, false},
		{-1, 0, true},
		{2, 7, true},
	}
	for _, test := range positionTests {
		result := InvalidPosition(test.xPos, test.yPos)
		if result != test.result {
			t.Errorf("Result for invalid position test for x: %v and y: %v was incorrect, got: %v, want: %v ",
				test.xPos, test.yPos, result, test.result)
		}
	}
}

func TestMove(t *testing.T) {
	moveTests := []struct {
		xPos, yPos int
		facing     string
		x, y       int
	}{
		{0, 0, "north", 0, 1},
		{4, 4, "east", 4, 4},
		{1, 4, "west", 0, 4},
		{0, 0, "south", 0, 0},
	}
	for _, test := range moveTests {
		x, y, _ := Move(test.xPos, test.yPos, test.facing)
		if x != test.x || y != test.y {
			t.Errorf("Expected (%v, %v), got (%v, %v)", test.x, test.y, x, y)
		}
	}
}
