package main

import (
	"fmt"
	"os"
)

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

func CallbackHandPrint() {
	for _, card := range p.InitHand {
		genericCard := card
		prettyPrint(genericCard)
		fmt.Println()
	}
	fmt.Printf("> ")
}

func CallbackSelect() {
	currentCard := p.InitHand[current]
	prettyPrint(currentCard)
}

func CallbackNext() {
	current += 1
	lastCard := len(p.InitHand) - 1
	if current >= lastCard {
		current = lastCard
	}
	currentCard := p.InitHand[current]
	prettyPrint(currentCard)
}

func CallbackPrev() {
	current -= 1
	if current <= 0 {
		current = 0
	}
	currentCard := p.InitHand[current]
	prettyPrint(currentCard)
}

func CallbackPlayCard() {
	board = append(board, &p.InitHand[current])
	p.InitHand = append(p.InitHand[:current], p.InitHand[current+1:]...)
	current -= 1
}
