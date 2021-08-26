package main

import (
	"fmt"
	"math"
)

func main() {
	// answer := 7 + 3 * 4
	// fmt.Println("Answer:", answer)

	// answer = (7 + 3) * 4
	// fmt.Println("Answer:", answer)

	// mutiplication
	// area of a circle
	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("The area of the circle is:", area)

	// integer division
	half := 1 / 2 
	fmt.Println("half division with integer division", half)

	halfFloat := 1 / 2 
	fmt.Println("half float division with integer division", halfFloat)

	// squaring (raising to the power)
	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("good three squared", goodThreeSquared)

	// modulus
	remainder := 50 % 3
	fmt.Println("remainder is", remainder)

	// unary operator
	x := 3
	x ++
	fmt.Println("x is now", x)

	x--
	x--
	fmt.Println("x is now", x)

	// precedence - follow the order of operations
	//for multiplication and division, it is redudant, so there is no need for brackets
	a := 12.0 * 3.0 / 4.0
	b := (12.0 * 3.0) / 4.0
	c := 12.0 * (3.0 / 4.0)
	fmt.Println("a", a, "b", b, "c",c)

	// unclear
	// integer division 
	unclear := 12 * (3/4)
	fmt.Println("unclear", unclear)

	// oarenthesis have highest precedence
	f := 12.0 / 3.0 /4.0
	fmt.Println("f", f)
	f = 12.0 / (3.0 /4.0)
	fmt.Println("f", f)

	// modulus operator
	for m := 1; m <= 12; m++ {
		fmt.Println("The month after", m, "is", m%12 + 1)
	}

}