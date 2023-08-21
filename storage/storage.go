package storage

import (
	"io"
	"os"
	"path/filepath"
)

type Storage interface {
	SaveFile(fileName string, file io.Reader) (string, error)
}

type LocalStorage struct {
	basePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		basePath: basePath,
	}
}

func (s *LocalStorage) SaveFile(fileName string, file io.Reader) (string, error) {
	// Generate a unique file name for the uploaded image
	uploadedFileName := fileName

	// Create and open the file in the storage folder
	uploadedFilePath := filepath.Join(s.basePath, uploadedFileName)
	uploadedFile, err := os.Create(uploadedFilePath)
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	// Copy the uploaded file's data into the storage file
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		return "", err
	}

	return uploadedFilePath, nil
}
