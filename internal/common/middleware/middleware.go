package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	//TODO validar na session
	isUserLoggedIn := true

	if !isUserLoggedIn {
		return c.Status(fiber.StatusUnauthorized).SendString("Acesso não autorizado")
	}

	return c.Next()
}
