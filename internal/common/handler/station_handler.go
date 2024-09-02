package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"metro-in/internal/common/constants"
	"metro-in/internal/common/middleware"
	"metro-in/internal/common/service"
)

// stationControllerImpl implementation for StationController
type stationControllerImpl struct {
	baseController BaseController
	stationService service.StationService
}

// StationController interface for subway lines routes
type StationController interface {
	RegisterRoutes(app *fiber.App)
	GetStationByName(*fiber.Ctx) error
}

// NewStationController constructor for StationController
func NewStationController(dbClient *gorm.DB) StationController {
	baseController := NewBaseController()
	return &stationControllerImpl{stationService: service.NewStationService(dbClient), baseController: baseController}
}

// RegisterRoutes set the http routes for the handler
func (h *stationControllerImpl) RegisterRoutes(app *fiber.App) {
	stationRouter := app.Group(constants.StationHandler)
	stationRouter.Get("/:name", middleware.AuthMiddleware, h.GetStationByName)
}

// GetStationByName godoc
func (h *stationControllerImpl) GetStationByName(c *fiber.Ctx) error {
	station, err := h.stationService.GetStationByName(c.Params(constants.Name))
	if err != nil {
		return h.baseController.RespondError(c, err)
	}

	return h.baseController.RespondSuccessWithBody(c, station)
}
