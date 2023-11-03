package mock

import (
	"metro-in/internal/common/customerrors"
	"metro-in/internal/common/entity"
	"metro-in/internal/common/repository"
	"metro-in/internal/common/repository/mock"
	"metro-in/internal/common/service"
)

// subwayLineServiceMock mocked implementation for service.SubwayLineService
type subwayLineServiceMockError struct {
	subwayLineRepository repository.SubwayLineRepository
}

// NewSubwayLineServiceMockError generate a mocked implementation of service.SubwayLineService
func NewSubwayLineServiceMockError() service.SubwayLineService {
	return &subwayLineServiceMockError{subwayLineRepository: mock.NewSubwayLineRepositoryMock()}
}

// GetAll godoc
func (c *subwayLineServiceMockError) GetAll() ([]entity.SubwayLine, error) {
	return []entity.SubwayLine{}, &customerrors.InternalServerError{}
}

// GetByID godoc
func (c *subwayLineServiceMockError) GetByID(id uint) (entity.SubwayLine, error) {
	return entity.SubwayLine{}, &customerrors.InternalServerError{}
}

// GetByCompanyID godoc
func (c *subwayLineServiceMockError) GetByCompanyID(companyID uint) ([]entity.SubwayLine, error) {
	return []entity.SubwayLine{}, &customerrors.InternalServerError{}
}
