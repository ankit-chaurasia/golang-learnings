package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
	// return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(fileName string) error {
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := os.ReadFile(fileName)
	if err != nil {
		// Option 1 - log the error and return a call to newDeck()
		// Option 2 - log the error and entirely quit the program
		log.Fatal(err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ", ")
	return deck(s)
	// os.Stdout.Write(byteSlice)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		// newPosition := rand.Intn(len(d) - 1)
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
