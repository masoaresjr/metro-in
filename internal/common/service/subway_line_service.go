package service

import (
	"gorm.io/gorm"
	"metro-in/internal/common/repository"
)

// SubwayServiceImpl implementation for SubwayLineService
type SubwayServiceImpl struct {
	subwayRepository repository.SubwayLineRepository
}

// SubwayLineService interface for subway line service
type SubwayLineService interface {
}

// NewSubwayLineService constructor for SubwayLineService
func NewSubwayLineService(dbClient *gorm.DB) SubwayLineService {
	return &SubwayServiceImpl{subwayRepository: repository.NewSubwayLineRepository(dbClient)}
}
