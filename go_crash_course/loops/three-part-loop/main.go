package main

import "fmt"

func main() {
	// // infinite for loop
	// for {
	// 	fmt.Println("Hi there an running infinitely")
	// }

	// // to range over a map
	// for _, x := range myMap {

	// }
	// the three part loop. It is for doing something a certain 
	// number of times. It has three parts:
	// 1. where your index begins
	// 2. the condition that allows you to run the loop
	// 3. the operation, e.g to increment a value by 1 (i++ or i+1)

	for i := 0; i <= 10; i++ {
		fmt.Println("i is", i)
	}


}