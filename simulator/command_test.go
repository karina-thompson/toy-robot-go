package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessCommand(t *testing.T) {
	testData := []struct {
		name                   string
		robot                  *Robot
		command                string
		xPosResult, yPosResult int
		facingResult           string
		err                    error
	}{
		{"Valid move command", &Robot{0, 0, north, true}, "move", 0, 1, north, nil},
		{"Valid left command", &Robot{1, 3, south, true}, "left", 1, 3, east, nil},
		{"Valid right command", &Robot{4, 4, west, true}, "right", 4, 4, north, nil},
		{"Valid place command", &Robot{}, "place 3,3,east", 3, 3, east, nil},
		{"Valid report command", &Robot{2, 0, north, true}, "report", 2, 0, north, nil},
		{"Invalid place command with no coordinates", &Robot{}, "place", 0, 0, "", errInvalidCommand},
		{"Invalid place coordinates", &Robot{}, "place 3,east", 0, 0, "", errInvalidCommand},
		{"Invalid place coordinates", &Robot{}, "place x,x,x", 0, 0, "", errInvalidPosition},
	}
	for _, testCase := range testData {
		assert := assert.New(t)
		err := ProcessCommand(testCase.robot, testCase.command)
		assert.Equal(testCase.robot.xPos, testCase.xPosResult, testCase.name)
		assert.Equal(testCase.robot.yPos, testCase.yPosResult, testCase.name)
		assert.Equal(testCase.robot.facing, testCase.facingResult, testCase.name)
		assert.Equal(err, testCase.err, testCase.name)
	}
}
