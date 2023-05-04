package storage

import (
	"encoding/json"
	"os"
	"sync"
)

type FileManager[T any] struct {
	fd *os.File
	mu sync.Mutex
}

func (f *FileManager[T]) StartSession(session string) error {
	file, err := os.Create(session + ".json")
	if err != nil {
		return err
	}
	file.WriteString("[")
	f.fd = file
	return nil
}

func (f *FileManager[T]) WriteToDisk(model T) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	bytes, err := json.Marshal(model)
	if err != nil {
		return err
	}

	suffix := []byte(", \n")
	bytes = append(bytes, suffix...)

	_, err = f.fd.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
