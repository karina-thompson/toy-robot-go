package simulator

import "fmt"

type Robot struct {
	xPos, yPos int
	facing     string
	onTable    bool
}

func (r *Robot) Place(xPos, yPos int, direction string) error {
	if invalidPosition(xPos, yPos) || !isValidDirection(direction) {
		return errInvalidPosition
	}
	r.xPos, r.yPos, r.facing, r.onTable = xPos, yPos, direction, true
	return nil
}

func (r Robot) Report() error {
	err := r.checkOnTable()
	fmt.Println(r.xPos, r.yPos, r.facing)
	return err
}

func (r *Robot) Move() error {
	if err := r.checkOnTable(); err != nil {
		return err
	}
	xPos, yPos, err := moveRobot(r.xPos, r.yPos, r.facing)
	r.xPos, r.yPos = xPos, yPos
	return err
}

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
