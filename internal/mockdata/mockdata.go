package mockdata

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm-parquet-golang/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Generate random float numbers between a range
func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Generate random integer between a range
func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// InsertDynamicData inserts 50 random records into the metrics_aggregated_per_hour table
func InsertDynamicData(db *gorm.DB) error {
	for i := 0; i < 50; i++ {
		// Generate random data
		entityID := uuid.New().String()
		metricID := uuid.New().String()
		startTime := time.Now().Add(time.Duration(-i) * time.Hour)
		endTime := startTime.Add(1 * time.Hour)

		avgValue := randomFloat(30.0, 80.0)
		minValue := randomFloat(20.0, avgValue)
		maxValue := randomFloat(avgValue, 100.0)
		count := int64(randomInt(50, 150))
		oldestValue := randomFloat(minValue, avgValue)
		latestValue := randomFloat(avgValue, maxValue)
		trend := randomFloat(-5.0, 5.0)
		isProcessed := rand.Intn(2) == 1
		insertedAt := time.Now()
		totalValue := avgValue * float64(count)

		// Insert the record
		record := models.MetricsAggregatedPerHour{
			EntityID:               entityID,
			MetricID:               metricID,
			AggregateIntervalStart: startTime,
			AggregateIntervalEnd:   endTime,
			AvgValue:               &avgValue,
			MinValue:               &minValue,
			MaxValue:               &maxValue,
			Count:                  &count,
			OldestValue:            &oldestValue,
			LatestValue:            &latestValue,
			Trend:                  &trend,
			IsProcessed:            isProcessed,
			InsertedAt:             insertedAt,
			TotalValue:             &totalValue,
		}

		if err := db.Create(&record).Error; err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}
	return nil
}

func SeedData(db *gorm.DB) {
	err := InsertDynamicData(db)
	if err != nil {
		log.Fatalf("failed to seed mock data: %v", err)
	}
	log.Println("50 mock records inserted successfully.")
}
