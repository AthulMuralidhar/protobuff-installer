package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type promptContent struct {
	errorMsg string
	label    string
}


var protocPrompt = promptContent{
	"I hope you know what you are doing here...",
	"Would you like to install protoc? (y/n)",
}

var versionPrompt = promptContent{
	"You do understand semantic versioning, don't you?",
	"What version would you like installed? (semantic versioning pls)",
}

var installPathPrompt = promptContent{
	"",
	"Where would you like to install the protoc compiler? (Default is . )",
}

func promptGetProtocInput(pc promptContent) bool {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		switch input {
		case "y":
			return nil
		case "n":
			return nil

		}

		return errors.New(pc.errorMsg)
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	switch result {
	case "y":
		return true
	case "n":
		return false
	}
	return false
}