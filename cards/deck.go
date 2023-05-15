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
