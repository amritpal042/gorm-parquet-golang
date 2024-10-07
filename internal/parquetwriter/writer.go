package parquetwriter

import (
	"bytes"
	"net/http"

	"github.com/xitongsys/parquet-go-source/writerfile"
	"github.com/xitongsys/parquet-go/writer"
	"gorm.io/gorm"
)

// WriteToParquet writes data to a Parquet file with optimizations like buffering and connection handling
func WriteToParquet[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, pageSize int, results []T) error {
	ctx := r.Context()

	w.Header().Set("Content-Disposition", "attachment; filename=data.parquet")
	w.Header().Set("Content-Type", "application/octet-stream")

	buffer := &bytes.Buffer{}
	pw, err := writer.NewParquetWriter(writerfile.NewWriterFile(buffer), new(T), 4)
	if err != nil {
		return err
	}

	for _, record := range results {
		select {
		case <-ctx.Done():
			return nil
		default:
			if err := pw.Write(record); err != nil {
				return err
			}
		}
	}

	if _, err := w.Write(buffer.Bytes()); err != nil {
		return err
	}

	if err := pw.WriteStop(); err != nil {
		return err
	}

	return nil
}
