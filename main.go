package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/einnar82/go-exam/config"
	"github.com/einnar82/go-exam/handlers"
	"github.com/einnar82/go-exam/storage"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	config.LoadConfig()
	storageService := storage.NewLocalStorage(config.StorageFolderPath)

	db, err := sql.Open("postgres", config.DBConnectionString) // Replace with your PostgreSQL connection string
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()

	// Create the metadata table
	if err := createMetadataTable(db); err != nil {
		log.Fatal("Error creating metadata table:", err)
	}

	uploadHandler := handlers.NewUploadHandler(config.AuthToken, storageService, db)

	http.HandleFunc("/", handlers.FormHandler(config.AuthToken))
	http.HandleFunc("/upload", uploadHandler.HandleUpload)
	http.ListenAndServe(":8080", nil)
}

func createMetadataTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS image_metadata (
			id SERIAL PRIMARY KEY,
			content_type TEXT,
			size_bytes BIGINT,
			path TEXT
		)
	`)
	return err
}
