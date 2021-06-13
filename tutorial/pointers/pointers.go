package main

import (
	"fmt"
)

type name struct{
	Name string
	Age int
}

func changeName(n *name){
	(*n).Name="New Name"
	//or
	n.name="New Name" //both will work
}

func change(a *int, b *int){ //*int is pointer type telling function is accepting memory location as argument
	fmt.Println(*a,*b) //*a getting the value in the memory location
}

func main(){
	a:=10
	b:=20
	change(&a,&b) // passing the memory location of a and b to the function

	n:=name{"tom",18}
	changeName(&n)// passing the address of struct to the function
	fmt.Println(n)
}