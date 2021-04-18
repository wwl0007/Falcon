package models

import "gorm.io/gorm"

type RelativeHistory struct {
	gorm.Model
	Relation      string
	Cancer        string
	Age           int
	PatientDataID int
}

type FullHistory struct {
	Ethnicity       string `json:"ethnicity"`
	ConsentApproval bool   `json:"consent_approval"`
	CancerDX        bool   `json:"cancer_dx"`
	CancerDXType    string `json:"cancer_dx_type"`
	CancerDXAge     int    `json:"cancer_dx_age"`
	KnownBRCA       bool   `json:"known_brca"`
	KnownCancer     bool   `json:"known_cancer"`
	Method          string `json:"_method"`
	RelativeHistory []RelativeHistory
}

func NewFullHistory(
	ethnicity string,
	consentApproval bool,
	cancerDX bool,
	cancerDXType string,
	cancerDXAge int,
	knownBRCA bool,
	knownCancer bool,
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
) *RelativeHistory {
	r := RelativeHistory{
		Relation: relation,
		Cancer:   cancer,
		Age:      age,
	}

	return &r
}
