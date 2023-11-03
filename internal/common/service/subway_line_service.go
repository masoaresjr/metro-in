package service

import (
	"gorm.io/gorm"
	"metro-in/internal/common/customerrors"
	"metro-in/internal/common/entity"
	"metro-in/internal/common/repository"
)

// subwayLineServiceImpl implementation for SubwayLineService
type subwayLineServiceImpl struct {
	subwayLineRepository repository.SubwayLineRepository
}

// SubwayLineService interface for subway line service
type SubwayLineService interface {
	GetAll() ([]entity.SubwayLine, error)
	GetByID(id uint) (entity.SubwayLine, error)
	GetByCompanyID(companyID uint) ([]entity.SubwayLine, error)
}

// NewSubwayLineService constructor for SubwayLineService
func NewSubwayLineService(dbClient *gorm.DB) SubwayLineService {
	return &subwayLineServiceImpl{subwayLineRepository: repository.NewSubwayLineRepository(dbClient)}
}

// GetAll godoc
func (c *subwayLineServiceImpl) GetAll() (subwayLines []entity.SubwayLine, err error) {
	return c.subwayLineRepository.GetAll()
}

// GetByID godoc
func (c *subwayLineServiceImpl) GetByID(id uint) (entity.SubwayLine, error) {
	if id == 0 {
		return entity.SubwayLine{}, &customerrors.EmptyParameterError{ParameterName: "id"}
	}
	return c.subwayLineRepository.GetByID(id)
}

// GetByCompanyID godoc
func (c *subwayLineServiceImpl) GetByCompanyID(companyID uint) (subwayLines []entity.SubwayLine, err error) {
	if companyID == 0 {
		return []entity.SubwayLine{}, &customerrors.EmptyParameterError{ParameterName: "company_id"}
	}
	return c.subwayLineRepository.GetByCompanyID(companyID)
}
