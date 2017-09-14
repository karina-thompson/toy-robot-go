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
		{7, 3, "west", errInvalidPosition},
		{4, 3, "east", nil},
		{1, 3, "south", errInvalidPosition},
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
		onTable    bool
		err        error
	}{
		{0, 0, "north", 0, 1, true, nil},
		{4, 4, "east", 4, 4, true, errInvalidMove},
		{1, 4, "west", 0, 4, true, nil},
		{0, 0, "south", 0, 0, true, errInvalidMove},
		{1, 4, "up", 1, 4, true, errInvalidMove},
		{0, 0, "", 0, 0, false, errNotOnTable},
	}
	for _, testCase := range testData {
		r := Robot{testCase.xPos, testCase.yPos, testCase.facing, testCase.onTable}
		err := r.Move()
		assert.Equal(t, r.xPos, testCase.x)
		assert.Equal(t, r.yPos, testCase.y)
		if err != nil {
			assert.Equal(t, err, testCase.err)
		}
	}
}

func TestTurn(t *testing.T) {
	testData := []struct {
		facing, direction, result string
		onTable                   bool
	}{
		{north, left, west, true},
		{west, right, north, true},
		{south, left, east, true},
		{east, right, south, true},
		{south, left, "", false},
	}
	for _, testCase := range testData {
		r := Robot{0, 0, testCase.facing, testCase.onTable}
		err := r.Turn(testCase.direction)
		if err != nil {
			assert.Equal(t, err, errNotOnTable)
		} else {
			assert.Equal(t, r.facing, testCase.result)
		}
	}
}

func TestReport(t *testing.T) {
	testData := []struct {
		xPos, yPos int
		facing     string
		onTable    bool
	}{
		{0, 0, north, true},
		{3, 4, south, true},
		{0, 0, "", false},
	}
	for _, testCase := range testData {
		r := Robot{testCase.xPos, testCase.yPos, testCase.facing, testCase.onTable}
		err := r.Report()
		if err != nil {
			assert.Equal(t, err, errNotOnTable)
		}
	}
}
