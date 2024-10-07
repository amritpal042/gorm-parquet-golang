package handlers

import (
	"archive/zip"
	"bytes"
	"fmt"
	"gorm-parquet-golang/internal/pagination"
	"log"
	"net/http"

	"github.com/xitongsys/parquet-go-source/writerfile"
	"github.com/xitongsys/parquet-go/writer"
	"gorm.io/gorm"
)

// Generic handler function that works with any GORM model to download ZIP with Parquet files
func GenericZipDownloadHandler[T any, U any](w http.ResponseWriter, r *http.Request, db *gorm.DB, applyFilters func(*gorm.DB) *gorm.DB, transformFunc func(T) U) {
	ctx := r.Context() // Get context for cancellation

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=batch_data.zip")
	w.Header().Set("Content-Type", "application/zip")

	// Create a ZIP writer that writes directly to the response writer
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close() // Ensure the ZIP writer is closed properly

	// Estimate recommended page size based on file size
	pageSize := 100 // Set a default page size to 100 for now

	page := 1
	writtenRecords := 0

	// Paginate through the records and write each page as a separate Parquet file to the ZIP archive
	for {
		select {
		case <-ctx.Done():
			log.Println("Request cancelled by client or server; stopping.")
			return
		default:
		}

		// Paginate and fetch the next batch of records
		var results []T
		_, _, err := pagination.Paginate(db, page, pageSize, &results, applyFilters)
		if err != nil {
			http.Error(w, "Failed to paginate data", http.StatusInternalServerError)
			log.Printf("Failed to paginate data: %v", err)
			return
		}

		// If no more rows, finish
		if len(results) == 0 {
			break
		}

		// Create a new entry for each Parquet file in the ZIP archive
		zipFileName := fmt.Sprintf("data_page_%d.parquet", page)
		zipEntry, err := zipWriter.Create(zipFileName)
		if err != nil {
			http.Error(w, "Failed to create ZIP entry", http.StatusInternalServerError)
			log.Printf("Failed to create ZIP entry: %v", err)
			return
		}

		// Create an in-memory buffer and parquet writer for the transformed records
		buffer := &bytes.Buffer{}
		pw, err := writer.NewParquetWriter(writerfile.NewWriterFile(buffer), new(U), 4)
		if err != nil {
			http.Error(w, "Failed to create Parquet writer", http.StatusInternalServerError)
			log.Printf("Failed to create Parquet writer: %v", err)
			return
		}

		// Transform results from type T to type U
		for _, record := range results {
			transformedRecord := transformFunc(record)
			if err := pw.Write(transformedRecord); err != nil {
				http.Error(w, "Failed to write record to Parquet", http.StatusInternalServerError)
				log.Printf("Failed to write record to Parquet: %v", err)
				return
			}
			writtenRecords++ // Increment the counter for each successfully written record
		}

		// Finalize the Parquet writer
		if err := pw.WriteStop(); err != nil {
			http.Error(w, "Failed to close Parquet writer", http.StatusInternalServerError)
			log.Printf("Failed to close Parquet writer: %v", err)
			return
		}

		// Write the buffer containing the Parquet file data to the ZIP entry
		_, err = zipEntry.Write(buffer.Bytes())
		if err != nil {
			http.Error(w, "Failed to write buffer to ZIP entry", http.StatusInternalServerError)
			log.Printf("Failed to write buffer to ZIP entry: %v", err)
			return
		}

		// Move to the next page
		page++
	}

	// Check if any records were written before stopping
	if writtenRecords == 0 {
		http.Error(w, "No records were written to the Parquet files", http.StatusInternalServerError)
		log.Println("No records were written to the Parquet files")
		return
	}

	// Finalize and close the ZIP writer
	if err := zipWriter.Close(); err != nil {
		http.Error(w, "Failed to finalize the ZIP archive", http.StatusInternalServerError)
		log.Printf("Failed to finalize the ZIP archive: %v", err)
		return
	}
}
