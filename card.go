package main

import (
	"fmt"
	"strconv"
)

const (
	CARD_TWO   = 2
	CARD_THREE = 3
	CARD_FOUR  = 4
	CARD_FIVE  = 5
	CARD_SIX   = 6
	CARD_SEVEN = 7
	CARD_EIGHT = 8
	CARD_NINE  = 9
	CARD_TEN   = 10
	CARD_JACK  = 10
	CARD_QUEEN = 10
	CARD_KING  = 10
	CARD_ACE   = 11
)

const (
	SUIT_SPADES   = '\u2660'
	SUIT_HEARTS   = '\u2665'
	SUIT_DIAMONDS = '\u2666'
	SUIT_CLUBS    = '\u2663'
)

type Card struct {
	//for display purposes
	Symbol string

	// The suit of the card.
	Suit rune

	// The primary value of the card.
	Value int

	// Some cards (i.e. aces) have an alternate value.
	AlternateValue int
	
	//Position in the deck
	Position int
}

type SortableCards []Card

func (c SortableCards) Len() int{
	return len(c)
}

func (c SortableCards) Swap(i,j int){
	c[i], c[j] = c[j], c[i]
}

func (c SortableCards) Less(i,j int) bool{
	return c[i].Position < c[j].Position
}

func (card Card) String() string {
	return fmt.Sprintf("%s%c\t%d\t%d", card.Symbol, card.Suit, card.Value, card.Position)
}

func NewCard(value int, suit rune, position int) Card {
	card := Card{}
	card.Value = value
	card.Suit = suit
	card.Symbol = strconv.Itoa(value)
	card.Position = position

	switch value {
	default:
		card.AlternateValue = value
	case 11:
		card.Symbol = "J"
		card.Value = 10
	case 12:
		card.Symbol = "Q"
		card.Value = 10
	case 13:
		card.Symbol = "K"
		card.Value = 10
	case 14:
		card.Symbol = "A"
		card.Value = 11
		card.AlternateValue = 1

	}

	return card

}
