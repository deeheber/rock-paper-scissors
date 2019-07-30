package utils

import "github.com/AlecAivazis/survey"

func PlayAgain() bool {
	playAgain := ""

	prompt := &survey.Select{
		Message: "Would you like to play another round?",
		Options: []string{"yes", "no"},
	}
	survey.AskOne(prompt, &playAgain)

	return playAgain == "yes"
}
