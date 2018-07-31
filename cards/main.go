package main

import "fmt"

func main()  {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spade")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
