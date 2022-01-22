package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const exampleInput = "./example.input"
const actualInput = "./actual.input"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	Forward string = "forward"
	Down           = "up"
	Up             = "down"
)

type Command struct {
	cmd string
	val int
}

func parseCommand(command string) Command {
	commandFields := strings.Fields(command)
	instruction := commandFields[0]
	value, err := strconv.Atoi(commandFields[1])
	check(err)
	return Command{instruction, value}
}

func main() {
	file, err := os.Open(actualInput)
	check(err)

	scanner := bufio.NewScanner(file)
	var allCommands []Command
	for scanner.Scan() {
		line := scanner.Text()
		allCommands = append(allCommands, parseCommand((line)))
		check(err)
	}

	var x, y int = 0, 0

	for i := 0; i < len(allCommands); i++ {
		c := allCommands[i]
		println(c.cmd, c.val)
		switch c.cmd {
		case Forward:
			x += c.val
		case Up:
			y += c.val
		case Down:
			y -= c.val
		}
	}
	fmt.Println(x, y)
	fmt.Println(x * y)
}
