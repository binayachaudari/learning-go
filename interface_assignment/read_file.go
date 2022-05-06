package main

import (
	"fmt"
	"io"
	"os"
)

type logger struct {}

func main() {
	// Get file name from args
	fileName := os.Args[1]

	// Read File
	file, err := os.Open(fileName)

	// Error handler
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	log := logger{}
	io.Copy(log,file)
}

func (logger) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}
