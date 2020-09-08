package main

import (
	"errors"
	"fmt"
	"strconv"
)

func scanVolumes(uid string) ([]string, error) {
	// get usb devices TODO:: detect OS and read mount folder
	volumes, err := OSReadDir("/Volumes")
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
	return userVolumes[chosen-1], nil
}
