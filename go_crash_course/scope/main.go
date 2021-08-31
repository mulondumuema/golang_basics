package main

import "myapp/packageone"

var myVar = "I am a package level variable for the main package" 
func main() {
	var blockVar = "I am a block variable for the main function"
	packageone.PrintMe(myVar, blockVar)
}

