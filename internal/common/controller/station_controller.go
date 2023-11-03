package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"metro-in/internal/common/service"
)

// StationControllerImpl implementation for StationController
type StationControllerImpl struct {
	subwayService service.StationService
}

// StationController interface for subway lines routes
type StationController interface {
	GetByName(*fiber.Ctx) error
}

// NewStationController constructor for StationController
func NewStationController(dbClient *gorm.DB) StationController {
	return &StationControllerImpl{subwayService: service.NewStationService(dbClient)}
}

// GetByName godoc
func (c *StationControllerImpl) GetByName(ctx *fiber.Ctx) error {

	name := ctx.Params("name")
	if name == "" {
		//return error
	}

	//get station

	return nil
}
