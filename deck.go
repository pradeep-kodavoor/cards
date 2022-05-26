package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {

	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := string(bs)

	stringSplice := strings.Split(s, ",")
	return deck(stringSplice)
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// fmt.Println("Shuffle from : ", i, "to : ", newPosition)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
