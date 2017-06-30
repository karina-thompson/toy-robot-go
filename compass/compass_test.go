package compass

import "testing"

func TestTurn(t *testing.T) {
	turnTests := []struct {
		facing, direction, result string
	}{
		{"north", "left", "west"},
		{"west", "right", "north"},
		{"south", "left", "east"},
		{"east", "right", "south"},
	}
	for _, test := range turnTests {
		result := Turn(test.facing, test.direction)
		if result != test.result {
			t.Errorf("Turn %v from %v was incorrect, got: %v, want: %v ",
				test.direction, test.facing, result, test.result)
		}
	}
}
