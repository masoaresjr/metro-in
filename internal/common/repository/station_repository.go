package repository

import (
	"gorm.io/gorm"
	"metro-in/internal/common/entity"
)

// stationRepositoryImpl implementation for StationRepository
type stationRepositoryImpl struct {
	db *gorm.DB
}

// StationRepository interface for stations at database
type StationRepository interface {
	GetStationByName(name string) (*entity.Station, error)
}

// NewStationRepository constructor for StationRepository
func NewStationRepository(dbClient *gorm.DB) StationRepository {
	return &stationRepositoryImpl{db: dbClient}
}

func (s *stationRepositoryImpl) GetStationByName(name string) (*entity.Station, error) {
	return &entity.Station{}, nil
}
