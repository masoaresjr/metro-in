package main

import (
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/joho/godotenv/autoload"
	"metro-in/internal/config/database"
	"metro-in/internal/routes"
)

func main() { startApp() }

// StartApp runs database and initiate routes
func startApp() {
	db, err := database.ConnectDb()
	if err != nil {
		log.Info(err.Error())
		return
	}

	if err = database.RunMigrations(db); err != nil {
		log.Info(err.Error())
		return
	}

	routes.StartRoutes(db)
}
