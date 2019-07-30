package main

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"math/rand"
	"time"
)

type Scoreboard struct {
	User, Computer int
}

var (
	playAgain  = "no"
	totalScore = Scoreboard{}
)

func main() {
	userChoice := ""
	selectPrompt := &survey.Select{
		Message: "Make a selection:",
		Options: []string{"rock", "paper", "scissors"},
	}
	survey.AskOne(selectPrompt, &userChoice)

	// Generate computer choice
	computerChoice := GenerateChoice()

	fmt.Println("User choice: ", userChoice)
	fmt.Println("Computer choice: ", computerChoice)

	// Compare answers and declare winner
	if userChoice == "rock" {
		if computerChoice == "rock" {
			fmt.Println("Both picked rock, it's a tie!")
		}

		if computerChoice == "paper" {
			totalScore.Computer++
			fmt.Println("Paper covers rock...computer wins!")
		}

		if computerChoice == "scissors" {
			totalScore.User++
			fmt.Println("Rock crushes scissors...you win!")
		}
	}

	if userChoice == "paper" {
		if computerChoice == "rock" {
			totalScore.User++
			fmt.Println("Paper covers rock...you win!")
		}

		if computerChoice == "paper" {
			fmt.Println("Both picked paper, it's a tie!")
		}

		if computerChoice == "scissors" {
			totalScore.Computer++
			fmt.Println("Scissors cuts paper...computer wins!")
		}
	}

	if userChoice == "scissors" {
		if computerChoice == "rock" {
			totalScore.Computer++
			fmt.Println("Rock crushes scissors...computer wins!")
		}

		if computerChoice == "paper" {
			totalScore.User++
			fmt.Println("Scissors cuts paper...you win!")
		}

		if computerChoice == "scissors" {
			fmt.Println("Both picked scissors...it's a tie!")
		}
	}

	// Ask if you'd like to play again
	playAgainPrompt := &survey.Select{
		Message: "Would you like to play another round?",
		Options: []string{"yes", "no"},
	}
	survey.AskOne(playAgainPrompt, &playAgain)

	if playAgain == "yes" {
		main()
	} else {
		fmt.Println("Thanks for playing, here's the final results")
		fmt.Println("--------------------------------------------")
		fmt.Printf("User: %v\n", totalScore.User)
		fmt.Printf("Computer: %v\n", totalScore.Computer)
	}
}

func GenerateChoice() (choice string) {
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
