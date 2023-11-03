package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"metro-in/internal/common/service"
)

// stationControllerImpl implementation for StationController
type stationControllerImpl struct {
	stationService service.StationService
}

// StationController interface for subway lines routes
type StationController interface {
	GetByName(*fiber.Ctx) error
}

// NewStationController constructor for StationController
func NewStationController(dbClient *gorm.DB) StationController {
	return &stationControllerImpl{stationService: service.NewStationService(dbClient)}
}

// GetByName godoc
func (c *stationControllerImpl) GetByName(ctx *fiber.Ctx) error {

	station, err := c.stationService.GetStationByName(ctx.Params("name"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error: %s", err.Error()))
	}

	err = ctx.JSON(station)

	return err
}
