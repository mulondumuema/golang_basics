package main

import "fmt"

// interface type

type Animal interface {
	Says() string
	HowManyLegs() int
}

// Dog is the type for dogs

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string { // creating a receiver for dog
	return d.Sound
}
func (d *Dog) HowManyLegs() int { // creating a receiver for dog
	return d.NumberOfLegs
}

// Cat is the type for cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string { // creating a receiver for dog
	return c.Sound
}
func (c *Cat) HowManyLegs() int { // creating a receiver for dog
	return c.NumberOfLegs
}

func main() {
	// ask a riddle
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	Riddle(&dog)

	cat := Cat{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
		HasTail:      true,
	}

	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
// so an interface circumvents the need to write two functions that look/ do exactly the same thing. It requires fewer lines of code and it is not too difficult to implement
// Note: once you define an interface, you should list all the functions that the interface must have inorder to be that interface. e.g if Dog only had Says() function it would not satisfy the requirements for being an animal. So it needs to have both Says() and HowManyLegs functions so as to use the functions defined for Animal type and Dog type
