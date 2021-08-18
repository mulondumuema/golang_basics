package main

import (
	"log"
	"myapp/staff"
)

var employee = []staff.Employee {
	{FirstName: "John", LastName: "Smith", Salary: 300000, FullTime:true},
	{FirstName: "Sally", LastName: "Jones", Salary: 200000, FullTime:true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 100000, FullTime:true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 500000, FullTime:true},
	{FirstName: "Magaret", LastName: "Carter", Salary: 900000, FullTime:true},
}

func main(){
	mystaff := staff.Office {
		AllStaff : employee,
	}

	// mystaff.All()

	// log.Println(mystaff.All())
	staff.OverPaidLimit = 500000
	log.Println("Over paid staff",mystaff.Overpaid())
	log.Println("Under paid staff",mystaff.Underpaid())

	
}
