package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// card = "Five of Diamonds"

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.print()

	// greeting := "Hi There!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// fmt.Println(cards.toString())

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_card")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
