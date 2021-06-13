package main

import(
	"fmt"
)

func updateSlice(s []string){
	s[0]="bye"
	fmt.Println(s)
}

func print(val *string){
	fmt.Println(&val) // everything in go is pass by value, thats why eventhough we are passing pointer go will create a varible for storing that value in the stack while function is called. That is why this is printing different values.
}

func main(){
	ary:=[]string{"Hello","How", "are","you?"}
	updateSlice(ary)
	fmt.Println(ary)

	val:="hello"

	fmt.Println(&val)
	print(&val)
}

//eventhough we are not passing the pointers its updating in the original string array
// slice is actually a pointer to the first element to the array
// eventhough it works using pass by value hence a copy will be created the slice will be pointing to the first element.
