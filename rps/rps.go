package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate computer choice
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(3)
	var computerChoice string

	if randNum == 0 {
		computerChoice = "rock"
	}

	if randNum == 1 {
		computerChoice = "paper"
	}

	if randNum == 2 {
		computerChoice ="scissors"
	}

	fmt.Println("Computer choice: ", computerChoice)
}
