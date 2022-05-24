package main

import "fmt"

// Create a new type of 'deck'
// Slice of string
type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuites := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func deal(d deck, hand int) (deck, deck) {
	return d[:hand], d[hand:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
