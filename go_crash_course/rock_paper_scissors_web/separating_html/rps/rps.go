package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
	// PLAYERWINS   = 1
	// COMPUTERWINS = 2
	// DRAW         = 3
)

type Round struct {
	// Winner         int    `json:"winner"`
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

var winMessage = []string{
	"Good job!",
	"Nice work!",
	"You should buy a lottery ticket",
}

var loseMessage = []string{
	"Too bad!",
	"Try again!",
	"This is just not your day.",
}

var drawMessage = []string {
	"Great minds think alike.",
	"Uh oh. Try again.",
	"Nobody wins, but you can try again.",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	
	// winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	// generate a random number from 0 - 2, which we will use to pick a random message
	messageInt := rand.Intn(2)
	// declare a variable to hold the message
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		// winner = DRAW
		// populate message from drawMessage
		message = drawMessage[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		// winner = PLAYERWINS
		// populate message from winMessage
		message = winMessage[messageInt]
	} else {
		roundResult = "Computer wins!"
		// winner = COMPUTERWINS
		// populate message from loseMessage
		message = loseMessage[messageInt]
	}

	var result Round
	// result.Winner = winner
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
