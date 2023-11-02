package repository

import "metro-in/internal/config/database"

// SubwayLineRepositoryImpl implementation for SubwayLineRepository
type SubwayLineRepositoryImpl struct {
	db database.Database
}

// SubwayLineRepository interface for subway lines at database
type SubwayLineRepository interface {
}

// NewSubwayLineRepository constructor for SubwayLineRepository
func NewSubwayLineRepository(dbClient *database.Database) SubwayLineRepository {
	return &SubwayLineRepositoryImpl{db: *dbClient}
}
