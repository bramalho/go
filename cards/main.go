package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println("Deck:")
	cards.print()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("Hand:")
	hand.print()

	fmt.Println("Remaining Deck:")
	remainingDeck.print()

	fmt.Println("Deck to string: " + cards.toString())
	cards.saveToFile("cards")
}
