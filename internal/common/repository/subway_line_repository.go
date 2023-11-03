package repository

import (
	"gorm.io/gorm"
	"metro-in/internal/common/customerrors"
	"metro-in/internal/common/entity"
)

// subwayLineRepositoryImpl implementation for SubwayLineRepository
type subwayLineRepositoryImpl struct {
	db *gorm.DB
}

// SubwayLineRepository interface for subway lines at database
type SubwayLineRepository interface {
	GetAll() (subwayLines []entity.SubwayLine, err error)
	GetByID(id uint) (subwayLine entity.SubwayLine, err error)
	GetByCompanyID(companyID uint) (subwayLine []entity.SubwayLine, err error)
}

// NewSubwayLineRepository constructor for SubwayLineRepository
func NewSubwayLineRepository(dbClient *gorm.DB) SubwayLineRepository {
	return &subwayLineRepositoryImpl{db: dbClient}
}

// GetAll godoc
func (c *subwayLineRepositoryImpl) GetAll() (subwayLines []entity.SubwayLine, err error) {
	err = c.db.Find(&subwayLines).Error
	return
}

// GetByID godoc
func (c *subwayLineRepositoryImpl) GetByID(id uint) (subwayLine entity.SubwayLine, err error) {
	if id == 0 {
		err = &customerrors.EmptyParameterError{ParameterName: "id"}
	}
	err = c.db.Where("id = ?", id).First(&subwayLine).Error
	return
}

// GetByCompanyID godoc
func (c *subwayLineRepositoryImpl) GetByCompanyID(companyID uint) (subwayLine []entity.SubwayLine, err error) {
	if companyID == 0 {
		err = &customerrors.EmptyParameterError{ParameterName: "company_id"}
	}
	return []entity.SubwayLine{}, nil
}
