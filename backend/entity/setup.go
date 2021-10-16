package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	database.AutoMigrate(
		&Pharmacist{}, &Medicine{}, &Admission{}, &TreatmentRecord{}, &MedicationRecord{},
	)

	db = database

	//Pharmacist
	db.Model(&Pharmacist{}).Create(&Pharmacist{
		Name: "Chawarat  Narit",
		Pid:  "1400011111111",
	})
	db.Model(&Pharmacist{}).Create(&Pharmacist{
		Name: "Pichanon  Srisongmuang",
		Pid:  "1400011111112",
	})
	db.Model(&Pharmacist{}).Create(&Pharmacist{
		Name: "Chin Love",
		Pid:  "1400011111113",
	})
	//Medicine
	db.Model(&Medicine{}).Create(&Medicine{
		Name:  "PARACETAMOL 500 MG",
		Type:  "TAB",
		Price: 1,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Name:  "CEFCINIR 100 MG",
		Type:  "CAP",
		Price: 14,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Name:  "PREDNISLONE 5 MG",
		Type:  "TAB",
		Price: 1,
	})
	//Admission
	db.Model(&Admission{}).Create(&Admission{
		PatientID:       "I562",
		Patient_Name:    "huy",
		RoomID:          "RM2002",
		AdmitTime:       time.Now(),
		Right_Treatment: "GM0001",
	})
	db.Model(&Admission{}).Create(&Admission{
		PatientID:       "I563",
		Patient_Name:    "guli",
		RoomID:          "RM2005",
		AdmitTime:       time.Time{},
		Right_Treatment: "IV0003",
	})
	db.Model(&Admission{}).Create(&Admission{
		PatientID:       "I564",
		Patient_Name:    "ball",
		RoomID:          "RM2004",
		AdmitTime:       time.Time{},
		Right_Treatment: "IV0002",
	})
	var huy Admission
	var guli Admission
	db.Raw("SELECT * FROM admissions WHERE patient_id=?", "I562").Scan(&huy)
	db.Raw("SELECT * FROM admissions WHERE patient_id=?", "I562").Scan(&guli)

	//TreatmentRecord
	db.Model(&TreatmentRecord{}).Create(&TreatmentRecord{
		Model:          gorm.Model{},
		Length_of_stay: 3,
		Treatment:      "Heart Transplant",
		Food_type:      3000,
		Med_amount:     3,
		Cost:           50000,
		Equipment_id:   002,
		Med:            Medicine{},
		Admission:      huy,
	})
	db.Model(&TreatmentRecord{}).Create(&TreatmentRecord{
		Length_of_stay: 3,
		Treatment:      "Gastric lavage",
		Food_type:      3001,
		Med_amount:     3,
		Cost:           50000,
		Equipment_id:   001,
		Med:            Medicine{},
		Admission:      guli,
	})
	//MedicationRecord
	db.Model(&MedicationRecord{}).Create(&MedicationRecord{
		Amount:         0,
		RecordTime:     time.Time{},
		PharmaID:       new(uint),
		Pharma:         Pharmacist{},
		MedID:          new(uint),
		Med:            Medicine{},
		TreatmentID:    new(uint),
		AdmissionID:    new(uint),
		Admission_Name: "",
	})
}
