package main

import "fmt"

func main() {
	apples := 18
	oranges := 9

	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// operations are: > < >= <= 
	fmt.Printf("%d > %d: %t", apples, oranges, apples > oranges)
	fmt.Println()
	fmt.Printf("%d < %d: %t", apples, oranges, apples < oranges)
	fmt.Println()
	fmt.Printf("%d >= %d: %t", apples, oranges, apples >= oranges)
	fmt.Println()
	fmt.Printf("%d <= %d: %t", apples, oranges, apples <= oranges)
	fmt.Println()

	// note: the operations can only be done on similar types
	
}