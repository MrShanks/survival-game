package main

import "fmt"

type Character struct {
	Kind      string
	Name      string
	Health    int
	Attack    int
	Abilities []Ability
}

type Ability struct {
	name string
	desc string
}

var WindFury = Ability{
	name: "WindFury",
	desc: "Character can attack twice",
}
var DivineShield = Ability{
	name: "Divine Shield",
	desc: "First time the Character takes damage won't take any damage",
}

func NewCharacter(kind, name string) *Character {
	return &Character{
		Kind:   kind,
		Name:   name,
		Health: 10,
		Attack: 3,
		Abilities: []Ability{
			WindFury,
			DivineShield,
		},
	}
}

func PrintCharacterInfo(ch *Character) {
	fmt.Println()
	fmt.Printf("Kind: %s\n", ch.Kind)
	fmt.Printf("Name: %s\n", ch.Name)
	fmt.Printf("Health: %d\n", ch.Health)
	fmt.Printf("Attack: %d\n", ch.Attack)
	fmt.Printf("Abilities: \n")
	for _, ability := range ch.Abilities {
		fmt.Printf("  - Name: %s\n", ability.name)
		fmt.Printf("    Desc: %s\n", ability.desc)
	}
	fmt.Println()
}
