package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// check for any error
func check(err error) {
	if err != nil {
		fmt.Printf("Error Happened %s \n", err)
		os.Exit(1)
	}
}

// Main function
func main() {

	// Get the user input - target folder needs to be organized
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Which folder do you want to organize? - ")
	scanner.Scan()

	targetFolder := scanner.Text()

	// check the folder exists or not
	_, err := os.Stat(targetFolder)
	if os.IsNotExist(err) {
		fmt.Println("Folder does not exist.")
		os.Exit(1)
	} else {
		// create default folders such as Images, Music, Docs, Others, Videos
		createDefaultFolders(targetFolder)

		//Oraganize folders
		organizeFolder(targetFolder)
	}

}

// function to create default folders such as Images, Music, Docs, Others, Videos
func createDefaultFolders(targetFolder string) {
	defaultFolders := []string{"Music", "Videos", "Docs", "Images", "Others"}

	for _, folder := range defaultFolders {
		_, err := os.Stat(folder)
		if os.IsNotExist(err) {
			os.Mkdir(filepath.Join(targetFolder, folder), 0755)
		}
	}
}

// function to Oraganize folders
func organizeFolder(targetFolder string) {
	// read the dir
	filesAndFolders, err := os.ReadDir(targetFolder)
	check(err)

	// to track how many files moved
	noOfFiles := 0

	for _, filesAndFolder := range filesAndFolders {
		// check for files
		if !filesAndFolder.IsDir() {
			fileInfo, err := filesAndFolder.Info()
			check(err)

			//get the file full path
			oldPath := filepath.Join(targetFolder, fileInfo.Name())
			fileExt := filepath.Ext(oldPath)

			// switch case to move files based on ext
			switch fileExt {
			case ".png", ".jpg", ".jpeg":
				newPath := filepath.Join(targetFolder, "Images", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".mp4", ".mov", ".avi", ".amv":
				newPath := filepath.Join(targetFolder, "Videos", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".pdf", ".docx", ".csv", ".xlsx":
				newPath := filepath.Join(targetFolder, "Docs", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".mp3", ".wav", ".aac":
				newPath := filepath.Join(targetFolder, "Music", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			default:
				newPath := filepath.Join(targetFolder, "Others", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			}
		}
	}

	// print how many files moved
	if noOfFiles > 0 {
		fmt.Printf("%v number of files moved\n", noOfFiles)
	} else {
		fmt.Printf("No files moved")
	}
}
