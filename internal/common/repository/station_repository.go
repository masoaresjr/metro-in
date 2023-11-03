package repository

import (
	"gorm.io/gorm"
)

// StationRepositoryImpl implementation for StationRepository
type StationRepositoryImpl struct {
	db *gorm.DB
}

// StationRepository interface for stations at database
type StationRepository interface {
}

// NewStationRepository constructor for StationRepository
func NewStationRepository(dbClient *gorm.DB) StationRepository {
	return &StationRepositoryImpl{db: dbClient}
}
