package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"gorm-parquet-golang/services"
	"log"
	"strings"
)

// ProcessZipFile processes an in-memory ZIP file and reads each .parquet file
func ProcessZipFile(zipData []byte, zipFilename string) error {
	// Open the ZIP archive in memory
	zipReader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return fmt.Errorf("failed to read ZIP file %s: %w", zipFilename, err)
	}

	// Process each file in the ZIP
	for _, zipFile := range zipReader.File {
		log.Printf("Found file: %s\n", zipFile.Name)

		// Only process .parquet files
		if strings.HasSuffix(zipFile.Name, ".parquet") {
			// Open the parquet file within the ZIP archive
			parquetFile, err := zipFile.Open()
			if err != nil {
				return fmt.Errorf("failed to open parquet file: %w", err)
			}
			defer parquetFile.Close()

			// Read the parquet file
			err = services.ReadParquet(parquetFile)
			if err != nil {
				return fmt.Errorf("failed to read parquet file: %w", err)
			}
		}
	}

	return nil
}
