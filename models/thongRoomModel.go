package models

import "gorm.io/gorm"

type RomeType string

const (
	Male     RomeType = "MALE"
	Female   RomeType = "FEMALE"
	Disabled RomeType = "DISABLED"
)

type ThongRoom struct {
	gorm.Model
	RoomType    RomeType
	IsAvailable bool
	ThongID     uint
	Reserve     []Reserve
}
