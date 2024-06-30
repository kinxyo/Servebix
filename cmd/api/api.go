package api

import (
	"database/sql"
	"github.com/kinxyo/Servebix.git/service/home"
	"github.com/kinxyo/Servebix.git/service/user"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type API struct {
	addr string
	db   *sql.DB
}

func NewAPI(addr string, db *sql.DB) *API {
	return &API{addr: addr, db: db}
}

func (api *API) Start() error {

	router := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	/* Add Services here ⭐ */

	// HOME
	homeHandler := home.NewHandler()
	homeHandler.RegisterRoutes(router)

	// USER
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(router)

	/* ----------------- */

	server := &http.Server{Addr: api.addr, Handler: handler}

	log.Printf("listening on %s", api.addr)

	return server.ListenAndServe()
}
