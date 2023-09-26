package cmd

import (
	"fmt"
	"github.com/AthulMuralidhar/protobuff-installer/pkg/cmd/downloader"
	"github.com/AthulMuralidhar/protobuff-installer/pkg/cmd/semvar"
	"github.com/AthulMuralidhar/protobuff-installer/pkg/cmd/unzip"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"os"
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

	if err := protocInstaller(sm, installPath); err != nil {
		log.Fatalln(err)
	}
}

func protocInstaller(sm semvar.SemVar, installPath string) error {
	// https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-aarch_64.zip
	// https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	url := sm.ToProtocURL()

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("err during getcwd call: %s", err)
	}
	f, err := downloader.DownloadAndCreateFile(sugar, url, cwd)

	defer os.Remove(f.Name())

	// basically this:  unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
	err = os.Mkdir("protoc", 0777)
	if err != nil {
		return fmt.Errorf("error during making protoc directory: %w", err)
	}

	err = unzip.UnzipToDir(f.Name(), cwd+"protoc")
	if err != nil {
		return fmt.Errorf("error during unzipping: %w", err)
	}

	sugar.Info("protocInstall: done")
	return nil
}
