package main

func main() {
	// var card string = "Ace of spades"
	/* card := newCard()
	fmt.Println(card)
	cards := deck{card, "Six of Hearts"}
	cards = append(cards, "Two of Spades") */
	// cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 3)
	// hand.print()
	// remainingCards.print()
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of diamonds"
}
