package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// the while loop executes as long as the condition is true
	rand.Seed(time.Now().UnixNano())

	i := 1000
	// execute a loop while i > 100

	for i > 100 {
		// get a random number
		i = rand.Intn(1000) + 1
		if i > 100 {
			fmt.Println("i is", i, "so loop keeps going")
		}
	}

	fmt.Println("Got", i, "and broke out of loop")
}