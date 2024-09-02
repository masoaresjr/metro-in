package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"metro-in/internal/common/handler"
	"os"
)

// StartRoutes start http routes and listen to then
func StartRoutes(dbClient *gorm.DB) {

	var handlersRouters = []RouteConfigurer{
		handler.NewSubwayLineController(dbClient),
		handler.NewStationController(dbClient),
	}

	app := fiber.New()

	for _, routeConfigurer := range handlersRouters {
		routeConfigurer.RegisterRoutes(app)
	}

	if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
