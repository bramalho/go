package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println("Deck:")
	cards.print()

	hand, remainingDeck := deal(cards, 10)

	fmt.Println("Hand:")
	hand.print()

	fmt.Println("Remaining Deck:")
	remainingDeck.print()

	fmt.Println("Deck to string: " + cards.toString())
	cards.saveToFile("files/cards")

	fmt.Println("New deck from file:")
	tenCards := newDeckFromFile("files/tenCards")
	tenCards.print()

	fmt.Println("Shuffle")
	cards.shuffle()
	cards.print()
}
