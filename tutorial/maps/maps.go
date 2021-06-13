package main

import (
	"fmt"
)

func length(m map[string]int){
	fmt.Println(len(m))
}

func main(){
	marks:=map[string]int{
		"physics":10,
		"maths":16,
		"chemisry":15,
		"biology":20,
		"english":22,
	}
	fmt.Println(marks)

	// var marks map[string]int
	// marks:=make(map[string]int) // these both syntax are correct for declaring a map

	// deleting from map
	delete(marks,"english")
	fmt.Println(marks)

	// iterating over map
	for k,v:= range marks{
		fmt.Println(k,v)
	}

	// passing to a function
	length(marks)
}