package handlers

import (
	"net/http"
	"thong-khon-api/config"
	"thong-khon-api/models"
	"thong-khon-api/services"

	"github.com/gin-gonic/gin"
)

func CreateThong(c *gin.Context) {

	var body struct {
		Name       string
		IsMale     bool
		IsFemale   bool
		IsDisabled bool
		Lat        string
		Long       string
	}

	c.Bind(&body)

	if body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name missing"})
		return
	}

	if !body.IsMale {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IsMale missing"})
		return
	}

	if !body.IsFemale {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IsFeMale missing"})
		return
	}

	if !body.IsDisabled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IsDisabled missing"})
		return
	}

	if body.Lat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lat missing"})
		return
	}

	if body.Long == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "long missing"})
		return
	}

	thong := models.Thong{
		IsOpen:     true,
		Name:       body.Name,
		Ratting:    0,
		IsMale:     body.IsMale,
		IsFemale:   body.IsFemale,
		IsDisabled: body.IsDisabled,
		Lat:        body.Lat,
		Long:       body.Long,
	}

	createdThong, err := services.CreateThong(thong)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": createdThong,
	})
}

func CreateThongRoom(c *gin.Context) {

	var body struct {
		ThongId  string
		RoomType string
	}

	c.Bind(&body)

	if body.ThongId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thongId missing"})
		return

	}
	if body.RoomType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RoomType missing"})
		return

	}
	if body.RoomType != string(models.Male) && body.RoomType != string(models.Female) && body.RoomType != string(models.Disabled) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RoomType  invalid"})
		return
	}

	var thong models.Thong

	if err := config.DB.First(&thong, body.ThongId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Thong not found"})
		return
	}

	thongRoom := models.ThongRoom{
		ThongID:     thong.ID,
		RoomType:    models.RomeType(body.RoomType),
		IsAvailable: true,
	}

	createdThongRoom, err := services.CreateThongRoom(thongRoom)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": createdThongRoom,
	})

}

func GetAllThong(c *gin.Context) {

	thongList, err := services.GetAllThong()

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": thongList,
	})
}

func GetThongById(c *gin.Context) {
	id := c.Param("id")

	thongIncludedThongRoom, err := services.GetThongById(id)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": thongIncludedThongRoom,
	})

}
