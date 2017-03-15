package main

import (
	"fmt"
	//"strconv"
	"time"
)

func main() {
	var shoe Shoe
	shoe = NewShoe(6, "name")
	
	for _, deck := range shoe.Decks {
		for _, card := range deck.Cards {
			fmt.Println(card.String())
		}
	}

time.Sleep(10000 * time.Millisecond)
}
