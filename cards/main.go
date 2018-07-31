package main

import "fmt"

func main()  {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spade")

	fmt.Println(cards)
}

func newCard() string {
	return "Ace of Spades"
}
