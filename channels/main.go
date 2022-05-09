package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://binayachaudari.com.np",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://twitter.com",
	}
	
	// channels
	c := make(chan string)

	for _, url := range urls {
		// Go routines
		go checkStatus(url, c)
	}

	// Receive value from channel
	for l := range c {
		// function literal with Go Routine
		go func(url string){
			time.Sleep(5* time.Minute)
			// never share variables between Go routines
			// Pass it as a argument or communicate throug channel
			checkStatus(url, c)
		}(l)
	}
}

func checkStatus(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "is down!")

		// Send value through channel
		c <- url
		return
	}

	fmt.Println(url, "is up!")

	// Send value through channel
	c <- url
}
