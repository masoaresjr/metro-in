package entity

import "gorm.io/gorm"

// Station entity
type Station struct {
	gorm.Model
	Name        string       `json:"name"`
	SubwayLines []SubwayLine `gorm:"many2many:subway_line_stations;"`
}
