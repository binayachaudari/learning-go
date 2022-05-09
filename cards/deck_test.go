package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("Expected deck length of 24, but got %v", len(d))
	}

	if d[0] != "A of Club" {
		t.Errorf("Expected 1st card to be A of Club, but got %v", d[0])
	}

	if d[len(d)-1] != "5 of Spades" {
		t.Errorf("Expected 1st card to be 5 of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T){
	testFilename := "_decktesting"
	os.Remove(testFilename)

	deck := newDeck()
	deck.saveToFile(testFilename)


	loadedDeck := readFromFile(testFilename)
	if len(loadedDeck) != 24 {
		t.Errorf("Expected deck length of 24, but got %v", len(loadedDeck))
	}

	os.Remove(testFilename)
}
