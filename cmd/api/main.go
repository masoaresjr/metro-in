package main

import (
	_ "github.com/joho/godotenv/autoload"
	"metro-in/internal/config/database"
	"metro-in/internal/routes"
)

func main() { startApp() }

// StartApp runs database and initiate routes
func startApp() {
	routes.StartRoutes(database.ConnectDb())
}