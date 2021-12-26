package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// define all the constant variable
const (
	MAIN_FOLDER_NAME         = "parent"
	SUB_FOLDER_NAME          = "child"
	EMPTY_FOLDER_NAME        = "empty-folder"
	FILE_UNDER_MAIN_FOLDER   = "parent.txt"
	FILE_UNDER_SUB_FOLDER    = "child.txt"
	MAIN_FOLDER_FILE_CONTENT = "Hello this is from main folder file"
	SUB_FOLDER_FILE_CONTENT  = "Hello this is from sub folder file"
	ZIP_FILE_NAME            = "golang-folder.zip"
)

// check for error and stop the execution
func checkForError(err error) {
	if err != nil {
		fmt.Println("Error - ", err)
		os.Exit(1)
	}
}

// create folder if not exists
func createFolder(folderPath string) {
	// check if folder exists
	_, err := os.Stat(folderPath)

	if os.IsNotExist(err) {
		// create folder if no folder
		err = os.Mkdir(folderPath, 0755)
		checkForError(err)
		fmt.Printf("%s created successfully\n", folderPath)
	} else {
		fmt.Printf("%s - already exists\n", folderPath)
		os.Exit(1)
	}
}

// create file if not exists

func createFile(filePath, content string) {

	fmt.Println(filePath)
	// check if file exists
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		// create file if no file
		f, err := os.Create(filePath)
		checkForError(err)
		defer f.Close()

		_, err = f.WriteString(content)
		checkForError(err)
		fmt.Printf("%s created successfully\n", filePath)
	} else {
		fmt.Printf("%s - already exists\n", filePath)
		os.Exit(1)
	}
}

// Main function
func main() {

	var targetFilePaths []string
	// create a main dir
	createFolder(MAIN_FOLDER_NAME)

	// create a sub folder
	createFolder(filepath.Join(MAIN_FOLDER_NAME, SUB_FOLDER_NAME))

	// create a empty sub folder
	createFolder(filepath.Join(MAIN_FOLDER_NAME, EMPTY_FOLDER_NAME))

	// create a file under main folder
	createFile(filepath.Join(MAIN_FOLDER_NAME, FILE_UNDER_MAIN_FOLDER), MAIN_FOLDER_FILE_CONTENT)

	// create a file under sub folder
	createFile(filepath.Join(MAIN_FOLDER_NAME, SUB_FOLDER_NAME, FILE_UNDER_SUB_FOLDER), SUB_FOLDER_FILE_CONTENT)

	// get filepaths in all folders
	err := filepath.Walk(MAIN_FOLDER_NAME, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path)
		if info.IsDir() {
			targetFilePaths = append(targetFilePaths, path+"/")
			return nil
		}
		targetFilePaths = append(targetFilePaths, path)
		return nil
	})
	checkForError(err)

	// zip file logic starts here
	ZipFile, err := os.Create(ZIP_FILE_NAME)
	checkForError(err)
	defer ZipFile.Close()

	zipWriter := zip.NewWriter(ZipFile)
	defer zipWriter.Close()

	for _, targetFilePath := range targetFilePaths {

		//fmt.Println(targetFilePath)

		fileInfo, err := os.Stat(targetFilePath)
		checkForError(err)

		if fileInfo.IsDir() {
			_, err := zipWriter.Create(targetFilePath)
			checkForError(err)
		} else {
			file, err := os.Open(targetFilePath)
			checkForError(err)
			defer file.Close()

			w, err := zipWriter.Create(targetFilePath)
			checkForError(err)
			_, err = io.Copy(w, file)
			checkForError(err)
		}

	}

}
