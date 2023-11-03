package mock

import (
	"metro-in/internal/common/entity"
	"metro-in/internal/common/repository"
)

// stationRepositoryMock mocked implementation for repository.StationRepository
type stationRepositoryMock struct {}

// NewStationRepositoryMock generate a mocked implementation of repository.StationRepository
func NewStationRepositoryMock() repository.StationRepository {
	return &stationRepositoryMock{}
}

func (s *stationRepositoryMock) GetStationByName(name string) (result entity.Station, err error) {
	return result, nil
}