package handlers

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/einnar82/go-exam/storage"
)

type UploadHandler struct {
	authToken string
	storage   storage.Storage
	db        *sql.DB // Add a database connection

}

func NewUploadHandler(authToken string, storage storage.Storage, db *sql.DB) *UploadHandler {
	return &UploadHandler{
		authToken: authToken,
		storage:   storage,
		db:        db,
	}
}

func (h *UploadHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	auth := r.FormValue("auth")
	if auth != h.authToken {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	file, fileHeader, err := r.FormFile("data")
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Check if uploaded file is an image
	if !strings.HasPrefix(fileHeader.Header.Get("Content-Type"), "image/") {
		http.Error(w, "File is not an image", http.StatusForbidden)
		return
	}

	// Limit image size to 8MB
	fileSizeLimit := int64(8 * 1024 * 1024)
	if fileHeader.Size > fileSizeLimit {
		http.Error(w, "File size exceeds limit", http.StatusRequestEntityTooLarge)
		return
	}

	// Save the uploaded image and get the saved file path
	uploadedFilePath, err := h.storage.SaveFile(fileHeader.Filename, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Insert image metadata into the database
	imageMetadata := fmt.Sprintf("Content-Type: %s, Size: %d bytes, Path: %s", fileHeader.Header.Get("Content-Type"), fileHeader.Size, uploadedFilePath)
	if err := h.insertMetadata(imageMetadata, fileHeader, uploadedFilePath); err != nil {
		http.Error(w, "Error saving metadata", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "File uploaded successfully!")
}

func (h *UploadHandler) insertMetadata(metadata string, fileHeader *multipart.FileHeader, uploadedFilePath string) error {
	_, err := h.db.Exec(`
		INSERT INTO image_metadata (content_type, size_bytes, path) VALUES ($1, $2, $3)
	`, fileHeader.Header.Get("Content-Type"), fileHeader.Size, uploadedFilePath)
	return err
}
