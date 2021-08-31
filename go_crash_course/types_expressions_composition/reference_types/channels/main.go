package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

// // function are used to run goroutines
// func main() {
// 	go doSomething("Hello, world") // this is exucuted as a goroutine, i.e it runs independently as its own main function. It is fired off, runs in the background and the program keeps executing (This is called concurrency).
// 	// so for example if you have a go routine that sends an email after a request fo invoice, the goroutine runs in the background and once the invoice is requested, the goroutine send the email. So to speak with goroutines, we use channels 

// 	fmt.Println("This is another message") // this gets printed before the goroutine
// 	for {
// 		//  do nothing
// 	}
// }

// func doSomething(s string) {
// 	until := 0
// 	for {
// 		fmt.Println("s is", s)
// 		until = until + 1
// 		if until >= 5 {
// 			break
// 		}
// 	}
// }

// lets write a goroutine with a channel

var keyPressChan chan rune  // creating a channel variable. A channel can only use one data type. Rune is a single character that is used to build a string. After you have declared the variable this way, you cannot use it unless you the make keyword like we used for maps

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")

	_ = keyboard.Open()

	defer func ()  {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		// using the channel to send the char information

		keyPressChan <- char // note the arrow direction is for to send something to a channel
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan // to receive something from a channel
		fmt.Println("You pressed", string(key))
	}
}

// however the obove channel will block until something is passed
// channels are useful to pass information from one program to another, in goroutines
// channels are one of those powerful features of go