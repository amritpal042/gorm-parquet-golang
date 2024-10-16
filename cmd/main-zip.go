package main

// curl -X POST -F "zipfile=@/path/to/yourfile.zip" http://localhost:8080/upload
import (
	"gorm-parquet-golang/handlers"
	"log"
	"net/http"
)

func main() {
	// Set up the upload handler
	http.HandleFunc("/upload", handlers.UploadHandler)

	// Start the server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
