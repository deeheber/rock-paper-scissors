package utils

import "github.com/AlecAivazis/survey/v2"

func PlayAgain() (playAgain bool) {
	prompt := &survey.Confirm{
		Message: "Would you like to play another round?",
		Default: true,
	}
	survey.AskOne(prompt, &playAgain)

	return
}
