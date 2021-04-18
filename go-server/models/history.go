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
	ConsentApproval bool
	CancerDX        bool
	CancerDXType    string
	CancerDXAge     int
	KnownBRCA       bool
	KnownCancer     bool
	RelativeHistory []RelativeHistory
}

func NewFullHistory(
	consentApproval bool,
	cancerDX bool,
	cancerDXType string,
	cancerDXAge int,
	knownBRCA bool,
	knownCancer bool,
	relativeHistory []RelativeHistory,
) *FullHistory {
	f := FullHistory{
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
