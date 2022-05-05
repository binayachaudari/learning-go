package main

import "fmt"

func main() {
	deckOfCards := newDeck()

	// Multiple value returend from the function
	hand, remainingCards := deal(deckOfCards, 5)

	// Print items
	hand.print()
	remainingCards.print()

	// Save to File
	hand.saveToFile("hand")
	remainingCards.saveToFile("remainingCards")

	// Read from file
	readHand := readFromFile("hand")
	readRemainingCards := readFromFile("remainingCards")

	fmt.Println(readHand)
	fmt.Println(readRemainingCards)

	// Randomly shuffle using random number generator
	deckOfCards.shuffle()
	deckOfCards.print()
}

