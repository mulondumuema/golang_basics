package main

import "fmt"

// a nested loop is where there is a loop within a loop

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Print("i is", i)
		for j := 1; j <= 3; j++ {
			fmt.Print(" j: ", j)
		}
		fmt.Println()
	}
}

//  we are using the three part loop, but nesting can be applied in any loop