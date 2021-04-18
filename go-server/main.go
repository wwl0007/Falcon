package main

import (
	"io"
	"log"
	"net/http"

	"github.com/wwl0007/Project3/database"
	"github.com/wwl0007/Project3/models"
)

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	dbRef := database.New("domo.db")

	fh := models.NewFullHistory(false, false, "", 0, true, true, nil)
	p := models.NewPatientData(false, "A", "adssa", *fh)
	u := &models.PatientData{
		IsPathogenic: true,
		Gene:         "asad",
		HistoryClass: "Adsad",
		FullHistory:  *fh,
	}

	dbRef.Create(u)
	dbRef.Create(p)

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
