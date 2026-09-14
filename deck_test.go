package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("NewDeck returned %d cards, expected 52", len(d))
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()

	sd := newDeck()
	sd.shuffle()

	for index := range sd {
		if sd[index] != d[index] {
			return
		}
	}

	t.Errorf("Shuffle didn't work")
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	fn := "myTestDeck"

	t.Cleanup(func() {
		err := os.Remove(fn)
		if err != nil {
			t.Fatal(err)
		}
	})

	d := newDeck()
	d.saveToFile(fn)

	loadedDeck := newDeckFromFile(fn)

	if loadedDeck == nil {
		t.Errorf("Could not load deck from file")
	}

	for index, card := range d {
		if card != loadedDeck[index] {
			t.Errorf("Loaded deck does not match saved deck")
			break
		}
	}
}
