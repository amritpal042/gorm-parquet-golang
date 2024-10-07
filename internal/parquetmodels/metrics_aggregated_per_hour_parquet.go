package parquetmodels

type MetricsAggregatedPerHourParquet struct {
	EntityID               string   `parquet:"name=EntityID, type=BYTE_ARRAY, convertedtype=UTF8"`
	MetricID               string   `parquet:"name=MetricID, type=BYTE_ARRAY, convertedtype=UTF8"`
	AggregateIntervalStart int64    `parquet:"name=AggregateIntervalStart, type=INT64"`
	AggregateIntervalEnd   int64    `parquet:"name=AggregateIntervalEnd, type=INT64"`
	AvgValue               *float64 `parquet:"name=AvgValue, type=DOUBLE"`
	MinValue               *float64 `parquet:"name=MinValue, type=DOUBLE"`
	MaxValue               *float64 `parquet:"name=MaxValue, type=DOUBLE"`
	Count                  *int64   `parquet:"name=Count, type=INT64"`
	OldestValue            *float64 `parquet:"name=OldestValue, type=DOUBLE"`
	LatestValue            *float64 `parquet:"name=LatestValue, type=DOUBLE"`
	Trend                  *float64 `parquet:"name=Trend, type=DOUBLE"`
	IsProcessed            bool     `parquet:"name=IsProcessed, type=BOOLEAN"`
	InsertedAt             int64    `parquet:"name=InsertedAt, type=INT64"`
	TotalValue             *float64 `parquet:"name=TotalValue, type=DOUBLE"`
}
