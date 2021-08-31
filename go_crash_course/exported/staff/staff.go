package staff

import "log"

var OverPaidLimit = 750000
var underPaidLimit = 200000

type Employee struct {
	FirstName string
	LastName string
	Salary int
	FullTime bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			overpaid = append(overpaid, x)
		}
	}

	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee

	myFunction()
	for _, x := range e.AllStaff {
		if x.Salary < underPaidLimit {
			underpaid = append(underpaid, x)
		}
	}

	return underpaid
}

func (e *Office) notVisible() {
	log.Println("Hello, world")
}

func myFunction() {   // this is a function without a receiver. It hav no acces to any of the values in the Office type like the other functions
	log.Println("I am a function")
}
// variables, functions that begin with uppercase letter can be exported while those with lower case letters are nort exported.