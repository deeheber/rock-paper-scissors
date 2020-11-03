package utils

import "github.com/fatih/color"

func CalculateWinner(computerChoice, userChoice string) (winner string) {
	if userChoice == "rock" {
		if computerChoice == "rock" {
			winner = "Tie"
			color.Yellow("Both picked rock, it's a tie!")
		}

		if computerChoice == "paper" {
			winner = "Computer"
			color.Red("Paper covers rock...computer wins!")
		}

		if computerChoice == "scissors" {
			winner = "User"
			color.Green("Rock crushes scissors...you win!")
		}
	}

	if userChoice == "paper" {
		if computerChoice == "rock" {
			winner = "User"
			color.Green("Paper covers rock...you win!")
		}

		if computerChoice == "paper" {
			winner = "Tie"
			color.Yellow("Both picked paper, it's a tie!")
		}

		if computerChoice == "scissors" {
			winner = "Computer"
			color.Red("Scissors cuts paper...computer wins!")
		}
	}

	if userChoice == "scissors" {
		if computerChoice == "rock" {
			winner = "Computer"
			color.Red("Rock crushes scissors...computer wins!")
		}

		if computerChoice == "paper" {
			winner = "User"
			color.Green("Scissors cuts paper...you win!")
		}

		if computerChoice == "scissors" {
			winner = "Tie"
			color.Yellow("Both picked scissors...it's a tie!")
		}
	}

	return
}
