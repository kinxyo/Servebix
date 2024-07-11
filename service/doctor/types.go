package doctor

import (
	"database/sql"
	"time"
)

type Credentials struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Session struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// model
type Auth struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Phone    sql.NullString `json:"phone"`
	Password string         `json:"-"`
}

// model
type Doctor struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Avatar         string    `json:"avatar"`
	Specialization string    `json:"specialization"`
	Active         bool      `json:"active"`
	Rating         int       `json:"rating"`
	Language       string    `json:"language"`
	Nationality    string    `json:"nationality"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
