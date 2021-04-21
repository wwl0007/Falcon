package main

import (
	"github.com/wwl0007/Project3/database"
	"github.com/wwl0007/Project3/routes"
)

func main() {
	database.DBRef()
	routes.ServeREST()
}
