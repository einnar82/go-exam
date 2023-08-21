# Setup Instructions for Go Exam

## Introduction

This document provides step-by-step instructions for setting up and running Go Exam, a Go-based HTTP server for handling file uploads and saving image metadata to a PostgreSQL database.

## Prerequisites

Before you begin, ensure you have the following:

-   Docker and Docker Compose installed on your system.
-   Basic familiarity with Docker and using the command line.

## Step 1: Clone the Repository

1.  Open your terminal.
    
2.  Navigate to the directory where you want to clone the repository.
    
3.  Run the following command to clone the repository:
    
    shCopy code
    
    `git clone github.com/einnar82/go-exam` 
    

## Step 2: Configure Environment Variables

1.  Navigate to the project's `config` folder.
    
2.  Open the `config.go` file.
    
3.  Update the following configuration values according to your preferences:
    
    -   `AuthToken`: Set the authorization token for the file uploads.
    -   `StorageFolderPath`: Set the path to the storage folder within the project.
    -   `DBConnectionString`: Set the PostgreSQL database connection string.

## Step 3: Run the Application

1.  In your terminal, navigate to the root directory of the cloned repository.
    
2.  Run the following command to start the application and PostgreSQL database:
    
    shCopy code
    
    `docker-compose up -d` 
    
    This will launch the application and expose it on port 8080.
    

## Step 4: Access the Application

1.  In your terminal, Run `go run main.go`
2.  Open your web browser.
3.  Visit `http://localhost:8080` to access the file upload form.

## Step 5: Using pgAdmin

1.  Open your web browser.
2.  Visit `http://localhost:5050`.
3.  Log in using the default email (`admin@example.com`) and password (`admin`).
4.  Add a new server connection:
    -   Host: `postgres`
    -   Port: `5432`
    -   Database name, username, and password as specified in the `docker-compose.yml` file.

## Step 6: Uploading Files and Metadata

1.  On the file upload form (`http://localhost:8080`), select an image file and provide the authorization token.
2.  Click the "Upload" button.
3.  The image file will be saved in the storage folder, and its metadata will be stored in the PostgreSQL database.

## Step 7: Cleaning Up

When you're done testing the application, you can clean up:

1.  In the terminal, navigate to the directory containing the `docker-compose.yml` file.
    
2.  Run the following command to stop and remove the containers:
    
    shCopy code
    
    `docker-compose down` 
    

## Conclusion

Congratulations! You've successfully set up and run Go Exam. Feel free to explore the uploaded files and metadata using the provided web interface and pgAdmin.
