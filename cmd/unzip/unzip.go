package unzip

import (
	"archive/zip"
	"fmt"
	"os"
	"path"
)

func UnzipToDir(source, dest string) error {
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
	fmt.Println("unzipToDir done")
	return nil
}
