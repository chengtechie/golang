package main

import (
	"fmt"
	"os"
)

var size int // init first assign later

func printArray() {
	fmt.Println("")
	size = 5
	fmt.Println(size)

	//Array (fixed size)
	//Slice (dynamic size)
	elements := []int{1, 2, 3}
	fmt.Println(elements)
}

func checkEvenOdd() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even, ", i)
		} else {
			fmt.Printf("%d is odd, ", i)
		}
	}
}

func readFile(fileName string) {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(dat))
}
