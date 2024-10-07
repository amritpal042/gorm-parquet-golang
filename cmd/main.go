package main

import (
	"log"
	"net/http"
	"time"

	"gorm-parquet-golang/internal/db"
	"gorm-parquet-golang/internal/handlers"
	"gorm-parquet-golang/internal/models"
	"gorm-parquet-golang/internal/parquetmodels"
	"gorm-parquet-golang/internal/transformers"

	"gorm.io/gorm"
)

func main() {
	// Set up the PostgreSQL connection
	//docker run --name postgres147 -e POSTGRES_PASSWORD=

	connection, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Your application logic here
	log.Println("Connected to the database successfully!")

	// HTTP handler for paginated metrics data
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		handlers.GenericHandler[models.MetricsAggregatedPerHour](w, r, connection, func(query *gorm.DB) *gorm.DB {
			// Add any filters if needed
			return query
		})
	})

	http.HandleFunc("/metrics-zip", func(w http.ResponseWriter, r *http.Request) {
		// Define how to apply filters for this model
		applyFilters := func(query *gorm.DB) *gorm.DB {
			if r.URL.Query().Get("startDate") != "" && r.URL.Query().Get("endDate") != "" {
				startDate, err := time.Parse("2006-01-02", r.URL.Query().Get("startDate"))
				if err != nil {
					http.Error(w, "Invalid startDate format", http.StatusBadRequest)
					return query
				}
				endDate, err := time.Parse("2006-01-02", r.URL.Query().Get("endDate"))
				if err != nil {
					http.Error(w, "Invalid endDate format", http.StatusBadRequest)
					return query
				}
				query = query.Where("aggregate_interval_start >= ? AND aggregate_interval_end <= ?", startDate, endDate)
			}
			return query
		}

		// Call the generic ZIP download handler for this specific model
		handlers.GenericZipDownloadHandler[models.MetricsAggregatedPerHour, parquetmodels.MetricsAggregatedPerHourParquet](w, r, connection, applyFilters, transformers.TransformMetricsToParquet)
	})

	// Start the HTTP server
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
