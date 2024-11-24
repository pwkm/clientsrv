package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	err_client_no_id      = errors.New("the client has no unique ID")
	err_client_no_name    = errors.New("the client has no name")
	err_client_no_profile = errors.New("for this client is no profile defined")
	err_client_no_login   = errors.New("for this client is no login defined")
)

// ------------------------------
// Aggregate Client
// ------------------------------

type Client struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	Name    string
	Login   *Login
	Profile *Profile
}

// --------------------------------
// function : NewLogin
// - Business validation for creating a new client
// --------------------------------

func NewClient(id uuid.UUID, name string, login *Login, profile *Profile) (*Client, error) {

	// Validation of the client info
	if id == uuid.Nil {
		return nil, err_client_no_id
	}
	if name == "" {
		return nil, err_client_no_name
	}
	if profile == nil {
		return nil, err_client_no_profile
	}
	if login == nil {
		return nil, err_client_no_login
	}

	// Create and return a client structure
	return &Client{
		ID:      id,
		Name:    name,
		Login:   login,
		Profile: profile,
	}, nil
}
