package main

type Deck struct {
	Characters []struct {
		Name             string `yaml:"name"`
		Cost             int    `yaml:"cost"`
		Health           int    `yaml:"health"`
		Attack           int    `yaml:"attack"`
		SpecialAbilities []struct {
			Name        string `yaml:"name"`
			Description string `yaml:"description"`
			Effect      string `yaml:"effect,omitempty"`
		} `yaml:"special_abilities"`
	} `yaml:"characters"`
}
