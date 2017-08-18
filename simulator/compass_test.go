package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidDirection(t *testing.T) {
	testCases := []struct {
		direction string
		result    bool
	}{
		{"north", true},
		{"up", false},
		{"west", true},
		{"down", false},
	}
	for _, testCase := range testCases {
		result := isValidDirection(testCase.direction)
		assert.Equal(t, result, testCase.result)
	}
}
