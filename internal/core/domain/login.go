package domain

import (
	"client/internal/utils"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	err_login_empty_fields     = errors.New("email and login can not be empty")
	err_login_invalid_email    = errors.New("email is not valid email address")
	err_login_invalid_password = errors.New("passsword is not conform the policy")
	err_login_hash_password    = errors.New("can not hash the password")
)

// -----------------------------
// Login is a core domain entity
// ------------------------------
type Login struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Email    string
	Password string
	ClientID uuid.UUID //Foreign key to user
}

// --------------------------------
// function : NewLogin
// - Business validation for creating a new client
// --------------------------------

func NewLogin(
	email, password string, clientid uuid.UUID,
) (*Login, error) {

	// Validation of the login information
	if email == "" || password == "" {
		return nil, err_login_empty_fields
	}
	if !utils.ValidEmail(email) {
		return nil, err_login_invalid_email
	}
	if !utils.ValidPassword(password) {
		return nil, err_login_invalid_password
	}

	// Hash Password - generate a bcrypt hash for a given password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err_login_hash_password
	}
	hash := string(bytes)

	// Create and return a login structure
	return &Login{
		ID:       uuid.New(),
		Email:    email,
		Password: hash,
		ClientID: clientid,
	}, nil

}
