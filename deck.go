package main

import ()

type Deck struct {
	Cards []Card
}

func generateSuit(suit rune, deck Deck) Deck {
	for i := 1; i < 14; i += 1 {
		deck.Cards = append(deck.Cards, NewCard(i+1, suit, i))
	}

	return deck
}

// Returns a new set of cards. Whew.
func NewDeck() Deck {
	deck := Deck{}
	deck = generateSuit(SUIT_HEARTS, deck)
	deck = generateSuit(SUIT_DIAMONDS, deck)
	deck = generateSuit(SUIT_CLUBS, deck)
	deck = generateSuit(SUIT_SPADES, deck)
	return deck
}
