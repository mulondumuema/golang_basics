package main

import "fmt"

func main() {
	a := [2][3]int{}
	

	fmt.Println(`We are creating a 2 by 3 array. Press 'Enter' once done entering each element`)
	fmt.Print("->")
	fmt.Println()
	
	for i := 0; i < 2; i++ {
		for j :=0 ; j < 3; j++ {
			fmt.Scanln(&a[i][j])
		}
	}

	fmt.Println(a)

	

}

