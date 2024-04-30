package main

import (
	"fmt"
	"log"
	"os"

	"k8s.io/apimachinery/pkg/util/yaml"
)

func LoadDeck() *Deck {

	yamlFile, err := os.ReadFile("characters.yaml")
	if err != nil {
		log.Println(err)
	}

	deck := &Deck{}
	err = yaml.Unmarshal(yamlFile, deck)
	if err != nil {
		log.Println(err)
	}

	return deck
}

func main() {
	deck := LoadDeck()
	fmt.Println(deck)
	Repl()
}
