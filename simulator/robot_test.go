package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlace(t *testing.T) {
	testData := []struct {
		name       string
		xPos, yPos int
		facing     string
		err        error
	}{
		{"Valid place coordinates", 0, 0, "north", nil},
		{"Valid place coordinates", 4, 3, "east", nil},
		{"Invalid place coordinates", 7, 3, "west", errInvalidPosition},
		{"Invalid place coordinates", 1, 3, "south", errInvalidPosition},
	}
	for _, testCase := range testData {
		assert := assert.New(t)
		r := Robot{}
		err := r.Place(testCase.xPos, testCase.yPos, testCase.facing)
		if err == nil {
			assert.Equal(r.xPos, testCase.xPos, testCase.name)
			assert.Equal(r.yPos, testCase.yPos, testCase.name)
			assert.Equal(r.facing, testCase.facing, testCase.name)
		}
		if err != nil {
			assert.Equal(r.yPos, 0, testCase.name)
			assert.Equal(r.xPos, 0, testCase.name)
			assert.Equal(err, testCase.err, testCase.name)
		}
	}
}

func TestMove(t *testing.T) {
	testData := []struct {
		name       string
		xPos, yPos int
		facing     string
		x, y       int
		onTable    bool
		err        error
	}{
		{"Valid move", 0, 0, "north", 0, 1, true, nil},
		{"Valid move", 1, 4, "west", 0, 4, true, nil},
		{"Move would make robot fall off table", 4, 4, "east", 4, 4, true, errInvalidMove},
		{"Move would make robot fall off table", 0, 0, "south", 0, 0, true, errInvalidMove},
		{"Invalid direction", 1, 4, "up", 1, 4, true, errInvalidMove},
		{"Robot not on table", 0, 0, "", 0, 0, false, errNotOnTable},
	}
	for _, testCase := range testData {
		assert := assert.New(t)
		r := Robot{testCase.xPos, testCase.yPos, testCase.facing, testCase.onTable}
		err := r.Move()
		assert.Equal(r.xPos, testCase.x, testCase.name)
		assert.Equal(r.yPos, testCase.y, testCase.name)
		if err != nil {
			assert.Equal(err, testCase.err, testCase.name)
		}
	}
}

func TestTurn(t *testing.T) {
	testData := []struct {
		name, facing, direction, result string
		onTable                         bool
	}{
		{"Valid turn", north, left, west, true},
		{"Valid turn", west, right, north, true},
		{"Valid turn", south, left, east, true},
		{"Valid turn", east, right, south, true},
		{"Robot not on table", south, left, "", false},
	}
	for _, testCase := range testData {
		assert := assert.New(t)
		r := Robot{0, 0, testCase.facing, testCase.onTable}
		err := r.Turn(testCase.direction)
		if err != nil {
			assert.Equal(err, errNotOnTable, testCase.name)
		} else {
			assert.Equal(r.facing, testCase.result, testCase.name)
		}
	}
}

func TestReport(t *testing.T) {
	testData := []struct {
		name       string
		xPos, yPos int
		facing     string
		onTable    bool
	}{
		{"Robot on table", 0, 0, north, true},
		{"Robot on table", 3, 4, south, true},
		{"Robot not on table", 0, 0, "", false},
	}
	for _, testCase := range testData {
		r := Robot{testCase.xPos, testCase.yPos, testCase.facing, testCase.onTable}
		err := r.Report()
		if err != nil {
			assert.Equal(t, err, errNotOnTable, testCase.name)
		}
	}
}
