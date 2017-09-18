package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidPosition(t *testing.T) {
	testCases := []struct {
		name       string
		xPos, yPos int
		result     bool
	}{
		{"Valid position", 0, 0, false},
		{"Valid position", 3, 4, false},
		{"Invalid position", -1, 0, true},
		{"Invalid position", 2, 7, true},
	}
	for _, testCase := range testCases {
		result := invalidPosition(testCase.xPos, testCase.yPos)
		assert.Equal(t, result, testCase.result, testCase.name)
	}
}
