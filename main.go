package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang-rock-paper-scissors/utils"
)

var (
	playAgain  = "no"
	totalScore = map[string]int{"Computer": 0, "User": 0, "Tie": 0, "Rounds": 0}
)

func main() {
	userChoice := utils.GetUserChoice()

	computerChoice := utils.GenerateComputerChoice()

	fmt.Println("User choice: ", userChoice)
	fmt.Println("Computer choice: ", computerChoice)

	winner := utils.CalculateWinner(computerChoice, userChoice)
	totalScore[winner]++
	totalScore["Rounds"]++

	playAgain := utils.PlayAgain()

	if playAgain {
		main()
	} else {
		fmt.Println("Thanks for playing, here's the final results")
		fmt.Println("--------------------------------------------")
		fmt.Printf("User: %v\n", totalScore["User"])
		fmt.Printf("Computer: %v\n", totalScore["Computer"])
		fmt.Printf("Tie: %v\n", totalScore["Tie"])
		fmt.Printf("Number of games played: %v\n", totalScore["Rounds"])

		if (totalScore["User"] > totalScore["Computer"]) {
			color.Green("Final result: Congratualations, you won overall!")
		}

		if (totalScore["User"] < totalScore["Computer"]) {
			color.Red("Final result: The computer did better overall. Better luck next time.")
		}

		if (totalScore["User"] == totalScore["Computer"]) {
			color.Yellow("Final result: It was a tie overall.")
		}
	}
}
