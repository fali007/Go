package main

import (
	"fmt"
	"encoding/json"
)

type card struct{
	Suit      string   `json:"suit"`
	Value     int      `json:"value"`
}

func (c *card) print(){
	js,_:=json.Marshal(c)
	fmt.Printf("%s",js)
}

func main(){
	var c card
	c.Suit="Heart"
	c.Value=2
	c.print()

	cd:=card{"Spade", 3} //we use the order of the fields for assignment here.
	cd=card{Suit:"Spade",Value: 3}
	cd.print()
}

// pointer receiver - here we are passing the pointer of the struct as the receiver to the function and this helps us to change the value in the location without giving any return value
// This can be used for any object maipulation tasks where we can manipulate without returing the large object this can be helpful for optimisong and making code faster.
