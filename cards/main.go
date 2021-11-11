package main

import "fmt"

func main() {
	var cardsArr [2]string	// Array
	fmt.Println(cardsArr)

	cards := []string{newCard()}	// Slice
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}

	a := deck{"a"}
	a.print()
}

func newCard() string {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"	// ":=" only used for declare variable, use "=" for assigning value
	return card
}
