package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"metro-in/internal/common/controller"
)

var subwayLineController controller.SubwayLineController

// StartRoutes start http routes and listen to then
func StartRoutes(dbClient *gorm.DB) {

	subwayLineController = controller.NewSubwayLineController(dbClient)

	app := fiber.New()

	startSubwayRoutes(app)
	startUserRoutes(app)
	startLineRoutes(app)
	startStationRoutes(app)

	err := app.Listen(":3000")

	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

func authMiddleware(c *fiber.Ctx) error {

	//TODO validar na session
	isUserLoggedIn := true

	if !isUserLoggedIn {
		// O usuário não está logado, então bloqueie a solicitação com um código de status 401
		return c.Status(fiber.StatusUnauthorized).SendString("Acesso não autorizado")
	}

	return c.Next()
}

func startSubwayRoutes(app *fiber.App) {

	subwayRouter := app.Group("subway")

	subwayRouter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
}

func startUserRoutes(app *fiber.App) {
	userRouter := app.Group("user")

	userRouter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
}

func startLineRoutes(app *fiber.App) {
	lineRouter := app.Group("line")

	lineRouter.Get("/", authMiddleware, subwayLineController.GetAll)
}

func startStationRoutes(app *fiber.App) {
	stationRouter := app.Group("station")

	stationRouter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
}
