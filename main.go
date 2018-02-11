package main

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
