package models

import "gorm.io/gorm"

type Thong struct {
	gorm.Model
	IsOpen     bool
	Name       string
	Ratting    int
	IsMale     bool
	IsFemale   bool
	IsDisabled bool
	Lat        string
	Long       string
	ThongRoom  []ThongRoom
}
