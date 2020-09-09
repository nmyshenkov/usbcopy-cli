package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// VolumesPath contains path to volume place
const VolumesPath = "/Volumes"

func scanVolumes(uid string) ([]string, error) {
	// get usb devices TODO:: detect OS and read mount folder
	volumes, err := OSReadDir(VolumesPath)
	if err != nil {
		return nil, err
	}

	var userVolumes []string

	for _, volume := range volumes {
		// if file is not owned current user - continue
		if strconv.Itoa(int(volume.Owner)) != uid {
			continue
		}

		// needs just dirs
		if !volume.IsDir {
			continue
		}

		userVolumes = append(userVolumes, volume.Name)
	}

	if len(userVolumes) == 0 {
		return nil, errors.New("user's devices not found")
	}

	return userVolumes, nil
}

func getChosenVolume(userVolumes []string) (string, error) {
	fmt.Println("External device list:")
	for i, volume := range userVolumes {
		fmt.Println(i+1, "-", volume)
	}
	var chosen int

	fmt.Print("Choose external device: ")

	if _, err := fmt.Scan(&chosen); err != nil {
		return "", err
	}
	return VolumesPath + "/" + userVolumes[chosen-1] + "/", nil
}

func copyFile(sourceFile, dstFile string) (int64, error) {
	sourceFileStat, err := os.Stat(sourceFile)
	if err != nil {
		return 0, err
	}

	// we can copy just files
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", sourceFile)
	}

	source, err := os.Open(sourceFile)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dstFile)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
