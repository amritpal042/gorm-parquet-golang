package models

import "time"

// MetricsAggregatedPerHour is the GORM model that stores hourly metrics data.
type MetricsAggregatedPerHour struct {
	EntityID               string    `gorm:"column:entity_id;type:uuid"`
	MetricID               string    `gorm:"column:metric_id;type:uuid"`
	AggregateIntervalStart time.Time `gorm:"column:aggregate_interval_start"`
	AggregateIntervalEnd   time.Time `gorm:"column:aggregate_interval_end"`
	AvgValue               *float64  `gorm:"column:avg_value"`
	MinValue               *float64  `gorm:"column:min_value"`
	MaxValue               *float64  `gorm:"column:max_value"`
	Count                  *int64    `gorm:"column:count"`
	OldestValue            *float64  `gorm:"column:oldest_value"`
	LatestValue            *float64  `gorm:"column=latest_value"`
	Trend                  *float64  `gorm:"column:trend"`
	IsProcessed            bool      `gorm:"column:is_processed"`
	InsertedAt             time.Time `gorm:"column:inserted_at"`
	TotalValue             *float64  `gorm:"column:total_value"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (MetricsAggregatedPerHour) TableName() string {
	return "metrics_aggregated_per_hour"
}
