package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// installCmd represents the new command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "main command the the install functionality",
	Long:  `main command the the install functionality`,
	Run: func(cmd *cobra.Command, args []string) {
		protocInstall()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	rootCmd.AddCommand(installCmd)
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

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
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

	//fmt.Printf("Input: %s\n", result)

	//return errors.New("dont know wtf happened here")
	return false
}

func protocInstall() {
	protocPrompt := promptContent{
		"I hope you know what you are doing here...",
		"Would you like to install protoc? (y/n)",
	}
	installProtoc := promptGetProtocInput(protocPrompt)
	if installProtoc {
		versionPrompt := promptContent{
			"You do understand symantic versioning, don't you?",
			"What version would you like installed? (symantic versioning pls)",
		}
		version := promptGetProtocVersion(versionPrompt)

	}

	protocInstaller(installProtoc, version)
}

func promptGetProtocVersion(prompt promptContent) interface{} {

}

func protocInstaller(word bool) {

}
