package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type vietnameseBot struct{}

func main() {
	eb := englishBot{}
	vb := vietnameseBot{}

	printGreeting(eb)
	printGreeting(vb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (vietnameseBot) getGreeting() string {
	return "Xin chào"
}
