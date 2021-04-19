package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wwl0007/Project3/controllers"
	"github.com/wwl0007/Project3/database"
	"github.com/wwl0007/Project3/models"
)

func getAllPatientsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.GetAllPatients())
}

func getPatientHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Could not find the patient")
		return
	}
	json.NewEncoder(w).Encode(database.GetPatient(id))
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
	router.HandleFunc("/patients", getAllPatientsHandler).Methods("GET")
	router.HandleFunc("/patients/{id}", getPatientHandler).Methods("GET")
	router.HandleFunc("/patients", putPatientHandler).Methods("PUT")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Println("Server started, listening at port :8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
