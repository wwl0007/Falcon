package controllers

import (
	"math/rand"

	"github.com/wwl0007/Project3/database"
	"github.com/wwl0007/Project3/models"
)

func UpdateOrCreateNewPatientData(patientData models.PatientDataREST, calculateHistoryClass bool) {
	toAdd := models.PatientDataRESTToGORM(&patientData)

	if database.PatientDataExists(int(patientData.ID)) {
		database.UpdatePatientData(toAdd)
	} else {
		if calculateHistoryClass {
			toAdd.HistoryClass = AssignHistoryClass(toAdd)
		}
		database.AddPatientData(toAdd)
	}
}

func UpdateOrCreateNewRelativeHistory(relativeHistory models.RelativeHistoryREST) {
	toAdd := models.RelativeHistoryRESTToGORM(&relativeHistory)

	if database.RelativeHistoryExists(int(relativeHistory.ID)) {
		database.UpdateRelativeHistory(toAdd)
	} else {
		database.AddRelativeHistory(toAdd)
	}
}

func AssignHistoryClass(p *models.PatientData) string {
	val := rand.Intn(400) % 4
	switch val {
	case 0:
		return "strong_personal"
	case 1:
		return "strong_family"
	case 2:
		return "not_strong"
	case 3:
		return "none"
	}
	return "none"
}
