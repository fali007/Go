package main
/*
	go build hello.go - build the machine code for this program
	./hello will run to execute the machine code compiled in above command
	go run hello.go - will build and execute
*/

/*
	Package = project = workspace -is a collection of common source code files.
	package can have many related files inside it.
	first line of the file should declare the package it belongs to.

	We call it main - 
		We can have executable and helper packages(dependencies) in go
		While compiling if package main, it will generate runnable aka executable code. If not main it wont be executable.
		We also should create a func main in executable package
*/

import (
	"fmt"
	"reflect"
)
/*
	This import gives the access to all the code that conatins in package fmt.
*/

/*
	How to initialise a function
	func - name - (arguments) - returning values
*/

func main(){
	/*
		The func is mandatory for executable file. Starting point.
		Here we are not taking any arguments and not returning anything.
	*/
	fmt.Println("Hello World")
	apps:=[]App{}
	fmt.Println(reflect.TypeOf(apps))
}
