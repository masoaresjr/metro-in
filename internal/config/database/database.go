package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// Database struct to the database connections
type Database struct {
	db *gorm.DB
}

// ConnectDb initiate db connection with the given .env variables
func ConnectDb() Database {
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
		log.Fatal("Failed to connect to database. \n", err)
	}

	//log.Println("running migrations")
	//if err := db.AutoMigrate(&entity.Fact{}); err != nil {
	//	log.Fatal(err)
	//	return
	//}

	return Database{
		db: db,
	}
}
