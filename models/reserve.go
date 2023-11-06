package models

import "gorm.io/gorm"

type Status string

const (
	Reserved Status = "RESERVED"
	Using    Status = "USING"
	Success  Status = "SUCCESS"
)

type Reserve struct {
	gorm.Model
	Status      Status
	UserID      uint
	ThongRoomID uint
}
