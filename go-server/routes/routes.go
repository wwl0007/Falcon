package routes

import (
	"encoding/json"
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

func getPatientListHandler(w http.ResponseWriter, r *http.Request) {
	items, err := strconv.Atoi(r.FormValue("items"))
	if err != nil {
		w.Write([]byte("items could not be converted to an int\n"))
	}

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		w.Write([]byte("items could not be converted to an int\n"))
	}

	json.NewEncoder(w).Encode(database.GetPaginatedPaients(items, offset))
}

func getPatientHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Could not find the patient")
		return
	}

	patient, err := database.GetPatient(id)
	if err == nil {
		json.NewEncoder(w).Encode(patient)
	} else {
		json.NewEncoder(w).Encode("Could not find the relative history")
	}
}

func getRelativeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Could not find the relative history")
		return
	}

	history, err := database.GetRelativeHistory(id)
	if err == nil {
		json.NewEncoder(w).Encode(history)
	} else {
		json.NewEncoder(w).Encode("Could not find the relative history")
	}
}

func putPatientHandler(w http.ResponseWriter, r *http.Request) {
	var patientData models.PatientDataREST
	_ = json.NewDecoder(r.Body).Decode(&patientData)

	if !(len(patientData.FullHistory.RelativeAge) == len(patientData.FullHistory.RelativeRelation) &&
		len(patientData.FullHistory.RelativeAge) == len(patientData.FullHistory.RelativeCancer)) {
		json.NewEncoder(w).Encode("Not all relative's data is filled out")
		return
	}

	controllers.UpdateOrCreateNewPatientData(patientData, true)
}

func deletePatientHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Could not find the patient")
		return
	}
	database.DeletePatient(id)
}

func deleteRelativeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Could not find the relative history")
		return
	}
	database.DeleteRelativeHistory(id)
}

func putRelativeHistory(w http.ResponseWriter, r *http.Request) {
	var relativeHistory models.RelativeHistoryREST
	_ = json.NewDecoder(r.Body).Decode(&relativeHistory)

	controllers.UpdateOrCreateNewRelativeHistory(relativeHistory)
}

func ServeREST() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/relativeHistory", putRelativeHistory).Methods("PUT")
	router.HandleFunc("/relativeHistory/{id}", getRelativeHistoryHandler).Methods("GET")
	router.HandleFunc("/relativeHistory/{id}", deleteRelativeHistoryHandler).Methods("DELETE")
	router.HandleFunc("/patients", getAllPatientsHandler).Methods("GET")
	router.HandleFunc("/patientList", getPatientListHandler).Methods("GET").Queries("items", "{items:[0-9]+}", "offset", "{offset:[0-9]+}")
	router.HandleFunc("/patients/{id}", getPatientHandler).Methods("GET")
	router.HandleFunc("/patients", putPatientHandler).Methods("PUT")
	router.HandleFunc("/patients/{id}", deletePatientHandler).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Println("Server started, listening at port :8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
