package main

import "fmt"

type contactInfo struct {
	email string
	postalCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	ironMan := person{
		firstName: "Tony", 
		lastName: "Stark",
		contact: contactInfo{
			email: "ironman@email.com",
			postalCode: 45210,
		},
	}

	ironMan.updateName("Iron Man") // pass by value, name won't be updated

	// ironManPointer := &ironMan // memory address of `ironMan`
	// ironManPointer.updateNameByPointer("Anthony Edward") // pass by reference

	// shortcut way of passing pointer
	ironMan.updateNameByPointer("Anthony Edward") // pass by reference

	ironMan.print()
}

// pass by value
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// pass by reference
func (pointerToPerson *person) updateNameByPointer(newFirstName string) {
	// get value at the given pointer `*pointerToPerson`
	(*pointerToPerson).firstName = newFirstName
}

// Receiver function
func (p person) print() {
		fmt.Printf("%+v", p)
}
