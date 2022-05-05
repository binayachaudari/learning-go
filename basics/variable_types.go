package main

import "fmt"

func main() {
	/* 
		Basic fundamental variable types
			- bool
			- int
			- string
			- float64
	*/

	// var card string = "Ace of Spades"
	card := "Ace of Spades" // infer variable type
	number := getNumber() // inter variable type from function return type
	
	fmt.Println(card)
	fmt.Println(number)
}

// return type of `int`
func getNumber() int {
	return 10
}
