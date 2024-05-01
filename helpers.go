package main

import "fmt"

func prettyPrint(card Card) {
	fmt.Printf("- Name: %s\n", card.Name)
	fmt.Printf("  Cost: %d\n", card.Cost)
	fmt.Printf("  Health: %d\n", card.Health)
	fmt.Printf("  Attack: %d\n", card.Attack)
}
