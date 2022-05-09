package main

func main() {
	// vehicle type from `vehicle_type.go`
	vehicles := vehicle{"Car", "Jeep", "Motercycle", getBus()}
	
	vehicles = append(vehicles, "Mini Van") 

	// methods from `vehicle_type.go`
	// `vehicles` gets passed to the receiver reference variable
	vehicles.print()
}
 
func getBus() string {
	return "Bus"
}
