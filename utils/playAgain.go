package utils

import "github.com/AlecAivazis/survey"

func PlayAgain() (playAgain bool) {
	prompt := &survey.Confirm{
		Message: "Would you like to play another round?",
		Default: true,
	}
	survey.AskOne(prompt, &playAgain)

	return
}
