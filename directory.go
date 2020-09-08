package main

import (
	"os"
	"syscall"
)

type File struct {
	Name  string
	Owner uint32
	IsDir bool
}

// OSReadDir get all files from directory
func OSReadDir(root string) ([]File, error) {
	var files []File
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	_ = f.Close()
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {

		fileSysInfo, ok := file.Sys().(*syscall.Stat_t)
		if !ok {
			continue
		}

		files = append(files, File{
			Name:  file.Name(),
			Owner: fileSysInfo.Uid,
			IsDir: file.Mode().IsDir(),
		})
	}
	return files, nil
}
