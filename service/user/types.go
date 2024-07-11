package user

import (
	"database/sql"
	"time"
)

// model
type User struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Phone     sql.NullString `json:"phone"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"createdAt"`
}

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
