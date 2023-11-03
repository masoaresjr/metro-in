package entity

import "gorm.io/gorm"

// SubwayLine entity
type SubwayLine struct {
	gorm.Model
	Name      string `json:"name" gorm:"text"`
	Number    uint   `json:"number"`
	CompanyID uint   `json:"company_id"`
	Company   Company
	Stations  []SubwayLineStation `gorm:"-"`
}
