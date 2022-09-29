package stephen

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func ExecuteDeck(cards Deck) {
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("---")
	remainingCards.shuffle()
	fmt.Println(remainingCards.toString())

	papers := Deck{"Paper 1"}
	papers.print()
}

func NewDeck() Deck {
	cards := Deck{}
	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, suit+" "+value)
		}
	}
	return cards
}

func (d Deck) print() {
	for i, element := range d {
		fmt.Println(i, element)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// one line swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
