package cmd

import (
	"errors"
	"github.com/AthulMuralidhar/protobuff-installer/pkg/cmd/semvar"
	"go.uber.org/zap"
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

func promptGetProtocVersion(logger *zap.SugaredLogger, pc promptContent) semvar.SemVar {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		logger.Error("Prompt failed:", zap.Error(err))
		os.Exit(1)
	}

	semVarList := strings.Split(result, ".")
	var sm semvar.SemVar

	sm.Major, err = strconv.Atoi(semVarList[0])
	sm.Minor, err = strconv.Atoi(semVarList[1])

	return sm
}

func promptGetProtocInput(logger *zap.SugaredLogger, pc promptContent) bool {
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
		logger.Error("Prompt failed:", zap.Error(err))
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
