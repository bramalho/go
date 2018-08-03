package main

import "fmt"

func main()  {
	cards := newDeck()

	hand, remaingDeck := deal(cards, 5)

	fmt.Println("Hand:")
	hand.print()

	fmt.Println("Remaining Deck:")
	remaingDeck.print()
}
