package main

import "fmt"

func getDefaultMap() map[int]string {
	fmt.Println("")
	//colors := map[string]string{
	//	"red": "#ff0000",
	//}
	colors := make(map[int]string)
	colors[0] = "a"
	colors[1] = "b"
	colors[2] = "c"
	delete(colors, 1)
	return colors
}

func printMap(m map[int]string) {
	for key, value := range m {
		fmt.Printf("Key: %v Value: %v\n", key, value)
	}
}
