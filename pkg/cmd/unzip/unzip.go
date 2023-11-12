package unzip

import (
	"archive/zip"
	"go.uber.org/zap"
	"os"
	"path"
)

func ToDir(logger *zap.SugaredLogger, source, dest string) error {
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
			logger.Error("cannot open file")
			return err
		}
		name := path.Join(dest, file.Name)
		os.MkdirAll(path.Dir(name), 0777)
		create, err := os.Create(name)
		if err != nil {

			logger.Error("cannot create file")
			return err
		}
		defer create.Close()
		create.ReadFrom(open)
	}
	logger.Info("unzipToDir done")
	return nil
}
