package controller

import (
	"net/http"

	"github.com/S1CKK/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /treatmentrecord
func CreateTreatmentRecord(c *gin.Context) {
	var treatment_record entity.TreatmentRecord
	if err := c.ShouldBindJSON(&treatment_record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&treatment_record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": treatment_record})
}

// GET /treatmentrecord/:id
func GetTreatmentRecord(c *gin.Context) {
	var treatment_record entity.TreatmentRecord
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM treatment_records WHERE id = ?", id).Scan(&treatment_record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatment_record})
}

// GET /treatmentrecords
func ListTreatmentRecord(c *gin.Context) {
	var treatment_record []entity.TreatmentRecord
	if err := entity.DB().Raw("SELECT * FROM treatment_records").Scan(&treatment_record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatment_record})
}

// DELETE /treatmentrecords/:id
func DeleteTreatmentRecord(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM treatment_records WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatment record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /treatmentrecords
func UpdateTreatmentRecord(c *gin.Context) {
	var treatment_record entity.TreatmentRecord
	if err := c.ShouldBindJSON(&treatment_record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", treatment_record.ID).First(&treatment_record); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatment record not found"})
		return
	}

	if err := entity.DB().Save(&treatment_record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatment_record})
}
