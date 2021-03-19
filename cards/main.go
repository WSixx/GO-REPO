package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//fmt.Println(card)
	cards := newDeck()
	fmt.Println(cards.toString())

}
