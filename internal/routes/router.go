package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"metro-in/internal/common/controller"
	"os"
)

var subwayLineController controller.SubwayLineController
var stationController controller.StationController

// StartRoutes start http routes and listen to then
func StartRoutes(dbClient *gorm.DB) {

	subwayLineController = controller.NewSubwayLineController(dbClient)
	stationController = controller.NewStationController(dbClient)

	app := fiber.New()

	startSubwayRoutes(app)
	startUserRoutes(app)
	startLineRoutes(app)
	startStationRoutes(app)

	err := app.Listen(":" + os.Getenv("APP_PORT"))

	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

func authMiddleware(c *fiber.Ctx) error {

	//TODO validar na session
	isUserLoggedIn := true

	if !isUserLoggedIn {
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

	stationRouter.Get("/:name", authMiddleware, stationController.GetStationByName)
}
