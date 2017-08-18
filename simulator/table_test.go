package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidPosition(t *testing.T) {
	testCases := []struct {
		xPos, yPos int
		result     bool
	}{
		{0, 0, false},
		{3, 4, false},
		{-1, 0, true},
		{2, 7, true},
	}
	for _, testCase := range testCases {
		result := invalidPosition(testCase.xPos, testCase.yPos)
		assert.Equal(t, result, testCase.result)
	}
}
