package main

func main()  {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spade")

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
