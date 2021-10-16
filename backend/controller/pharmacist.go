package controller

import (
	"net/http"

	"github.com/S1CKK/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// GET /users
// List all users
func ListUsers(c *gin.Context) {
	var pharmacist []entity.Pharmacist
	if err := entity.DB().Raw("SELECT * FROM pharmacists").Scan(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}

// GET /user/:id
// Get user by id
func GetUser(c *gin.Context) {
	var pharmacist entity.Pharmacist
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM pharmacists WHERE id = ?", id).Scan(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}

// POST /pharmacists
func CreateUser(c *gin.Context) {
	var pharmacist entity.Pharmacist
	if err := c.ShouldBindJSON(&pharmacist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}

// PATCH /pharmacists
func UpdateUser(c *gin.Context) {
	var pharmacist entity.Pharmacist
	if err := c.ShouldBindJSON(&pharmacist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", pharmacist.ID).First(&pharmacist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}

// DELETE /pharmacists/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM pharmacists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
