package routes

import "github.com/gofiber/fiber/v2"

type RouteConfigurer interface {
	RegisterRoutes(app *fiber.App)
}
