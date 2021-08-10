package main

import (
	"bufio"
	"fmt"
	"os"
	"myapp/doctor"
	
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		userInput, _ := reader.ReadString('\n')

	    fmt.Println(userInput)
	}

	
}

