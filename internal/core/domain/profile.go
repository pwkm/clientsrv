package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/pwkm/clientsrv/internal/utils"
)

var (
	err_profile_invalid_email = errors.New("the email address is not valid")
)

// Client is a core domain entity
type Profile struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Street    string
	Number    string
	Postcode  int
	Community string
	Email     string
	Birthday  time.Time
	ClientID  uuid.UUID //foreign key
}

// --------------------------------
// function : NewProfile
// - Business validation for creating a new profile
// --------------------------------

func NewProfile(
	street, number string,
	postcode int,
	community, email string,
	birthday time.Time,
) (*Profile, error) {

	// Validation of the profile information
	if email == "" {
		return nil, err_profile_invalid_email
	}
	if !utils.ValidEmail(email) {
		return nil, err_profile_invalid_email
	}

	// Create and return a login structure
	return &Profile{
		Id:        uuid.New(),
		Street:    street,
		Number:    number,
		Postcode:  postcode,
		Community: community,
		Email:     email,
		Birthday:  birthday,
	}, nil

}
