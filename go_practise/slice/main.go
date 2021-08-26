package main

import "fmt"

func main() {
	var rows,col byte
	a := [2][3]int{}
	

	fmt.Println(`We are creating a 2 by 3 array. Press 'Enter' once done entering each element`)
	fmt.Print("->")
	fmt.Println()
	fmt.Println("Enter number of rows")
	fmt.Scanln(&rows)
	fmt.Println("Enter number of col")
	fmt.Scanln(&col)

	for i := 0; i < rows; i++ {
		for j :=0 ; j < 3; j++ {
			fmt.Scanln(&a[i][j])
		}
	}

	fmt.Println(a)

	

}

