package main

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>go server running!</h1>")
	})

	router.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Healthy!"))
	})

	http.ListenAndServe(":8000", handler)

}
