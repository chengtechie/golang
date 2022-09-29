package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func executeDeck(cards deck) {
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("---")
	remainingCards.shuffle()
	fmt.Println(remainingCards.toString())

	papers := deck{"Paper 1"}
	papers.print()
}

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, suit+" "+value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, element := range d {
		fmt.Println(i, element)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// one line swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
