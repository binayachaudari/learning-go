package main

import "fmt"

// Create new type `vehicle` that extends existing type
// slice of strings
type vehicle []string

/*
syntax: func <receiver> <function-name> {...}
receiver of type vehicle, function name is `print`
any variable inside our application of type `vechile` now gets access to `print` method

The variable `v` is reference to the actual copy of variable of type vehicle that we're working with
like 
	- `this` in JS
	- `self` in Python
*/
func (v vehicle) print(){
	for idx, item := range v{
		fmt.Println(idx, item)
	}
}
