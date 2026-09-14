package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	d := newDeckFromFile("meuDecks.txt")
	d.print()

	d.shuffle()
	d.print()

	io.Copy(deckWriter{}, deckReader{})
}

type deckReader struct{}

type deckWriter struct{}

func (d deckReader) Read(p []byte) (n int, err error) {
	data, err := os.ReadFile("meuDeck")
	if err != nil {
		return 0, err
	}

	n = copy(p, data)

	return n, io.EOF
}

func (d deckWriter) Write(p []byte) (n int, err error) {
	fmt.Println(strings.Split(string(p), "\n"))
	return len(p), nil
}
