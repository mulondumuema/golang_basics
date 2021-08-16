package main

import (
	"log"
	"runtime/trace"
)

var myInt int
var myUnint uint // for positive numbers

var myFloat32 float32

var myFloat64 float64 // for large numbers

func main() {
	myInt = 10
	myUnint = 20


	myFloat32 = 100.1000000001
	myFloat64 = 100.1000000001

	log.Println(myInt, myUnint, myFloat32, myFloat64)

	myString := ""

	log.Println(myString)

	myString = "John" // you haven't changed the value of myString. This is because strings in go are immutable. That means when yu are trying to change a value of a new string what you are actually doing is creating a new string and storing that in your variable. BUt anyway, you can treat strings as mutable but know they are immutable.

	log.Println(myString)

	var myBool = true
	
}