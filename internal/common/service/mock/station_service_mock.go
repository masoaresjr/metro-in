package mock

import (
	"metro-in/internal/common/entity"
	"metro-in/internal/common/repository"
	"metro-in/internal/common/repository/mock"
	"metro-in/internal/common/service"
)

// stationServiceMock mocked implementation for service.StationService
type stationServiceMock struct {
	stationRepository repository.StationRepository
}

// NewStationServiceMock generate a mocked implementation of service.StationService
func NewStationServiceMock() service.StationService {
	return &stationServiceMock{ stationRepository: mock.NewStationRepositoryMock() }
}

func (s *stationServiceMock) GetStationByName(name string) (*entity.Station, error) {
	return &entity.Station{}, nil
}