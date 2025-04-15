package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// card = "Five of Diamonds"

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	// cards.print()
}
