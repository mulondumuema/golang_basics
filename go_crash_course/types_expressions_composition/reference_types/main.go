package main

import "fmt"

// reference types(pointers, slices, maps, functions, channels)

//  pointers is something that points to a specific location in memory

// func main() {
// 	var myInt int   // here we have created a variable myInt that we assign to hold the value 10 but it is also possible to declare variables which point to the location where that 10 is stored
// 	myInt = 10

// 	fmt.Println(myInt)
// }

// why would one need a pointer?
// func main() {
// 	x := 10
// 	myFirstPointer := &x // the '&' allows us to access the location for x. i.e the variable myFirstPointer holds the address where x is stored in memory

// 	fmt.Println("x is", x)
// 	fmt.Println("myFirstPointer is", myFirstPointer)

// 	*myFirstPointer = 15
// 	fmt.Println("x is now", x) // even without touching the variable x, we have change the value that it holds

// 	changeValueOfPointer(&x)

// 	fmt.Println("the value now is", x)
// }


// func changeValueOfPointer(num *int) {
// 	*num = 25
// }


// slices
// func main()  {
// 	var animals []string
// 	animals = append(animals, "dog")
// 	animals = append(animals, "fish")
// 	animals = append(animals, "cat")
// 	animals = append(animals,"horse")

// 	fmt.Println(animals) //they are appended in the order that they are added i.e FIFO

// 	// for _, x := range animals { //the _ is the index which in this case is not used
// 	// 	fmt.Println(x)
// 	// }

// 	for i, x := range animals { //the _ is the index which in this case is not used
// 		fmt.Println(i, x)
// 	}

// 	fmt.Println("The slice is", len(animals))

// 	fmt.Println("Element 0 is",animals[0])

// 	fmt.Println("The first two elements are", animals[0:2])

// 	// fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

// 	// sort.Strings(animals)

	
// }


// maps
// maps are reference type so they do not use pointers

func main() {
	intMap := make(map[string]int) //to create a map variable you need to use the make function otherwise you'll create a nil map and you can't do anything with it e.g myMap map[string]string

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// delete(intMap, "four")

	// to check if an element is in a map

	el, ok := intMap["four"]

	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}

	// to change a value

	intMap["two"] = 4

	// maps are commonly used in go because they are fast, easy to use and don't require pointers

	for key, value := range intMap {
		fmt.Println(key, value) // maps are not and cannot be sorted
	
	}

}