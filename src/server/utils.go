package server

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// PrintDirectoryTree prints the structure of the given directory in a tree format.
func PrintDirectoryTree(path string, indent string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(indent + "+ " + file.Name())
			PrintDirectoryTree(filepath.Join(path, file.Name()), indent+"  ")
		} else {
			fmt.Println(indent + "- " + file.Name())
		}
	}
}
