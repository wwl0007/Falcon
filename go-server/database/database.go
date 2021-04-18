package database

import (
	"log"

	"github.com/wwl0007/Project3/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbRef *gorm.DB

func DBRef() *gorm.DB {
	if dbRef == nil {
		dbRef = new("domo.db")
	}
	return dbRef
}

func new(dbFilename string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFilename), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.PatientData{}, &models.RelativeHistory{})

	return db
}

func UpdatePatientData(toAdd *models.PatientData) {
	db := DBRef()
	err := db.Save(toAdd).Error
	if err != nil {
		log.Printf("%+v", err)
	}
}

func AddPatientData(toAdd *models.PatientData) {
	db := DBRef()
	err := db.Create(toAdd).Error
	if err != nil {
		log.Printf("%+v", err)
	}
}

func GetAllPatients() []models.PatientData {
	var patients []models.PatientData
	db := DBRef()
	err := db.Preload("RelativeHistory").Find(&patients).Error
	if err != nil {
		log.Printf("%v", err)
	}
	return patients
}

func PatientDataExists(ID int) bool {
	if ID == 0 {
		return false
	}

	var patientData models.PatientData
	db := DBRef()
	err := db.First(&patientData, ID).Error
	if err == nil {
		return true
	} else {
		return false
	}
}
