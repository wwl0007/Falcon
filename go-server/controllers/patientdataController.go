package controllers

import (
	"github.com/wwl0007/Project3/database"
	"github.com/wwl0007/Project3/models"
)

func UpdateOrCreateNewPatientData(patientData models.PatientDataREST, calculateHistoryClass bool) models.PatientData {
	toAdd := models.PatientDataRESTToGORM(&patientData)

	if database.PatientDataExists(int(patientData.ID)) {
		return database.UpdatePatientData(toAdd)
	} else {
		if calculateHistoryClass {
			toAdd.HistoryClass = AssignHistoryClass(toAdd)
		}
		return database.AddPatientData(toAdd)
	}
}

func UpdateOrCreateNewRelativeHistory(relativeHistory models.RelativeHistoryREST) models.RelativeHistory {
	toAdd := models.RelativeHistoryRESTToGORM(&relativeHistory)

	if database.RelativeHistoryExists(int(relativeHistory.ID)) {
		return database.UpdateRelativeHistory(toAdd)
	} else {
		return database.AddRelativeHistory(toAdd)
	}
}

func AssignHistoryClass(p *models.PatientData) string {
	class := "none"

	if len(p.FullHistory.RelativeHistory) > 4 {
		class = "strong_family"
	} else if p.FullHistory.KnownCancer == "yes" {
		class = "strong_personal"
	} else if p.FullHistory.CancerDX == "yes" || len(p.FullHistory.RelativeHistory) > 0 {
		class = "not_strong"
	}
	return class
}
