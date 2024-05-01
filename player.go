package main

type Player struct {
	InitHand []Card
	Mana     int
}

func (p *Player) DealCards() {
	p.InitHand = deck.Cards[:3]
}
