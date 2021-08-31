package main

import "fmt"

// Assigning functions to types

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {  // the '(a *Animal)' is a receiver and it is used to link a function to a type
	fmt.Printf("a %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) HowManyLegs(){
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
}

func main() {
	myTotal := sumMany(1, 2, 4, 5, 6, 9)
	fmt.Println(myTotal)

	var dog Animal 
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.Says()

	cat := Animal {
		Name: "cat",
		Sound: "meow",
		numberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()

}

// writing a function that can take any number if arguments (this is called a variadic function)
// if you have a function that includes multiple parameters and one of the has to be variadic, then your variadic has to be at the end so that your compiler can know which is which

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}


