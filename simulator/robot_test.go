package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlace(t *testing.T) {
	testData := []struct {
		xPos, yPos int
		facing     string
		err        error
	}{
		{0, 0, "north", nil},
		{7, 3, "west", ErrInvalidPosition},
		{4, 3, "east", nil},
		{1, 3, "south", ErrInvalidPosition},
	}
	for _, testCase := range testData {
		r := Robot{}
		err := r.Place(testCase.xPos, testCase.yPos, testCase.facing)
		if err == nil {
			assert.Equal(t, r.xPos, testCase.xPos)
			assert.Equal(t, r.yPos, testCase.yPos)
			assert.Equal(t, r.facing, testCase.facing)
		}
		if err != nil {
			assert.Equal(t, r.yPos, 0)
			assert.Equal(t, r.xPos, 0)
			assert.Equal(t, err, testCase.err)
		}
	}
}

func TestMove(t *testing.T) {
	testData := []struct {
		xPos, yPos int
		facing     string
		x, y       int
	}{
		{0, 0, "north", 0, 1},
		{4, 4, "east", 4, 4},
		{1, 4, "west", 0, 4},
		{0, 0, "south", 0, 0},
	}
	for _, testCase := range testData {
		r := Robot{testCase.xPos, testCase.yPos, testCase.facing, true}
		err := r.Move()
		if err == nil {
			assert.Equal(t, r.xPos, testCase.x)
			assert.Equal(t, r.yPos, testCase.y)
		}
	}
}

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
		result := turn(test.facing, test.direction)
		if result != test.result {
			t.Errorf("Turn %v from %v was incorrect, got: %v, want: %v ",
				test.direction, test.facing, result, test.result)
		}
	}
}
