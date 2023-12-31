package controllers

import (
	"models"
	"net/http"
	"strconv"

	"PBI/database"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	var newPhoto models.Photo

	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat foto"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"photo": newPhoto})
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo

	if err := database.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil foto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

func UpdatePhotoByID(c *gin.Context) {
	photoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID foto tidak valid"})
		return
	}

	var updatedPhoto models.Photo

	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingPhoto models.Photo
	if err := database.DB.First(&existingPhoto, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Foto tidak ditemukan"})
		return
	}

	existingPhoto.Title = updatedPhoto.Title
	existingPhoto.Caption = updatedPhoto.Caption
	existingPhoto.PhotoURL = updatedPhoto.PhotoURL

	if err := database.DB.Save(&existingPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui foto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photo": existingPhoto})
}

func DeletePhotoByID(c *gin.Context) {
	photoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID foto tidak valid"})
		return
	}

	if err := database.DB.Delete(&models.Photo{}, photoID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus foto"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
