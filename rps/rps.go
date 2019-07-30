package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/AlecAivazis/survey"
)

func main() {
	// Get user choice
	userChoice := ""
	prompt := &survey.Select{
			Message: "Make a selection:",
			Options: []string{"rock", "paper", "scissors"},
	}
	survey.AskOne(prompt, &userChoice)

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
	fmt.Println("User choice: ", userChoice)
	fmt.Println("Computer choice: ", computerChoice)

	// Compare answers and declare winner
	if userChoice == "rock" {
		if computerChoice == "rock" {
			fmt.Println("Both picked rock, it's a tie!")
		}

		if computerChoice == "paper" {
			fmt.Println("Paper covers rock...computer wins!")
		}

		if computerChoice == "scissors" {
			fmt.Println("Rock crushes scissors...you win!")
		}
	}

	if userChoice == "paper" {
		if computerChoice == "rock" {
			fmt.Println("Paper covers rock...you win!")
		}

		if computerChoice == "paper" {
			fmt.Println("Both picked paper, it's a tie!")
		}

		if computerChoice == "scissors" {
			fmt.Println("Scissors cuts paper...computer wins!")
		}
	}

	if userChoice == "scissors" {
		if computerChoice == "rock" {
			fmt.Println("Rock crushes scissors...computer wins!")
		}

		if computerChoice == "paper" {
			fmt.Println("Scissors cuts paper...you win!")
		}

		if computerChoice == "scissors" {
			fmt.Println("Both picked scissors...it's a tie!")
		}
	}
}
