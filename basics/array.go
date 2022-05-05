package main

import "fmt"

func main() {
	/*
		# Array
			- very basic/primitive data structure for holding list of records
			- fixed length of records

		# Slice
			- like an array that can grow or shrink (dynamic array)

			Both array and slices must be defined with a data type,
			every elements must be of same type
	*/

	// Slice
	vehicles := []string{"Car", "Jeep", "Motercycle", getBus()}
	
	// `append(variable, newValue)` to add new element in slice
	vehicles = append(vehicles, "Mini Van") // Immutabe (returns new slice)

	// Iterate over slice (forEach loop)
	for index, vehicle := range vehicles{
		fmt.Println(index, vehicle)
	}
}
 
func getBus() string {
	return "Bus"
}
