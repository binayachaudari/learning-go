// Every Go program must declate a package that it belongs to at the very beginning
// Package is a collection of common source codes
// `main` for generating executable file
package main

// Access to code writter in another packages
// `fmt` is a Go standard library, shorthand for `format`
// Standard library packages `https://golang.org/pkg`
import "fmt"

// Main function, code execution starts here
func main(){
	fmt.Println("Hello World")
}
