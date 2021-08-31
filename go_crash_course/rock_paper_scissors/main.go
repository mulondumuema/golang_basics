package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(2)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose ROCK")
	case PAPER:
		fmt.Println("Computer chose PAPER")
	case SCISSORS:
		fmt.Println("Computer chose SCISSORS")
	default:
	}

	fmt.Println("Player chose", playerChoice)
	
	fmt.Println()
	if playerValue == computerValue {
		fmt.Println("It's a draw")
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				fmt.Println("Computer Wins!")
			} else {
				fmt.Println("Player Wins!")
			}
		case PAPER:
			if computerValue == SCISSORS {
				fmt.Println("Computer Wins!")
			} else {
				fmt.Println("Player Wins!")
			}
		case SCISSORS:
			if computerValue == ROCK {
				fmt.Println("Computer Wins!")
			} else {
				fmt.Println("Player Wins!")
			}
		default:
			fmt.Println("Invalid choice!")
		}
	}

}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
