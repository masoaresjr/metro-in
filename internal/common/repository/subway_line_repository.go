package repository

import (
	"gorm.io/gorm"
)

// subwayLineRepositoryImpl implementation for SubwayLineRepository
type subwayLineRepositoryImpl struct {
	db *gorm.DB
}

// SubwayLineRepository interface for subway lines at database
type SubwayLineRepository interface {
}

// NewSubwayLineRepository constructor for SubwayLineRepository
func NewSubwayLineRepository(dbClient *gorm.DB) SubwayLineRepository {
	return &subwayLineRepositoryImpl{db: dbClient}
}
