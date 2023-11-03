package mock

import "metro-in/internal/common/repository"

// subwayLineRepositoryMock mocked implementation for repository.SubwayLineRepository
type subwayLineRepositoryMock struct {}

// NewSubwayLineRepositoryMock generate a mocked implementation of repository.SubwayLineRepository
func NewSubwayLineRepositoryMock() repository.SubwayLineRepository {
	return &subwayLineRepositoryMock{}
}