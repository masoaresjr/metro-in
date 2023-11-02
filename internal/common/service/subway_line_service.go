package service

import (
	"metro-in/internal/common/repository"
	"metro-in/internal/config/database"
)

// SubwayServiceImpl implementation for SubwayLineService
type SubwayServiceImpl struct {
	subwayRepository repository.SubwayLineRepository
}

// SubwayLineService interface for subway line service
type SubwayLineService interface {
}

// NewSubwayLineService constructor for SubwayLineService
func NewSubwayLineService(dbClient *database.Database) SubwayLineService {
	return &SubwayServiceImpl{subwayRepository: repository.NewSubwayLineRepository(dbClient)}
}
