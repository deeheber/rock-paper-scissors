package utils

import "github.com/AlecAivazis/survey/v2"

func GetUserChoice() (choice string) {
	selectPrompt := &survey.Select{
		Message: "Make a selection:",
		Options: []string{"rock", "paper", "scissors"},
	}
	survey.AskOne(selectPrompt, &choice)

	return
}
