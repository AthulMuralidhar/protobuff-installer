package cmd

import (
	"archive/zip"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

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

func init() {
	rootCmd.AddCommand(installCmd)
}

func protocInstall() {
	log.Println("protocInstall called")
	installProtoc := promptGetProtocInput(protocPrompt)
	if !installProtoc {
		log.Println("protocInstall: user does not want to install protoc, exiting...")
		return
	}
	sm := promptGetProtocVersion(versionPrompt)
	installPath := promptGetInstallPath(installPathPrompt)

	if err := protocInstaller( sm, installPath); err != nil {
		log.Fatalln(err)
	}
}

func protocInstaller(sm semVar, installPath string) error {
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
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Create the file
	out, err := os.CreateTemp("", "protoc.zip")
	if err != nil {
		return fmt.Errorf("error during creating temp zip file: %w", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error during writing to temp zip file: %w", err)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error during getting user home dir: %w", err)
	}
	// basically this:  unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
	err = unzip("test.zip", homeDir+"/.protoc")
	if err != nil {
		return fmt.Errorf("error during unzipping: %w", err)
	}

	// finally add to path
	// TODO: make this simple
	// FIXME: the problem here is that the protoc is not an executable and the export does not work due to this
	// fmt.Println("finally copy paste this into your terminal:")
	// fmt.Println(`export PATH="$PATH:$HOME/.protoc/bin"`)

	log.Println("protocInstall: done")
	return nil
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
