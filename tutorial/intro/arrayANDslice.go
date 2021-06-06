package main

import (
	"fmt"
)

func main(){
	/*
		Arrays are very primitive
			fixed length
		Slice is array with more features
			can grow or shrink
		Both should contain records of same type
	*/
	cards:= []string{"Ace of Diamonds",newCard()}
	cards=append(cards,"Queen of Hearts")

	for i,card:= range cards{
		fmt.Println(i,card)
	}

	/*
		i - index of element
		card - value of element
		range cards will take slice cards and loop over it
		run the inner code for each different card
		we are using ':=' because we are throwing away variables after each iteration
	*/
}

func newCard() string{
	return "Jack of Clubs"
}