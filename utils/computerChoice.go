package utils

import (
	"math/rand"
	"time"
)

func GenerateComputerChoice() (choice string) {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(3)

	if randNum == 0 {
		choice = "rock"
	}

	if randNum == 1 {
		choice = "paper"
	}

	if randNum == 2 {
		choice = "scissors"
	}

	return
}
