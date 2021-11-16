package main


func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"	// ":=" only used for declare variable, use "=" for assigning value
	return card
}
