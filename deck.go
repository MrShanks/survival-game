package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Deck struct {
	Cards []Card `yaml:"cards"`
}

type Card struct {
	Name   string `yaml:"name"`
	Kind   string `yaml:"kind"`
	Cost   int    `yaml:"cost"`
	Health int    `yaml:"health"`
	Attack int    `yaml:"attack"`
}

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
