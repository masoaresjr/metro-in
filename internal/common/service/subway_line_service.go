package service

import (
	"gorm.io/gorm"
	"metro-in/internal/common/repository"
)

// SubwayLineServiceImpl implementation for SubwayLineService
type SubwayLineServiceImpl struct {
	subwayLineRepository repository.SubwayLineRepository
}

// SubwayLineService interface for subway line service
type SubwayLineService interface {
}

// NewSubwayLineService constructor for SubwayLineService
func NewSubwayLineService(dbClient *gorm.DB) SubwayLineService {
	return &SubwayLineServiceImpl{subwayLineRepository: repository.NewSubwayLineRepository(dbClient)}
}
