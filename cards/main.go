package main

func main() {
	//var card string = "Ace of Spades"
	//fmt.Println(card)
	cards := newDeck()
	cards.saveToFile("my cards")

}
