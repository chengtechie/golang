package main // must have func main!

import (
	"learning/stephen"
	"os"
)

func main() {

	stephen.CheckEvenOdd()

	stephen.PrintArray()

	stephen.ExecuteDeck(stephen.NewDeck())

	stephen.PrintPeople()

	stephen.PrintMap(stephen.GetDefaultMap())

	stephen.PrintBot(stephen.ChineseBot{})
	stephen.PrintBot(stephen.EnglishBot{})

	stephen.GetUrl("https://google.com")

	stephen.TestInterface()

	stephen.ReadFile(os.Args[1])

	stephen.TestChannel()
}
