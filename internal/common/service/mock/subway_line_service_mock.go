package mock

import (
	"gorm.io/gorm"
	"metro-in/internal/common/entity"
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

// GetAll godoc
func (c *subwayLineServiceMock) GetAll() ([]entity.SubwayLine, error) {
	return []entity.SubwayLine{
		{
			Model: gorm.Model{
				ID: 1,
			},
			CompanyID: 2,
			Name: "Vermelha",
		},
	}, nil
}

// GetByID godoc
func (c *subwayLineServiceMock) GetByID(id uint) (entity.SubwayLine, error) {
	return entity.SubwayLine{
		Model: gorm.Model{
			ID: 1,
		},
		CompanyID: 2,
		Name: "Vermelha",
	}, nil
}

// GetByCompanyID godoc
func (c *subwayLineServiceMock) GetByCompanyID(companyID uint) ([]entity.SubwayLine, error) {
	if companyID == 0 {

	}
	return []entity.SubwayLine{}, nil
}
