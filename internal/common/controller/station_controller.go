package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"metro-in/internal/common/service"
)

// stationControllerImpl implementation for StationController
type stationControllerImpl struct {
	baseController BaseController
	stationService service.StationService
}

// StationController interface for subway lines routes
type StationController interface {
	GetStationByName(*fiber.Ctx) error
}

// NewStationController constructor for StationController
func NewStationController(dbClient *gorm.DB) StationController {
	baseController := NewBaseController()
	return &stationControllerImpl{stationService: service.NewStationService(dbClient), baseController: baseController}
}

// GetStationByName godoc
func (c *stationControllerImpl) GetStationByName(ctx *fiber.Ctx) error {

	station, err := c.stationService.GetStationByName(ctx.Params("name"))
	if err != nil {
		return c.baseController.RespondError(ctx, err)
	}

	return c.baseController.RespondSuccessWithBody(ctx, station)
}
