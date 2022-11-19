package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	ZIP_FILE_NAME    = "example.zip"
	MAIN_FOLDER_NAME = "parent"
)

type fileMeta struct {
	Path  string
	IsDir bool
}

func main() {
	var files []fileMeta
	err := filepath.Walk(MAIN_FOLDER_NAME, func(path string, info os.FileInfo, err error) error {
		files = append(files, fileMeta{Path: path, IsDir: info.IsDir()})
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	z, err := os.Create(ZIP_FILE_NAME)
	if err != nil {
		log.Fatalln(err)
	}
	defer z.Close()

	zw := zip.NewWriter(z)
	defer zw.Close()

	for _, f := range files {
		path := f.Path
		if f.IsDir {
			path = fmt.Sprintf("%s%c", path, os.PathSeparator)
		}

		w, err := zw.Create(path)
		if err != nil {
			log.Fatalln(err)
		}

		if !f.IsDir {
			file, err := os.Open(f.Path)
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()

			if _, err = io.Copy(w, file); err != nil {
				log.Fatalln(err)
			}
		}
	}
}
