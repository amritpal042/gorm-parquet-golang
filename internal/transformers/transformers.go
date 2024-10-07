package transformers

import (
	"gorm-parquet-golang/internal/models"

	"gorm-parquet-golang/internal/parquetmodels"
)

// TransformMetricsToParquet transforms MetricsAggregatedPerHour to MetricsAggregatedPerHourParquet
func TransformMetricsToParquet(metrics models.MetricsAggregatedPerHour) parquetmodels.MetricsAggregatedPerHourParquet {
	return parquetmodels.MetricsAggregatedPerHourParquet{
		EntityID:               metrics.EntityID,
		MetricID:               metrics.MetricID,
		AggregateIntervalStart: metrics.AggregateIntervalStart.UnixMilli(),
		AggregateIntervalEnd:   metrics.AggregateIntervalEnd.UnixMilli(),
		AvgValue:               metrics.AvgValue,
		MinValue:               metrics.MinValue,
		MaxValue:               metrics.MaxValue,
		Count:                  metrics.Count,
		OldestValue:            metrics.OldestValue,
		LatestValue:            metrics.LatestValue,
		Trend:                  metrics.Trend,
		IsProcessed:            metrics.IsProcessed,
		InsertedAt:             metrics.InsertedAt.UnixMilli(),
		TotalValue:             metrics.TotalValue,
	}
}
