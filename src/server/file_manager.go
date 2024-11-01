package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	pb "github.com/vexxuh/magnetflow_filemanager/src/generated/src/protobuf"
)

// Server struct implementing the FileManager gRPC service
type Server struct {
	pb.UnimplementedFileManagerServer
}

// StartWatcher method to watch directory changes and send events over gRPC stream
func (s *Server) StartWatcher(req *pb.WatchRequest, stream pb.FileManager_StartWatcherServer) error {
	directory := req.GetDirectory()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(directory)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Started watching directory: %s\n", directory)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			fmt.Printf("Event detected: %s\n", event)
			_ = stream.Send(&pb.WatchResponse{
				Event:    event.Op.String(),
				FileName: event.Name,
			})
			PrintDirectoryTree(directory, "")
		case err, ok := <-watcher.Errors:
			if !ok {
				return err
			}
			log.Printf("Error: %s\n", err)
		}
	}
}

// UploadFile method to handle file upload via gRPC
func (s *Server) UploadFile(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	directory := req.GetDirectory()
	fileName := req.GetFileName()
	content := req.GetContent()

	filePath := filepath.Join(directory, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return &pb.UploadResponse{Message: "Error creating file"}, err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return &pb.UploadResponse{Message: "Error writing to file"}, err
	}

	fmt.Printf("File %s uploaded successfully to %s\n", fileName, directory)
	return &pb.UploadResponse{Message: fmt.Sprintf("File %s uploaded successfully", fileName)}, nil
}
