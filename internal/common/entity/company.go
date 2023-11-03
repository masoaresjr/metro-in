package entity

import "gorm.io/gorm"

// Company entity
type Company struct {
	gorm.Model
	Name    	string
	Phone   	string
	SiteURL 	string
	ImageURL	string
}
