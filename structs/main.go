package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// type {name} struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// You can create var alex person without value, and Go will asign default to each type
	// Explicit parameters is convention
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// Casting: if receiver is *person, we can use *person or person
	jim.updateName("Jimmy")
	jim.print()
}

// Passing Parameters by "Reference" - Pointers !
// & and * are used - C++
// Turn address into value with *address
// Turn value into address value with &value
// *person -> Pointer to person
// This is only necessary for primitive types
func (pointerToPerson *person) updateName(newFirstName string) {
	// Manipulate value of what a pointer is pointing to
	(*pointerToPerson).firstName = newFirstName
}

// Passing Parameters by Copy - This is the default
func (p person) print() {
	// %+v is used for a full print, with name of parameters
	fmt.Printf("%+v", p)
}
