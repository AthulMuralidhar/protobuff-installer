package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type promptContent struct {
	errorMsg string
	label    string
}

type semVar struct {
	major int
	minor int
	patch int
}

func (s semVar) String() string {
	return "v" + strconv.Itoa(s.major) + "." + strconv.Itoa(s.minor) + "." + strconv.Itoa(s.patch)
}

const PB_URL = "https://github.com/protocolbuffers/protobuf/releases"

// installCmd represents the new command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "main command the the install functionality",
	Long:  `main command the the install functionality`,
	Run: func(cmd *cobra.Command, args []string) {
		protocInstall()
	},
}
var templates = &promptui.PromptTemplates{
	Prompt:  "{{ . }} ",
	Valid:   "{{ . | green }} ",
	Invalid: "{{ . | red }} ",
	Success: "{{ . | bold }} ",
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

func protocInstall() {
	var sm semVar
	protocPrompt := promptContent{
		"I hope you know what you are doing here...",
		"Would you like to install protoc? (y/n)",
	}
	installProtoc := promptGetProtocInput(protocPrompt)
	if installProtoc {
		versionPrompt := promptContent{
			"You do understand semantic versioning, don't you?",
			"What version would you like installed? (semantic versioning pls)",
		}
		sm = promptGetProtocVersion(versionPrompt)

	}

	protocInstaller(installProtoc, sm)
}

func promptGetProtocVersion(pc promptContent) semVar {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		semVarList := strings.Split(input, ".")

		if len(semVarList) != 3 {
			return errors.New(pc.errorMsg)
		}
		_, err := strconv.Atoi(semVarList[0])
		if err != nil {
			return err
		}
		_, err = strconv.Atoi(semVarList[1])
		if err != nil {
			return err
		}
		_, err = strconv.Atoi(semVarList[2])
		if err != nil {
			return err
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
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	semVarList := strings.Split(result, ".")
	var sm semVar

	sm.major, err = strconv.Atoi(semVarList[0])
	sm.minor, err = strconv.Atoi(semVarList[1])
	sm.patch, err = strconv.Atoi(semVarList[2])

	return sm
}

func protocInstaller(installProtoc bool, sm semVar) {
	if !installProtoc {
		return
	}

	// https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-aarch_64.zip
	// https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip
	url := fmt.Sprintf("%s/download/%s/protoc-%s-linux-x86_64.zip", PB_URL, sm.String(), sm.String())

	fmt.Printf("url: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer resp.Body.Close()
	//fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	// Create the file
	out, err := os.Create("test.zip")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Printf("err: %s", err)
}
