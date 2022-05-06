package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logger struct {}

func main() {
	 res, err := http.Get("https://binayachaudari.com.np/YIFY/Developer/")

	 if err != nil {
		 fmt.Println("Error: ", err)
		 os.Exit(1)
	 }

	 // io.Copy(<Writer Interface>, <Reader interface>)
	 io.Copy(os.Stdout, res.Body)

	 /*
	 	Custom method implementation in Writer interface
	 */
	 log := logger{}
	 io.Copy(log, res.Body)
}

// Writter interface
func (logger) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("Wrote bytes:", len(b))

	return len(b), nil
}
