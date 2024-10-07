package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect establishes a connection to the PostgreSQL database
func Connect() (*gorm.DB, error) {

	// docker run --name postgres147 -e POSTGRES_PASSWORD=mysecretpassword -d postgres -p 5432:5432
	//docker run --name postgres147 -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:14.7

	dsn := "host=localhost user=postgres password=mysecretpassword dbname=metrics_dashboard port=5432 sslmode=disable"

	// Set custom logger for GORM (optional)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 200, // Slow SQL threshold in milliseconds
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
