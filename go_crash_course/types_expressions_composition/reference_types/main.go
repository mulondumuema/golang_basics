package main

import "fmt"

// reference types(pointers, sices, maps, functions, channels)

//  pointers is something that points to a specific location in memory

// func main() {
// 	var myInt int   // here we have created a variable myInt that we assign to hold the value 10 but it is also possible to declare variables which point to the location where that 10 is stored
// 	myInt = 10

// 	fmt.Println(myInt)
// }

// why would one need a pointer?
func main() {
	x := 10
	myFirstPointer := &x // the '&' allows us to access the location for x. i.e the variable myFirstPointer holds the address where x is stored in memory

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println("x is now", x) // even without touching the variable x, we have change the value that it holds
}