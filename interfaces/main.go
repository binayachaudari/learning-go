package main

import "fmt"

type bot interface {
	getGreeting() string // bot now has a method getGreeting
}

/*
	type bot interface {
		getGreeting(string, int) (string, error)
		getBotVersion() float
		respondToUser(user) string
	}
*/

// members of bot
type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())	
}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "HOLA!"
}
