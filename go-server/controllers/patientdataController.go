package controllers

import (
	"math/rand"

	"github.com/wwl0007/Project3/models"
)

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
