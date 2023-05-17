package main

import "fmt"

// Map is Reference Type - Indexed Keys
// Struct is Value Type	 - Not Indexed Keys

func main() {
	//Very similar to dict (Python)
	// syntax is map[{type keys}]{type values}
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//Adding and deleting from Map
	colors["yellow"] = "laksjdf"
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(c map[string]string) {
	// Destructuring returns key, value
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
