package controller

import (
	"github.com/gofiber/fiber/v2"
	"metro-in/internal/common/service"
	"metro-in/internal/config/database"
)

// SubwayLineControllerImpl implementation for SubwayLineController
type SubwayLineControllerImpl struct {
	subwayService service.SubwayLineService
}

// SubwayLineController interface for subway lines routes
type SubwayLineController interface {
	GetAll(*fiber.Ctx) error
}

// NewSubwayLineController constructor for SubwayLineController
func NewSubwayLineController(dbClient *database.Database) SubwayLineController {
	return &SubwayLineControllerImpl{subwayService: service.NewSubwayLineService(dbClient)}
}

// GetAll godoc
func (c *SubwayLineControllerImpl) GetAll(*fiber.Ctx) error {
	return nil
}
