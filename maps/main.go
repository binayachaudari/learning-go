package main

import "fmt"

/*
 Maps are collection of key value pairs
 Kyes and values are statically typed, all the values and keys must be of same type
*/

func main() {
	/*
	 1. Way to declare map
	*/
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	fmt.Println(colors) // map[blue:#0000ff green:#00ff00 red:#ff0000]


	/*
	 2. Way to declare map
	*/
	var cars map[string]string

	fmt.Println(cars)


	/*
	 2. Way to declare map
	*/
	person := make(map[string]string)

	fmt.Println(person)


	/*
	**
	*/

	// Add key, we always use `[]` syntax to access map value
	colors["white"] = "#ffffff"

	// delete key
	delete(colors, "red")

	printMap(colors)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, " is", value)
	}
}
