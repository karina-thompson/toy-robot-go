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
	if len(os.Args) == 2 {
		input, _ = os.Open(os.Args[1])
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
		if err := simulator.ProcessCommand(&r, command); err != nil {
			fmt.Println(err)
		}
	}
}
