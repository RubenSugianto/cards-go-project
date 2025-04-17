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

	// numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for _, number := range numbers {
	// 	if number%2 == 0 {
	// 		fmt.Println(number, "is even")
	// 	} else {
	// 		fmt.Println(number, "is odd")
	// 	}
	// }

}
