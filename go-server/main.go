package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wwl0007/Project3/controllers"
	"github.com/wwl0007/Project3/database"
	"github.com/wwl0007/Project3/models"
)

func getPatientHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: GET patients")
	json.NewEncoder(w).Encode(database.GetAllPatients())
}

func putPatientHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: PUT patients")
	w.Header().Set("Content-Type", "application/json")
	var patientData models.PatientData
	_ = json.NewDecoder(r.Body).Decode(&patientData)
	if database.PatientDataExists(int(patientData.ID)) {
		database.UpdatePatientData(&patientData)
	} else {
		// Sort the patient data
		patientData.HistoryClass = controllers.AssignHistoryClass(&patientData)
		log.Println(patientData.HistoryClass)
		database.AddPatientData(&patientData)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/patients", getPatientHandler).Methods("GET")
	router.HandleFunc("/patients", putPatientHandler).Methods("PUT")

	log.Println("Server started, listening at port :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
