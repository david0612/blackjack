package main

import (
	
	"math/rand"
	
	"time"
)

type Shoe struct {
	Name  string
	Decks []Deck
}

func NewShoe(numberOfDecks int, name string) Shoe {
	shoe := Shoe{}
	shoe.Name = name
	for i := 0; i < numberOfDecks; i++ {
		shoe.Decks = append(shoe.Decks, NewDeck())
	}
	
	rand.Seed(time.Now().Unix())
	perm := rand.Perm(numberOfDecks * 52)

	for i, deck := range shoe.Decks {
		for j, _ := range deck.Cards {
			position := perm[ ( i * 52 ) + j ]
			positionAddress := &shoe.Decks[i].Cards[j].Position
			assignPosition(positionAddress, position)
		}
	}

	return shoe
}

func assignPosition(positionPointer *int, position int){
	*positionPointer = position
}
