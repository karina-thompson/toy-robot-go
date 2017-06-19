package compass

import "testing"

func TestTurn(t *testing.T) {
	turnTests := []struct {
		facing    string
		direction string
		result    string
	}{
		{"NORTH", "LEFT", "WEST"},
		{"WEST", "RIGHT", "NORTH"},
		{"SOUTH", "LEFT", "EAST"},
		{"EAST", "RIGHT", "SOUTH"},
		{"da", "do", ""},
	}
	for _, test := range turnTests {
		result := Turn(test.facing, test.direction)
		if result != test.result {
			t.Errorf("Turn %v from %v was incorrect, got: %v, want: %v ",
				test.direction, test.facing, result, test.result)
		}
	}
}
