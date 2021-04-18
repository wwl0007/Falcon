package database

import (
	"log"

	"github.com/wwl0007/Project3/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(dbFilename string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFilename), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.PatientData{})

	return db
}

func AddPatientData(db *gorm.DB, toAdd *models.PatientData) {
	err := db.Create(toAdd).Error
	if err != nil {
		log.Printf("%+v", err)
	}
}
