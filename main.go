package main

var p Player
var current int
var deck *Deck
var board []*Card

func main() {
	deck = LoadDeck()
	p.DealCards()
	Repl()
}
