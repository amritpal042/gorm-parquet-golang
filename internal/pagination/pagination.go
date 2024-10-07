package pagination

import (
	"gorm.io/gorm"
)

// Paginate is a generic function that handles paginated queries for any GORM model
func Paginate[T any](db *gorm.DB, page, pageSize int, result *[]T, applyFilters func(*gorm.DB) *gorm.DB) (*gorm.DB, int64, error) {
	var totalRows int64

	// Apply filters
	query := applyFilters(db)

	// Count total rows
	err := query.Model(result).Count(&totalRows).Error
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	err = query.Limit(pageSize).Offset(offset).Find(result).Error
	if err != nil {
		return nil, 0, err
	}

	return query, totalRows, nil
}
