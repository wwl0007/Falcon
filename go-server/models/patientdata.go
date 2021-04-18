package models

import "gorm.io/gorm"

type PatientData struct {
	gorm.Model
	IsPathogenic bool
	Gene         string
	HistoryClass string
	FullHistory
}

func NewPatientData(
	isPathogenic bool,
	gene string,
	historyClass string,
	fullHistory FullHistory) *PatientData {
	p := PatientData{
		IsPathogenic: isPathogenic,
		Gene:         gene,
		HistoryClass: historyClass,
		FullHistory:  fullHistory,
	}

	return &p
}
