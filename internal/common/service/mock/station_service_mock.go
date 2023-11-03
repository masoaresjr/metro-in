package mock

import (
	"errors"
	"metro-in/internal/common/custom_errors"
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
	if name == "corinthians-itaquera" {
		return &entity.Station{}, nil
	}
	return nil, custom_errors.Error{ Context: "Mock", Message: "Mocked Error", Err: errors.New("mocked")}
}