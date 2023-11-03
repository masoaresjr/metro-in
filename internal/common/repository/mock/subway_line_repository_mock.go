package mock

import (
	"metro-in/internal/common/entity"
	"metro-in/internal/common/repository"
)

// subwayLineRepositoryMock mocked implementation for repository.SubwayLineRepository
type subwayLineRepositoryMock struct{}

// NewSubwayLineRepositoryMock generate a mocked implementation of repository.SubwayLineRepository
func NewSubwayLineRepositoryMock() repository.SubwayLineRepository {
	return &subwayLineRepositoryMock{}
}

// GetAll godoc
func (c *subwayLineRepositoryMock) GetAll() (subwayLines []entity.SubwayLine, err error) {
	return
}

// GetByID godoc
func (c *subwayLineRepositoryMock) GetByID(id uint) (subwayLine entity.SubwayLine, err error) {
	return
}

// GetByCompanyID godoc
func (c *subwayLineRepositoryMock) GetByCompanyID(companyID uint) (subwayLine []entity.SubwayLine, err error) {
	return
}
