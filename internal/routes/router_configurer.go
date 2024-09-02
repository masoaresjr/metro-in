package routes

import "github.com/gofiber/fiber/v2"

// RouteConfigurer interface for router configuration
type RouteConfigurer interface {
	RegisterRoutes(app *fiber.App)
}
