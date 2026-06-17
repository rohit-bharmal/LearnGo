package main

func main() {

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5) // deal is a function that returns two decks
	// hand.print()                           // print is a method of the deck type which is defined in the deck.go file
	// remainingCards.print()                 // print is a method of the deck type which is defined in the deck.go file

	cards := newDeck()
	// fmt.Println(cards.toString()) // print the cards in a string format
	// cards.saveToFile("my_cards") // save the cards to a file
	// cards = newDeckFromFile("my_cards") // read the cards from a file
	// cards.print() // print the cards in a string format
	cards.shuffle()
	cards.print()
}
