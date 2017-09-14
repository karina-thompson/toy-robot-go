package simulator

import "fmt"

type Robot struct {
	xPos, yPos int
	facing     string
	onTable    bool
}

// Place places the Robot at the supplied position
func (r *Robot) Place(xPos, yPos int, direction string) error {
	if invalidPosition(xPos, yPos) || !isValidDirection(direction) {
		return errInvalidPosition
	}
	r.xPos, r.yPos, r.facing, r.onTable = xPos, yPos, direction, true
	return nil
}

// Report prints the Robot's current position to stdout
func (r Robot) Report() error {
	err := r.checkOnTable()
	fmt.Println(r.xPos, r.yPos, r.facing)
	return err
}

// Move moves the Robot 1 space in the direction it is facing
func (r *Robot) Move() error {
	if err := r.checkOnTable(); err != nil {
		return err
	}
	xPos, yPos, err := moveRobot(r.xPos, r.yPos, r.facing)
	r.xPos, r.yPos = xPos, yPos
	return err
}

// Turn moves the Robot 90 degrees in a left or right direction
func (r *Robot) Turn(direction string) error {
	err := r.checkOnTable()
	r.facing = turn(r.facing, direction)
	return err
}

func (r Robot) checkOnTable() error {
	if r.onTable {
		return nil
	}
	return errNotOnTable
}
