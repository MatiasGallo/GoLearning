package main

import "fmt"

type bot interface {
	// If a type implements getGreeting() string it becomes bot !
	// And gain access to the func printGreeting
	// If multiple funcs are requested, everyone needs to be implemented
	getGreeting() string
}

// Concrete type
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// The receiver is not used - We can remove it !
func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
