package main

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name     string
	Desc     string
	Callback func()
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:     "help",
			Desc:     "Print the help menu",
			Callback: CallbackHelp,
		},
		"exit": {
			Name:     "exit",
			Desc:     "Terminate the programm immediately",
			Callback: CallbackExit,
		},
	}
}

func CallbackExit() {
	os.Exit(0)
}

func CallbackHelp() {
	fmt.Println()
	fmt.Println("These are the available commands:")
	for _, cmd := range GetCommands() {
		fmt.Printf("- %s: %s\n", cmd.Name, cmd.Desc)
	}
	fmt.Printf("> ")
}
