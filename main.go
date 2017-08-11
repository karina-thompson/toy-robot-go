package main

import (
	"os"

	"github.com/karina-thompson/toy-robot-go/simulator"
	"github.com/karina-thompson/toy-robot-go/toyrobot"
)

func main() {
	var inputFile string
	r := simulator.Robot{}
	if len(os.Args) == 2 {
		inputFile = os.Args[1]
	}
	toyrobot.Run(&r, inputFile)
}
