package stephen

import (
	"fmt"
	"os"
)

var size int // init first assign later

func PrintArray() {
	fmt.Println("")
	size = 5
	fmt.Println(size)

	//Array (fixed size)
	//Slice (dynamic size)
	elements := []int{1, 2, 3}
	fmt.Println(elements)
}

func CheckEvenOdd() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even, ", i)
		} else {
			fmt.Printf("%d is odd, ", i)
		}
	}
}

func ReadFile(fileName string) {
	curr, err1 := os.Getwd()
	if err1 != nil {
		fmt.Print(err1)
	}
	dat, err2 := os.ReadFile(curr + "/stephen/data/" + fileName)
	if err2 != nil {
		fmt.Print(err2)
	}
	fmt.Print(string(dat) + "\n")
}
