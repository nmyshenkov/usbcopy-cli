package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"
)

func main() {

	// get user who runs program
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	userVolumes, err := scanVolumes(currentUser.Uid)
	if err != nil {
		log.Fatal(err)
	}

	choose, err := getChosenVolume(userVolumes)
	if err != nil {
		log.Fatal(err)
	}

	// get current dir
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// get file list
	files, err := OSReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}
	// sort string slice by name
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	// create directory
	if err := os.MkdirAll(choose+filepath.Base(currentDir), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Print("Copy: ", file.Name, ":")
		if _, err := copyFile(currentDir+"/"+file.Name, choose+filepath.Base(currentDir)+"/"+file.Name); err != nil {
			fmt.Println(" \x1b[31merror: " + err.Error() + "\x1b[0m")
			continue
		}

		fmt.Println(" \x1b[32mdone.\x1b[0m")
	}
}
