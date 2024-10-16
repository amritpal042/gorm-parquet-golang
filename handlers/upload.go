package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"your_project/utils"
)

// UploadHandler handles the uploading of ZIP files
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form and retrieve the file
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("zipfile")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read ZIP file into memory
	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		http.Error(w, "Failed to read the uploaded file", http.StatusInternalServerError)
		return
	}

	// Process ZIP file in memory
	if err := utils.ProcessZipFile(buf.Bytes(), handler.Filename); err != nil {
		http.Error(w, "Failed to process ZIP file", http.StatusInternalServerError)
		return
	}

	// Successfully processed
	fmt.Fprintln(w, "Successfully processed all Parquet files from ZIP.")
}
