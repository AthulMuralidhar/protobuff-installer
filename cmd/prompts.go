package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func promptGetInstallPath(prompt promptContent) string {
	return "."
}

func promptGetProtocVersion(pc promptContent) semVar {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		semVarList := strings.Split(input, ".")

		fmt.Printf("semvarlist: %v", semVarList)

		// FIXME
		//if len(semVarList) != 3 {
		//	return errors.New(pc.errorMsg)
		// FIXME
		//}
		//_, err := strconv.Atoi(semVarList[0])
		//if err != nil {
		//	return err
		// FIXME
		//}
		//_, err = strconv.Atoi(semVarList[1])
		//if err != nil {
		//	return err
		// FIXME
		//}
		//_, err = strconv.Atoi(semVarList[2])
		//if err != nil {
		//	return err
		//}
		return nil
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

	semVarList := strings.Split(result, ".")
	var sm semVar

	sm.major, err = strconv.Atoi(semVarList[0])
	sm.minor, err = strconv.Atoi(semVarList[1])
	// FIXME
	//sm.patch, err = strconv.Atoi(semVarList[2])

	return sm
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
