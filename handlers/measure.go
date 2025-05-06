package handlers

import (
	"net/http"
	"test1/database"
	"test1/models"

	"github.com/gin-gonic/gin"
)

func GetMeasure(c *gin.Context) {
	var measure models.Measure
	if err := database.DB.Find(&measure).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, measure)
}

func GetMeasureID(c *gin.Context) {
	id := c.Param("id")
	var measure models.Measure

	if err := database.DB.First(&measure, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "еденица не найдена"})
		return
	}
	c.JSON(http.StatusOK, measure)
}

func CreateMeasure(c *gin.Context) {
	var measure models.Measure

	if err := c.ShouldBindJSON(&measure); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не верный json"})
		return
	}

	if err := database.DB.Create(&measure).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать measure"})
		return
	}
	c.JSON(http.StatusOK, measure)
}

func UpdateMeasure(c *gin.Context) {
	id := c.Param("id")
	var measure models.Measure

	if err := database.DB.First(&measure, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "не удалось найти"})
		return
	}

	if err := c.ShouldBindJSON(&measure); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не верный json"})
		return
	}

	if err := database.DB.Save(&measure).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить"})
		return
	}
	c.JSON(http.StatusOK, measure)
}

func DeleteMeasure(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Measure{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось удалить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "удалено"})
}
