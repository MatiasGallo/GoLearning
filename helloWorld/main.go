package main // Executable package !

import "fmt"

func main() {
	// Static Typing ! Can be infered
	// var card string = "Ace of Spades" - Equivalent complete declaration
	firstCard := "Ace of Spades"
	fmt.Println(firstCard)

	secondCard := newCard()
	fmt.Println(secondCard)

	// Array fixed length
	// Slice - can grow and shrink (List!)
	cards := []string{"Ace of Diamonds", newCard()}
	// Returns new slice
	cards = append(cards, "Six of Spades")

	// Destructuring - Clearing i and card every loop
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// Explicit data type returns
func newCard() string {
	return "Five of Diamonds"
}
