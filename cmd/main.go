package main

import (
	"github.com/kinxyo/Servebix.git/cmd/api"
	"github.com/kinxyo/Servebix.git/database"
	"log"
)

func main() {

	db, _ := database.GetDB()

	database.DbInit(db)

	server := api.NewAPI(":8000", db)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
