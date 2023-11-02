package repository

import (
	"gorm.io/gorm"
)

// SubwayLineRepositoryImpl implementation for SubwayLineRepository
type SubwayLineRepositoryImpl struct {
	db *gorm.DB
}

// SubwayLineRepository interface for subway lines at database
type SubwayLineRepository interface {
}

// NewSubwayLineRepository constructor for SubwayLineRepository
func NewSubwayLineRepository(dbClient *gorm.DB) SubwayLineRepository {
	return &SubwayLineRepositoryImpl{db: dbClient}
}
