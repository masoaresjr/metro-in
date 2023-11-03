package mock

import (
	"metro-in/internal/common/repository"
	"metro-in/internal/common/repository/mock"
	"metro-in/internal/common/service"
)

// subwayLineServiceMock mocked implementation for service.SubwayLineService
type subwayLineServiceMock struct {
	subwayLineRepository repository.SubwayLineRepository
}

// NewSubwayLineServiceMock generate a mocked implementation of service.SubwayLineService
func NewSubwayLineServiceMock() service.SubwayLineService {
	return &subwayLineServiceMock{subwayLineRepository: mock.NewSubwayLineRepositoryMock()}
}
