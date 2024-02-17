package main

import "fmt"

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades") // Does not modify the existing, instead create a new one

	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	fmt.Println("--- hand ---")
	hand.print()
	fmt.Println("--- remainingCards ---")
	remainingCards.print()
	fmt.Println("--- cards ---")
	cards.print()

	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

	// cards := newDeck()
	cards.shuffle()
	cards.print()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	cardsFromFile := newDeckFromFile("my_cards")
	cardsFromFile.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
