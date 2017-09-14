package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessCommand(t *testing.T) {
	testData := []struct {
		robot      *Robot
		command    string
		xPos, yPos int
		facing     string
		err        error
	}{
		{&Robot{0, 0, north, true}, "move", 0, 1, north, nil},
		{&Robot{1, 3, south, true}, "left", 1, 3, east, nil},
		{&Robot{4, 4, west, true}, "right", 4, 4, north, nil},
		{&Robot{}, "place 3,3,east", 3, 3, east, nil},
		{&Robot{}, "place 3,east", 0, 0, "", errInvalidCommand},
		{&Robot{}, "place", 0, 0, "", errInvalidCommand},
		{&Robot{}, "place x,x,x", 0, 0, "", errInvalidPosition},
		{&Robot{2, 0, north, true}, "report", 2, 0, north, nil},
	}
	for _, testCase := range testData {
		err := ProcessCommand(testCase.robot, testCase.command)
		assert.Equal(t, testCase.robot.xPos, testCase.xPos)
		assert.Equal(t, testCase.robot.yPos, testCase.yPos)
		assert.Equal(t, testCase.robot.facing, testCase.facing)
		assert.Equal(t, err, testCase.err)
	}
}
