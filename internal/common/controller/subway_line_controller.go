package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"metro-in/internal/common/service"
)

// subwayLineControllerImpl implementation for SubwayLineController
type subwayLineControllerImpl struct {
	subwayService service.SubwayLineService
}

// SubwayLineController interface for subway lines routes
type SubwayLineController interface {
	GetAll(*fiber.Ctx) error
}

// NewSubwayLineController constructor for SubwayLineController
func NewSubwayLineController(dbClient *gorm.DB) SubwayLineController {
	return &subwayLineControllerImpl{subwayService: service.NewSubwayLineService(dbClient)}
}

// GetAll godoc
func (c *subwayLineControllerImpl) GetAll(*fiber.Ctx) error {
	return nil
}
