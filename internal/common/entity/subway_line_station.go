package entity

// SubwayLineStation entity
type SubwayLineStation struct {
	SubwayLineID uint `gorm:"index"`
	StationID    uint `gorm:"index"`
}
