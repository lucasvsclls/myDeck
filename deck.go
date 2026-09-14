package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var suits = [4]string{"Spades", "Clubs", "Hearts", "Diamonds"}

var ranks = [13]string{"Jack", "Queen", "King", "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

type deck []string

func newDeck() deck {
	var card string
	var d = make(deck, 0, 52)

	for _, suit := range suits {
		for _, rank := range ranks {
			card = rank + " of " + suit

			d = append(d, card)
		}
	}

	return d
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
	fmt.Println()
}

func deal(d deck, handSize int) (hand deck, theRest deck) {
	if handSize > len(d) {
		return
	}

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(filename string) {
	err := os.WriteFile(filename, []byte(strings.Join(d, "\n")), 0666)

	if err != nil {
		log.Fatal(err)
	}
}

func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal("\nError: \n", err)
	}

	return deck(strings.Split(string(data), "\n"))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := len(d) - 1; i > 0; i-- {
		random := r.Intn(i + 1)
		d[i], d[random] = d[random], d[i]
	}
}
