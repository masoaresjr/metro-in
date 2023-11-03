package service

import (
	"gorm.io/gorm"
	"metro-in/internal/common/entity"
	"metro-in/internal/common/repository"
)

// stationServiceImpl implementation for StationService
type stationServiceImpl struct {
	stationRepository repository.StationRepository
}

// StationService interface for station service
type StationService interface {
	GetStationByName(name string) (*entity.Station, error)
}

// NewStationService constructor for StationService
func NewStationService(dbClient *gorm.DB) StationService {
	return &stationServiceImpl{stationRepository: repository.NewStationRepository(dbClient)}
}

func (s *stationServiceImpl) GetStationByName(name string) (*entity.Station, error) {
	return &entity.Station{}, nil
}