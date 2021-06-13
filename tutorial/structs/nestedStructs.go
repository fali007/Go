package main

import (
	"fmt"
	"encoding/json"
)

type class struct{
	Grade int
	Strength int
	School string
	Students []person
}

type person struct{
	Name string
	Roll int
}

func (c *class) print(){
	js,_:=json.Marshal(c)
	fmt.Printf("\n%s",js)
}

func (c *class) pname(){
	c.Grade=1// here we can see the type miss match, go is taking care of it internally
}

func (c class) name(){
	c.Grade=2
}

func main(){
	cls:=class{Grade:10,Strength:40,School:"JNV",Students:[]person{{Name:"Hari",Roll:21},{Name:"Anju",Roll:38}}}
	cls.name()
	cls.print()
	cls.pname()
	cls.print()
}