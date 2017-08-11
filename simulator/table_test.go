package simulator

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
		result := invalidPosition(test.xPos, test.yPos)
		if result != test.result {
			t.Errorf("Result for invalid position test for x: %v and y: %v was incorrect, got: %v, want: %v ",
				test.xPos, test.yPos, result, test.result)
		}
	}
}
