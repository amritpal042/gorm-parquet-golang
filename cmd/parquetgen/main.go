package main

import (
	"gorm-parquet-golang/internal/models"
	"gorm-parquet-golang/internal/parquetgen"
	"log"
	"os"
)

func main() {
	// Define the GORM model you want to generate a Parquet struct for
	model := models.MetricsAggregatedPerHour{}

	// Generate the Parquet struct
	parquetStruct := parquetgen.GenerateParquetStruct(model, "MetricsAggregatedPerHour")

	// Define the output file for the generated Parquet struct
	outputFile := "internal/parquetmodels/metrics_aggregated_per_hour_parquet.go"

	// Write the generated Parquet struct to a Go file
	err := os.MkdirAll("internal/parquetmodels", os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	parquetgen.WriteToFile(outputFile, parquetStruct)

	log.Printf("Generated Parquet struct for MetricsAggregatedPerHour in %s\n", outputFile)
}
