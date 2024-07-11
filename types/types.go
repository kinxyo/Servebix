package types

import (
	"time"
)

type DoctorCreds struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

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
