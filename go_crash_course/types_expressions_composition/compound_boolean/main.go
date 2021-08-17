package main

import "fmt"


type Employee struct {
	Name string
	Age int
	Salary int
	FullTime bool
}

func main(){
	jack := Employee {
		Name: "Jack Doe",
		Age: 28,
		Salary: 400000,
		FullTime: false,
	}

	jill := Employee {
		Name: "Jill Johns",
		Age: 33,
		Salary: 600000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {

		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 0r older")
		} else {
			fmt.Println(x.Name, "is under 30")
		}

		if x.Age > 30 && x.Salary > 500000 {
			fmt.Println(x.Name, "makes more than 500000 and is over 30")
		} else {
			fmt.Println("Either", x.Name, "makes less than 500000 and is under 30")
		}

		
	}
}