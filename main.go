package main

import (
	"errors"
	"fmt"
)

type Robot struct {
	xPos, yPos int
	direction  string
	onTable    bool
}

func (r Robot) place(xPos int, yPos int, direction string) (Robot, error) {
	if xPos > 5 || xPos < 0 || yPos > 5 || yPos < 0 {
		return r, errors.New("Invalid position, robot could not be placed")
	}
	r.xPos, r.yPos, r.direction, r.onTable = xPos, yPos, direction, true
	return r, nil
}

func main() {
	robot := Robot{}
	robot, err := robot.place(9, 3, "NORTH")
	fmt.Println(robot, err)

	// var command string
	// fmt.Println("Toy Robot - please enter a command:")
	// fmt.Scan(&command)
	// fmt.Println(command)
}
