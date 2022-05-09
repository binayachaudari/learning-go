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

	/*
		Slice Range 
		vehicles[startIndexIncluding: endIndexNotIncluding]
			- startIndex and endIndex are optional
			- empty startIndex infer start form the beginning 
			- empty endIndex infer upto very end 
	*/
	fmt.Println(vehicles[0:2]) // [Car Jeep]
	fmt.Println(vehicles[:2]) // [Car Jeep]
	fmt.Println(vehicles[3:]) // [Bus Mini Van]

}
 
func getBus() string {
	return "Bus"
}
