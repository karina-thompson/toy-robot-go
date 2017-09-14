package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/karina-thompson/toy-robot-go/simulator"
)

func main() {
	var input io.Reader
	var err error
	if len(os.Args) == 2 {
		input, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Invalid input file path")
			return
		}
	} else {
		fmt.Println("Toy Robot - please enter a command:")
		input = os.Stdin
	}

	scanner := bufio.NewScanner(input)
	r := simulator.Robot{}

	for scanner.Scan() {
		command := scanner.Text()
		if strings.ToLower(command) == "exit" {
			return
		}
		if err = simulator.ProcessCommand(&r, command); err != nil {
			fmt.Println(err)
		}
	}
}
