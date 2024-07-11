package doctor

import "database/sql"

type Store struct {
	db *sql.DB
}

type DataAccessInterface interface {
	RegisterUser(user Credentials) error
	RegisterUser(user Credentials) error
}
