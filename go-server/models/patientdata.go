package models

import (
	"strconv"

	"gorm.io/gorm"
)

type PatientData struct {
	gorm.Model
	Pathogenic   string
	Gene         string
	HistoryClass string
	FullHistory
}

type PatientDataREST struct {
	ID           int
	Pathogenic   string
	Gene         string
	HistoryClass string          `json:"history_class"`
	FullHistory  FullHistoryREST `json:"full_history"`
}

func NewPatientData(
	isPathogenic string,
	gene string,
	historyClass string,
	fullHistory FullHistory) *PatientData {
	p := PatientData{
		Pathogenic:   isPathogenic,
		Gene:         gene,
		HistoryClass: historyClass,
		FullHistory:  fullHistory,
	}

	return &p
}

func PatientDataRESTToGORM(rest *PatientDataREST) *PatientData {
	cancerDXAge, _ := strconv.Atoi(rest.FullHistory.CancerDXAge)
	ret := NewPatientData(
		rest.Pathogenic,
		rest.Gene,
		rest.HistoryClass,
		*NewFullHistory(
			rest.FullHistory.Ethnicity,
			rest.FullHistory.ConsentApproval,
			rest.FullHistory.CancerDX,
			rest.FullHistory.CancerDXType,
			cancerDXAge,
			rest.FullHistory.KnownBRCA,
			rest.FullHistory.KnownCancer,
			nil,
		),
	)
	for i := range rest.FullHistory.RelativeAge {
		age, _ := strconv.Atoi(rest.FullHistory.RelativeAge[i])
		ret.RelativeHistory = append(
			ret.RelativeHistory,
			*NewRelativeHistory(
				rest.FullHistory.RelativeRelation[i],
				rest.FullHistory.RelativeCancer[i],
				age,
			),
		)
	}

	return ret
}
