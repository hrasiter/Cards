package main

import "fmt"

func main() {
	/*Declaring long way varible*/
	//var card string = "Ace of Space"

	/*Declaring short hand variable*/
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
