package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidDirection(t *testing.T) {
	testCases := []struct {
		name      string
		direction string
		result    bool
	}{
		{"Valid direction", "north", true},
		{"Valid direction", "west", true},
		{"Invalid direction", "up", false},
		{"Invalid direction", "down", false},
	}
	for _, testCase := range testCases {
		result := isValidDirection(testCase.direction)
		assert.Equal(t, result, testCase.result, testCase.name)
	}
}
