package main

import (
	"metro-in/internal/config/database"
	"metro-in/internal/routes"
)

// StartApp runs database and initiate routes
func StartApp() {
	database.ConnectDb()

	routes.StartRoutes()
}
