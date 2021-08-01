package main

import "fmt"

func main()  {
	for i := 1; i<=100; i++{
		if i % 3 == 0 && i % 5 == 0{ 
            fmt.Println("FuzzBuzz")
		}else if i % 3 == 0 {
			fmt.Println("Fuzz")
		} else if i % 5 == 0{
           fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}