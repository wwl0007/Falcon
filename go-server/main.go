package main

import (
	"io"
	"log"
	"net/http"

	"github.com/wwl0007/Project3/database"
)

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	dbRef := database.New("domo.db")
	defer func() {
		sqlDB, _ := dbRef.DB()
		err := sqlDB.Close()
		if err != nil {
			log.Fatalf("Could not close database: %v", err)
		}
	}()

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
