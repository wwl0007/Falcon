package database

import (
	"github.com/wwl0007/Project3/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(dbFilename string) *gorm.DB {
	// Hello world, the web server
	db, err := gorm.Open(sqlite.Open(dbFilename), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.PatientData{}, &models.RelativeHistory{})

	return db
}
