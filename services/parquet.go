package services

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/xitongsys/parquet-go-source/writerfile"
	"github.com/xitongsys/parquet-go/reader"
)

// ReadParquet reads a single Parquet file from the multipart file
func ReadParquet(file multipart.File) error {
	// Initialize the Parquet reader
	fr := writerfile.NewWriterFile(file)
	pr, err := reader.NewParquetReader(fr, nil, 4)
	if err != nil {
		return fmt.Errorf("failed to create parquet reader: %w", err)
	}
	defer pr.ReadStop()

	// Get the total number of rows
	numRows := int(pr.GetNumRows())
	log.Printf("Total rows in Parquet file: %d\n", numRows)

	// Read and process rows in batches
	batchSize := 10
	for i := 0; i < numRows; i += batchSize {
		readCount := utils.Min(batchSize, numRows-i)

		// Create a slice of maps for the batch
		rows := make([]map[string]interface{}, readCount)

		// Read the batch into the slice of maps
		if err := pr.Read(&rows); err != nil {
			return fmt.Errorf("failed to read batch: %w", err)
		}

		// Process each row in the batch
		for rowIndex, row := range rows {
			log.Printf("Row %d:\n", i+rowIndex)
			for colName, colValue := range row {
				log.Printf("  %s: %v\n", colName, colValue)
			}
		}
	}

	return nil
}
