package entity

import "gorm.io/gorm"

// SubwayLine entity
type SubwayLine struct {
	gorm.Model
	Name     string              `json:"name" gorm:"text"`
	Stations []SubwayLineStation `gorm:"many2many:subway_line_stations;"`
}
