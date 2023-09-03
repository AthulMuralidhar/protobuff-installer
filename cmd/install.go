package cmd

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
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
	//patch int
}

func (s semVar) String() string {
	//return "v" + strconv.Itoa(s.major) + "." + strconv.Itoa(s.minor) + "." + strconv.Itoa(s.patch)
	return "v" + strconv.Itoa(s.major) + "." + strconv.Itoa(s.minor)
}

// TODO: make cleanup func to remove zip file and the rest

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
	if !installProtoc {
		return
	}
	versionPrompt := promptContent{
		"You do understand semantic versioning, don't you?",
		"What version would you like installed? (semantic versioning pls)",
	}

	sm = promptGetProtocVersion(versionPrompt)

	installPathPrompt := promptContent{
		"",
		"Where would you like to install the protoc compiler? (Default is . )",
	}
	installPath := promptGetInstallPath(installPathPrompt)

	protocInstaller(installProtoc, sm, installPath)
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

func protocInstaller(installProtoc bool, sm semVar, installPath string) {
	if !installProtoc {
		return
	}

	// https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-aarch_64.zip
	// https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip

	// TODO: fix url making, the below works, but the hand made one does not
	// FIXME
	//url := fmt.Sprintf("%s/download/%s/protoc-%s-linux-x86_64.zip", PB_URL, sm.String(), sm.String())
	url := "https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip"

	fmt.Printf("url: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	// Create the file
	out, err := os.CreateTemp("", "protoc.zip")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	// basically this:  unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
	err = unzip("test.zip", homeDir+"/.protoc")
	if err != nil {
		panic(err)
	}

	// finally add to path
	// TODO: make this simple
	// FIXME: the problem here is that the protoc is not an executable and the export does not work due to this
	fmt.Println("finally copy paste this into your terminal:")
	fmt.Println(`export PATH="$PATH:$HOME/.protoc/bin"`)

	fmt.Println("done")
}

func unzip(source, dest string) error {
	read, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer read.Close()
	for _, file := range read.File {
		if file.Mode().IsDir() {
			continue
		}
		open, err := file.Open()
		if err != nil {
			fmt.Println("cannot open file")
			return err
		}
		name := path.Join(dest, file.Name)
		os.MkdirAll(path.Dir(name), 0777)
		create, err := os.Create(name)
		if err != nil {

			fmt.Println("cannot create file")
			return err
		}
		defer create.Close()
		create.ReadFrom(open)
	}
	fmt.Println("unzip done")
	return nil
}
