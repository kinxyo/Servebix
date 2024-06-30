package main

import (
	"github.com/kinxyo/Servebix.git/cmd/api"
	"log"
)

func main() {

	server := api.NewAPI(":8000", nil)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
