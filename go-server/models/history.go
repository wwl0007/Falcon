package models

import (
	"gorm.io/gorm"
)

type RelativeHistory struct {
	gorm.Model
	Relation      string
	Cancer        string
	Age           int
	PatientDataID int
}

type RelativeHistoryREST struct {
	ID            int
	Relation      string
	Cancer        string
	Age           int
	PatientDataID int
}

type FullHistory struct {
	Ethnicity       string
	ConsentApproval string
	CancerDX        string
	CancerDXType    string
	CancerDXAge     int
	KnownBRCA       string
	KnownCancer     string
	RelativeHistory []RelativeHistory
}

type FullHistoryREST struct {
	ConsentApproval  string   `json:"consent_approval"`
	Ethnicity        string   `json:"ethnicity"`
	CancerDX         string   `json:"cancer_dx"`
	CancerDXType     string   `json:"cancer_dx_type"`
	CancerDXAge      string   `json:"cancer_dx_age"`
	KnownBRCA        string   `json:"known_brca"`
	KnownCancer      string   `json:"known_cancer"`
	RelativeRelation []string `json:"rel_relation"`
	RelativeCancer   []string `json:"rel_cancer"`
	RelativeAge      []string `json:"rel_age"`
}

func NewFullHistory(
	ethnicity string,
	consentApproval string,
	cancerDX string,
	cancerDXType string,
	cancerDXAge int,
	knownBRCA string,
	knownCancer string,
	relativeHistory []RelativeHistory,
) *FullHistory {
	f := FullHistory{
		Ethnicity:       ethnicity,
		ConsentApproval: consentApproval,
		CancerDX:        cancerDX,
		CancerDXType:    cancerDXType,
		CancerDXAge:     cancerDXAge,
		KnownBRCA:       knownBRCA,
		KnownCancer:     knownCancer,
		RelativeHistory: relativeHistory,
	}

	return &f
}

func NewRelativeHistory(
	relation string,
	cancer string,
	age int,
	patientID int,
) *RelativeHistory {
	r := RelativeHistory{
		Relation:      relation,
		Cancer:        cancer,
		Age:           age,
		PatientDataID: patientID,
	}

	return &r
}

func RelativeHistoryRESTToGORM(rest *RelativeHistoryREST) *RelativeHistory {
	ret := NewRelativeHistory(rest.Relation, rest.Cancer, rest.Age, rest.PatientDataID)
	ret.ID = uint(rest.ID)
	return ret
}
