package main

import (
	"errors"
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

	// relational operators: e.g and
	// it is good to put parenthesis to increase readability
	second := 31
	minute := 1

	if (minute < 59) && ((second + 1) > 59) {
		minute++
	}

	fmt.Println(minute)

	a1 := 12.0
	b1 := 5.0

	// incase b is put to zero, the division by zero will return an error
	// to hadle it:

	// if b1 != 0 {
	// 	c := divideTwoNumbers(a1, b1)

	// 	if c == 2 {
	// 		fmt.Println("We found a two")
	// 	}
	// }

	// or you can use the short circuit evaluation
	// if b1 != 0 && divideTwoNumbers(a1, b1) == 2 {
	// 	fmt.Println("We found a two")
	// }

	// note: in this case you cannot put b != 0 after the evaluation since this might lead to a zero division being evaluated

	c, err := divideTwoNumbers(a1,b1)
	if err != nil {
		fmt.Println(err)
	} else {
		if c == 2.0 {
			fmt.Println("We found 2")
		}
	}

	// let's try above with the blank operator, to see what happens
	_, myerr := divideTwoNumbers(a1,b1)
	if err != nil {
		fmt.Println(myerr)
	} else {
		fmt.Println("the blank assignement works!")
	}

	i := 0

	for {
		i++
		if i > 3 {
			break
		}
		fmt.Println("i is", i)
	}
}

// func divideTwoNumbers(x, y int) int {
// 	return x / y
// }

// to handle the zero division, you can add the error handling in the divideTwoNumbers function
// in this case, the function will return two values i.e the output of the evaluation and error. the error is a built-in type from go

func divideTwoNumbers(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0") // you need to return something, so we return 0 and the error message
	}
	return x / y, nil
}