package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gorm-parquet-golang/internal/pagination"

	"gorm.io/gorm"
)

// GenericHandler handles paginated queries and responds with JSON
func GenericHandler[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, applyFilters func(*gorm.DB) *gorm.DB) {
	page := 1
	pageSize := 10

	if r.URL.Query().Get("page") != "" {
		if p, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil {
			page = p
		}
	}

	if r.URL.Query().Get("pageSize") != "" {
		if ps, err := strconv.Atoi(r.URL.Query().Get("pageSize")); err == nil {
			pageSize = ps
		}
	}

	var results []T
	_, totalRows, err := pagination.Paginate(db, page, pageSize, &results, applyFilters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"page":      page,
		"pageSize":  pageSize,
		"totalRows": totalRows,
		"results":   results,
	})
}

// Paginate is a generic pagination function that works with any GORM model

func estimateRowSizeAndPageSize[T any](db *gorm.DB, model T, maxFileSizeMB int) (int, error) {
	// Retrieve a small sample of data from the table
	var sample []T
	err := db.Limit(100).Find(&sample).Error
	if err != nil {
		return 0, fmt.Errorf("failed to sample data from the table: %w", err)
	}

	// Estimate row size based on JSON encoding (simplified approach)
	var sampleBuffer bytes.Buffer
	encoder := json.NewEncoder(&sampleBuffer)
	for _, row := range sample {
		err := encoder.Encode(row)
		if err != nil {
			return 0, fmt.Errorf("failed to encode row for size estimation: %w", err)
		}
	}

	// Calculate average row size
	averageRowSize := sampleBuffer.Len() / len(sample)

	// Calculate page size based on maximum file size
	maxFileSizeBytes := maxFileSizeMB * 1024 * 1024
	recommendedPageSize := maxFileSizeBytes / averageRowSize

	return recommendedPageSize, nil
}
