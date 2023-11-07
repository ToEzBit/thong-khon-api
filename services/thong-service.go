package services

import (
	"thong-khon-api/config"
	"thong-khon-api/models"
)

func CreateThong(thong models.Thong) (*models.Thong, error) {

	thongData := models.Thong{
		IsOpen:     thong.IsOpen,
		Name:       thong.Name,
		Ratting:    thong.Ratting,
		IsMale:     thong.IsMale,
		IsFemale:   thong.IsFemale,
		IsDisabled: thong.IsDisabled,
		Lat:        thong.Lat,
		Long:       thong.Long,
	}

	if err := config.DB.Create(&thongData).Error; err != nil {
		return nil, err
	}

	return &thongData, nil

}

func CreateThongRoom(thongRoom models.ThongRoom) (*models.ThongRoom, error) {
	thongRoomData := models.ThongRoom{
		ThongID:     thongRoom.ThongID,
		RoomType:    thongRoom.RoomType,
		IsAvailable: thongRoom.IsAvailable,
	}

	if err := config.DB.Create(&thongRoomData).Error; err != nil {
		return nil, err
	}

	return &thongRoomData, nil
}

func GetAllThong() (*[]models.Thong, error) {

	var thongs []models.Thong

	if err := config.DB.Find(&thongs).Error; err != nil {
		return nil, err
	}

	return &thongs, nil
}

func GetThongById(id string) (*models.Thong, error) {
	var thong models.Thong

	config.DB.Preload("ThongRoom").First(&thong, id)

	return &thong, nil
}
