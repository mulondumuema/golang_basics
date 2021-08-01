package main 

import "fmt"

func main(){
	fmt.Println("Enter temperature in Fahrenheit: ")
	var faren float64
	fmt.Scanf("%f", &faren)
    degrees := ((faren - 32)*5/9)
	fmt.Println(degrees)
}