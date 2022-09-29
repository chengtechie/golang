package stephen

import "fmt"

type bot interface {
	getGreeting() string
}

type ChineseBot struct{}
type EnglishBot struct{}

func PrintBot(b bot) {
	fmt.Println(b.getGreeting())
}

func (ChineseBot) getGreeting() string {
	return "Chinese"
}

func (EnglishBot) getGreeting() string {
	return "English"
}

type logWriter struct{}

// Count as interface of Writer because the same signature is initialized
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("I am a custom writer")
	return len(bs), nil
}

type shape interface {
	getArea() float64
}

type square struct {
	length float64
}

type triangle struct {
	height float64
	width  float64
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func (t triangle) getArea() float64 {
	return (t.width * t.height) / 2
}

func printShape(s shape) {
	fmt.Println("Area: ", s.getArea())
}

func TestInterface() {
	t := triangle{
		height: 5,
		width:  2,
	}
	s := square{
		length: 3,
	}
	printShape(t)
	printShape(s)
}
