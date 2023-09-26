package downloader

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
)

func DownloadAndCreateFile(sugar *zap.SugaredLogger, url string, cwd string) (*os.File, error) {
	//Create the file
	f, err := os.CreateTemp(cwd, "protoc.zip")
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

	return f, nil
}
