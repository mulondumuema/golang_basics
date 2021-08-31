package main

import "fmt"

// an expression is some bit of code that can be evaluated to a single value

func main() {
	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed: %t", name, age, rightHanded)  //name, age, righthanded are expressions
	fmt.Println()

	ageInTenYears := age + 10 // this is an expression

	fmt.Printf("In ten years, %s will be %d years old", name, ageInTenYears)   // name, ageInTenYears is an expression
	fmt.Println()
}