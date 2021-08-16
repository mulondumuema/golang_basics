package main

import "fmt"

// aggregate types (array, struct)

// arrays:
// arrays are not commonly used in go since we have slices which have more functionality

// func main() {
// 	var myStrings [3]int
// 	myStrings[0] = 1
// 	myStrings[1] = 2
// 	myStrings[2] = 3

// 	fmt.Println("First element in array is", myStrings[0])
// }

// struct
// structs are considered an aggregate type because struct can hold many pieces of information. It aggregates information

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	// you can define using dot notation:
	// var myCar Car
	// myCar.NumberOfTires = 4
	// myCar.Luxury = false
	// myCar.BucketSeats = true
	// myCar.Make = "vitz"
	// myCar.Year = 2000
	// myCar.Model = "tyu"

	// or you can define:
	myCar := Car {
		NumberOfTires: 4,
		Luxury: true,
		BucketSeats: true,
		Make: "Volvo",
		Model: "XC90",
		Year: 2019,
	}

		fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)

}