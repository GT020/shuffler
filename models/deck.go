package models

import "shuffler/utils"

type Deck struct {
	cards []Card
}

func NewClassicDeck() *Deck {

	suites := []CardSuit{Spade, Heart, Club, Diamond}

	values := []CardValue{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}

	cards := make([]Card, 0, 54)

	for _, suit := range suites {
		for _, value := range values {
			cards = append(cards, Card{suit, value})
		}
	}

	return &Deck{
		cards: cards,
	}
}

func (d *Deck) Shuffle() {
	for i := 0; i < len(d.cards); i++ {
		j := utils.SecureRandomInt(0, len(d.cards))
		k := utils.SecureRandomInt(0, len(d.cards))
		d.cards[k], d.cards[j] = d.cards[j], d.cards[k]
	}
}
