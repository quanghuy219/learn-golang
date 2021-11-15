package main


func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

func newCard() string {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"	// ":=" only used for declare variable, use "=" for assigning value
	return card
}
