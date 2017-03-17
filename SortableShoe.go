package main

import (
	"math/rand"
	"sort"
	"time"
)

type SortableShoe []Card

func (s SortableShoe) Len() int {
	return len(s)
}

func (s SortableShoe) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableShoe) Less(i, j int) bool {
	return s[i].Position < s[j].Position
}

func NewSortableShoe(numberOfDecks int) SortableShoe {

	sortableShoe := SortableShoe{}

	for i := 0; i < numberOfDecks; i++ {
		deck := NewDeck()
		for _, card := range deck.Cards {
			sortableShoe = append(sortableShoe, card)
		}
	}

	rand.Seed(time.Now().Unix())
	perm := rand.Perm(numberOfDecks * 52)

	for i := 0; i < len(sortableShoe); i++ {
		sortableShoe[i].Position = perm[i]
	}
	sort.Sort(sortableShoe)

	return sortableShoe
}
