package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"metro-in/internal/common/entity"
	"metro-in/internal/common/errors"
	"os"
)

var context = "Database"

var tables = []interface{}{
	&entity.SubwayLineStation{},
	&entity.Station{},
	&entity.SubwayLine{},
}

// ConnectDb initiate db connection with the given .env variables
func ConnectDb() (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, errors.Error{Context: context, Message: "Failed to connect to database", Err: err}
	}

	err = autoMigrateTables(db)
	if err != nil {
		return nil, errors.Error{Context: context, Message: "Failed during migrations", Err: err}
	}

	return db, nil
}

func autoMigrateTables(db *gorm.DB) error {
	for _, table := range tables {
		if err := migrateTable(db, table); err != nil {
			return err
		}
	}
	return nil
}

func migrateTable(db *gorm.DB, entity interface{}) error {
	return db.AutoMigrate(entity)
}
