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
	err = db.AutoMigrate(&models.PatientData{}, &models.RelativeHistory{})
	if err != nil {
		panic("Could not migrate")
	}
	return db
}

func UpdatePatientData(toAdd *models.PatientData) models.PatientData {
	db := DBRef()
	err := db.Save(toAdd).Error
	if err != nil {
		log.Printf("%+v", err)
	}
	return *toAdd
}

func AddPatientData(toAdd *models.PatientData) models.PatientData {
	db := DBRef()
	err := db.Create(toAdd).Error
	if err != nil {
		log.Printf("%+v", err)
	}
	return *toAdd
}

func GetPatient(id int) (models.PatientData, error) {
	db := DBRef()
	var patient models.PatientData
	err := db.Preload("RelativeHistory").First(&patient, id).Error
	if err != nil {
		return patient, err
	}

	return patient, err
}

func GetPaginatedPaients(numberOfItems int, offset int) []models.PatientData {
	var patients []models.PatientData
	db := DBRef()
	err := db.Limit(numberOfItems).Offset(offset).Preload("RelativeHistory").Find(&patients).Error
	if err != nil {
		log.Printf("%v", err)
	}
	return patients
}

func RelativeHistoryExists(ID int) bool {
	if ID == 0 {
		return false
	}

	var relativeHistory models.RelativeHistory
	db := DBRef()
	err := db.First(&relativeHistory, ID).Error
	if err == nil {
		return true
	} else {
		return false
	}
}

func UpdateRelativeHistory(toAdd *models.RelativeHistory) models.RelativeHistory {
	db := DBRef()
	err := db.Save(toAdd).Error
	if err != nil {
		log.Printf("%+v", err)
	}
	return *toAdd
}

func AddRelativeHistory(toAdd *models.RelativeHistory) models.RelativeHistory {
	db := DBRef()
	err := db.Create(toAdd).Error
	if err != nil {
		log.Printf("%+v", err)
	}
	return *toAdd
}

func GetRelativeHistory(id int) (models.RelativeHistory, error) {
	db := DBRef()
	var relative models.RelativeHistory
	err := db.First(&relative, id).Error
	if err != nil {
		return relative, err
	}

	return relative, err
}

func DeleteRelativeHistory(id int) {
	db := DBRef()
	relative := models.RelativeHistory{}
	relative.ID = uint(id)

	err := db.Delete(&relative).Error
	if err != nil {
		log.Printf("%+v", err)
	}
}

func DeletePatient(id int) {
	db := DBRef()
	patient := models.PatientData{}
	patient.ID = uint(id)

	err := db.Delete(&patient).Error
	if err != nil {
		log.Printf("%+v", err)
	}

	err = db.Where("patient_data_id = ?", id).Delete(&models.RelativeHistory{}).Error
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
