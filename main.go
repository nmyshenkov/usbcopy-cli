package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {

	// get user who runs program
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// get usb devices
	volumes, err := OSReadDir("/Volumes")
	if err != nil {
		log.Fatal(err)
	}

	var userVolumes []string

	for _, volume := range volumes {
		// if file is not owned current user - continue
		if strconv.Itoa(int(volume.Owner)) != currentUser.Uid {
			continue
		}
		userVolumes = append(userVolumes, volume.Name)
	}

	if len(userVolumes) == 0 {
		log.Fatal("User's devices not found")
	}

	// get current dir
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentDir)

	// get file list
	files, err := OSReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}
	// sort string slice by name
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	fmt.Println(files)

}
