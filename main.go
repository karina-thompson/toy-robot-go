package main

import (
	"os"

	"github.com/karina-thompson/toy-robot-go/toyrobot"
)

func main() {
	var inputFile string
	if len(os.Args) == 2 {
		inputFile = os.Args[1]
	}
	toyrobot.Run(inputFile)
}
