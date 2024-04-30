package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Repl() {
	fmt.Println("Welcome to Survival Game")
	fmt.Printf("> ")
	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		cmdName := strings.ToLower(input.Text())
		if len(cmdName) == 0 {
			continue
		}

		availableCommands := GetCommands()

		command, ok := availableCommands[cmdName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		command.Callback()
	}
}
