package optimizations

import (
	"bytes"
	"encoding/json"
)

// EstimateRowSize estimates the size of a row based on a sample and calculates the appropriate page size
func EstimateRowSize(sample interface{}) int {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(sample)
	if err != nil {
		return 0
	}
	return buffer.Len()
}

// AdjustPageSize adjusts the page size dynamically based on available memory and desired file size
func AdjustPageSize(sample interface{}, maxFileSizeMB int) int {
	rowSize := EstimateRowSize(sample)
	if rowSize == 0 {
		return 100 // Default page size
	}

	maxFileSizeBytes := maxFileSizeMB * 1024 * 1024
	return maxFileSizeBytes / rowSize
}
