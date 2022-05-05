package main

import "fmt"

// Define all the properties of a struct
// new custom type person with following fields
type person struct {
	firstName string
	lastName string
}

func main() {
	// create a person
	// alex := person{"Alex", "Anderson"} 
	alex := person{firstName: "Alex", lastName: "Anderson"} // other way to create person
	fmt.Println(alex) // output: {Alex Anderson}

	// Another way to create person
	var cornell person
	fmt.Println(cornell) // output: { }

	// print with properties `%+v`
	fmt.Printf("%+v", alex) // output: {firstName:Alex lastName:Anderson}

	// updating properties of structs
	cornell.firstName = "Cornell"
	cornell.lastName = "Mahoney"
	fmt.Printf("%+v", cornell) // output: {firstName:Cornell lastName:Mahoney}
}
