package user

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type StoreInterface interface {
	RegisterUser(user Credentials) error
	ValidateUser(user Credentials) (Session, error)
}

func (store *Store) ValidateUser(user Credentials) (Session, error) {

	var session Session

	// Fetch storedPassword
	storedPassword, err := store.FetchStoredPassword(user.Email)
	if err != nil {
		return session, err
	}

	// UnHash Password
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return session, fmt.Errorf("invalid password")
		}
		return session, err
	}

	// Return Session
	err = store.FetchSession(&session, user.Email)
	if err != nil {
		return session, err
	}

	return session, nil
}

func (store *Store) RegisterUser(user Credentials) error {

	// Check if email already exists
	exists, err := store.EmailExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Add user to database
	query := "INSERT INTO patient_creds ( name, email, password, phone ) VALUES ( ?, ?, ?, ? );"
	_, err = store.db.Exec(query, user.Name, user.Email, hashedPassword, user.Phone)
	if err != nil {
		return err
	}

	return nil
}

/*--------Helper Functions--------*/

func (store *Store) EmailExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM patient_creds WHERE email = ?)"
	err := store.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if email already exists: %v", err)
	}
	return exists, nil
}

func (store *Store) FetchStoredPassword(email string) (string, error) {
	var storedPassword string

	query := "SELECT password FROM patient_creds WHERE email = ?"

	err := store.db.QueryRow(query, email).Scan(&storedPassword)
	if err != nil {
		return "", fmt.Errorf("email not found: %v", err)
	}
	return storedPassword, nil
}

func (store *Store) FetchSession(session *Session, email string) error {

	query := "SELECT id, name FROM patient_creds WHERE email = ?"

	err := store.db.QueryRow(query, email).Scan(&session.ID, &session.Name)
	if err != nil {
		return fmt.Errorf("failed to fetch session: %v", err)
	}
	return nil
}

// STORE DEFINED

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}
