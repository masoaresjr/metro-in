package service

import (
	"gorm.io/gorm"
	"metro-in/internal/common/repository"
)

// subwayLineServiceImpl implementation for SubwayLineService
type subwayLineServiceImpl struct {
	subwayLineRepository repository.SubwayLineRepository
}

// SubwayLineService interface for subway line service
type SubwayLineService interface {
}

// NewSubwayLineService constructor for SubwayLineService
func NewSubwayLineService(dbClient *gorm.DB) SubwayLineService {
	return &subwayLineServiceImpl{subwayLineRepository: repository.NewSubwayLineRepository(dbClient)}
}
