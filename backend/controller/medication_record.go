package controller

import (
	"net/http"

	"github.com/S1CKK/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /medication_records
func CreateMedicationRecord(c *gin.Context) {

	var medicationrecord entity.MedicationRecord
	//var pharmacist entity.Pharmacist
	//var admission entity.Admission
	var medicine entity.Medicine
	var treatmentrecord entity.TreatmentRecord

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร medicationrecord
	if err := c.ShouldBindJSON(&medicationrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา treatmentrecord ด้วย id
	if tx := entity.DB().Where("id = ?", medicationrecord.TreatmentID).First(&treatmentrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatment record not found"})
		return
	}

	// 10: ค้นหา medicine ด้วย id
	if tx := entity.DB().Where("id = ?", medicationrecord.MedID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}

	// 11: สร้าง MedicationRecord
	wv := entity.MedicationRecord{
		Med:       medicine,        // โยงความสัมพันธ์กับ Entity Resolution
		Treatment: treatmentrecord, // โยงความสัมพันธ์กับ Entity Video

		RecordTime: medicationrecord.RecordTime, // ตั้งค่าฟิลด์ watchedTime
	}

	// 12: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /medication_record/:id
func GetMedicationRacord(c *gin.Context) {
	var medicationrecord entity.MedicationRecord
	id := c.Param("id")
	if err := entity.DB().Preload("TreatmentRecord").Preload("Medicine").Raw("SELECT * FROM medication_records WHERE id = ?", id).Find(&medicationrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicationrecord})
}

// GET /medication_records
func ListMedicationRacord(c *gin.Context) {
	var medicationrecord []entity.MedicationRecord
	if err := entity.DB().Preload("TreatmentRecord").Preload("Medicine").Raw("SELECT * FROM watch_videos").Find(&medicationrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicationrecord})
}

// DELETE /medication_records/:id
func DeleteMedicationRacord(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medication_record WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medication record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /medication_records
func UpdateMedicationRacord(c *gin.Context) {
	var medicationrecord entity.MedicationRecord
	if err := c.ShouldBindJSON(&medicationrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", medicationrecord.ID).First(&medicationrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medication record not found"})
		return
	}

	if err := entity.DB().Save(&medicationrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicationrecord})
}
