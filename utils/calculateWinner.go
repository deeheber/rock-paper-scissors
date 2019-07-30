package utils

import "fmt"

func CalculateWinner(computerChoice, userChoice string) (winner string) {
		if userChoice == "rock" {
		if computerChoice == "rock" {
			winner = "Tie"
			fmt.Println("Both picked rock, it's a tie!")
		}

		if computerChoice == "paper" {
			winner = "Computer"
			fmt.Println("Paper covers rock...computer wins!")
		}

		if computerChoice == "scissors" {
			winner = "User"
			fmt.Println("Rock crushes scissors...you win!")
		}
	}

	if userChoice == "paper" {
		if computerChoice == "rock" {
			winner = "User"
			fmt.Println("Paper covers rock...you win!")
		}

		if computerChoice == "paper" {
			winner = "Tie"
			fmt.Println("Both picked paper, it's a tie!")
		}

		if computerChoice == "scissors" {
			winner = "Computer"
			fmt.Println("Scissors cuts paper...computer wins!")
		}
	}

	if userChoice == "scissors" {
		if computerChoice == "rock" {
			winner = "Computer"
			fmt.Println("Rock crushes scissors...computer wins!")
		}

		if computerChoice == "paper" {
			winner = "User"
			fmt.Println("Scissors cuts paper...you win!")
		}

		if computerChoice == "scissors" {
			winner = "Tie"
			fmt.Println("Both picked scissors...it's a tie!")
		}
	}

	return
}
