package main

import (
	"fmt"
)

var number int

func main(){
	/*
		Variable declaration - static type language
		bool, string, int, float
		1- var - (variable) - name - type = value variable contains, it should always be of the assigned type
		2- name:="value of variable", compiler identifies the type automatically. Only use the colon during initialisation.
		3- Declare the variable outside the main function and initialise it inside function.

		**files in the same package dont have to be imported in go
	*/
	var card string = "ace of spades"
	fmt.Println(card)

	card1:=newCard()
	fmt.Println(card1)

	card1="ace of hearts"
	// card1:="ace of hearts" //will give an warning - no new variable warning
	fmt.Println(card1)

	number=10
	fmt.Println(number)
}

func newCard() string{
	return "5 of Diamonds"
}