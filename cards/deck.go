package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
// Extendes/borrow []string
type deck []string

// Receiver function
// Parameters are copied ! Not referenced
// By convention the receiver is the first letter of type deck -> d
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Replace variable unused with '_'
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// slice[startIncluded : endNotIncluded]
// Deal from 0 to handSize - and another one from handSize to end
// New references that point at subsections of d - Never directly modified d.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
