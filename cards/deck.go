package main

// import multiple standard library
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Club", "Diamonds", "Heart", "Spades"}
	cardValues := []string{"A", "1", "2", "3", "4","5"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+ " of "+ suit)
		}
	}

	return cards
}

// Function arguments
// return multiple values
func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print(){
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string { 
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	// convert deck -> to string -> to byte slice
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) (deck) {
	// if nothing went wrong error will be `nil`
	bs, err := ioutil.ReadFile(filename)

	// Error Handling
	if err != nil {
		// Option #1 log the error and return a call to `newDeck`
		// Option #2 log the error and exit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ", ") // convert byte slice to string and split

	return deck(s)
}

func (d deck) shuffle() {
	for idx := range d {
		// sudo random generator, depends upon some seed value
		// will use exact same seed, hence same exact sequence
		// newPosition := rand.Intn(len(d) -1)	

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		newPosition := r.Intn(len(d)-1)
		d[idx], d[newPosition] = d[newPosition], d[idx]
	}
}
