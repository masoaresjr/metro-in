package service

import (
	"gorm.io/gorm"
	"metro-in/internal/common/repository"
)

// StationServiceImpl implementation for StationService
type StationServiceImpl struct {
	stationRepository repository.StationRepository
}

// StationService interface for station service
type StationService interface {
}

// NewStationService constructor for StationService
func NewStationService(dbClient *gorm.DB) StationService {
	return &StationServiceImpl{stationRepository: repository.NewStationRepository(dbClient)}
}
