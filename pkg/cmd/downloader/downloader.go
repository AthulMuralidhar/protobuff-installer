package downloader

import (
	"fmt"
	"github.com/AthulMuralidhar/protobuff-installer/pkg/cmd/semvar"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
)

func DownloadAndCreateFile(sugar *zap.SugaredLogger, sm semvar.SemVar, url string, cwd string) (*os.File, error) {
	sm = checkIfValid(sugar, sm)

	f, err := os.CreateTemp(cwd, "protoc"+sm.String()+".zip")
	if err != nil {
		return nil, fmt.Errorf("error during creating temp zip file: %w", err)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("err during http.Get: %s", err)
	}
	defer resp.Body.Close()
	sugar.Debugf("response status: %s", resp.Status)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	//Write the body to file
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error during writing to temp zip file: %w", err)
	}

	// check if zip file is empty
	fileInfo, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("error during getting file info: %w", err)
	}

	if fileInfo.Size() == 0 {
		return nil, fmt.Errorf("zip file is empty")
	}
	return f, nil
}

func checkIfValid(sugar *zap.SugaredLogger, sm semvar.SemVar) semvar.SemVar {
	// check if the current sem var exists on releases
	// if the given one is higher than the latest release, revert to latest release sem var and return
	resp, err := http.Get("https://github.com/protocolbuffers/protobuf/releases")
	if err != nil {
		sugar.Error(err)
		return sm
	}
	if resp == nil {
		sugar.Error("reponse from checkIfValid is nil")
		return sm
	}
	if resp.StatusCode != http.StatusOK {
		sugar.Error("returned response is not 200")
		return sm
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	return sm
}
