package handlers

import (
	"net/http"
	"sneakers/database"
	"sneakers/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetBrands(c *gin.Context) {
	var brands []models.Brand

	// Пагинация
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	query := database.DB.Offset((page - 1) * limit).Limit(limit)
	if err := query.Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brands)
}


func GetBrandByID(c *gin.Context) {
	var brand models.Brand
	if err := database.DB.First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func CreateBrand(c *gin.Context) {
	var brand models.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&brand)
	c.JSON(http.StatusCreated, brand)
}

func UpdateBrand(c *gin.Context) {
	var brand models.Brand
	if err := database.DB.First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&brand)
	c.JSON(http.StatusOK, brand)
}

func DeleteBrand(c *gin.Context) {
	if err := database.DB.Delete(&models.Brand{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not delete brand"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted"})
}
