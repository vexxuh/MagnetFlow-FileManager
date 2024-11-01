package server

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

// StartDirectoryWatcher sets up a watcher on the given directory and logs events.
func StartDirectoryWatcher(directory string, watcher *fsnotify.Watcher) {
	err := watcher.Add(directory)
	if err != nil {
		log.Fatal("Failed to add watcher:", err)
	}
	fmt.Printf("Started watching directory: %s\n", directory)
}
