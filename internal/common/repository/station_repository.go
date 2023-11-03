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
	GetStationByName(name string) (result entity.Station, err error)
}

// NewStationRepository constructor for StationRepository
func NewStationRepository(dbClient *gorm.DB) StationRepository {
	return &stationRepositoryImpl{db: dbClient}
}

func (r *stationRepositoryImpl) GetStationByName(name string) (result entity.Station, err error) {
	err = r.db.Where("name = ?", name).First(&result).Error
	return result, err
}
