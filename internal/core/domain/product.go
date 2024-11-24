package domain

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

var (
	err_product_empty_fields = errors.New("name can not be empty")
)

// -----------------------------
// Login is a core domain entity
// ------------------------------
type Product struct {
	ID                 uuid.UUID
	Name               string
	Description        string
	AnnualContribution int
}

// --------------------------------
// function : NewLogin
// - Business validation for creating a new client
// --------------------------------

func NewProduct(name, descript string, ac int) (*Product, error) {
	// validation of the information
	if name == "" || ac == 0 {
		log.Println("error: ", err_product_empty_fields)
		return nil, err_product_empty_fields
	}

	return &Product{
		ID:                 uuid.New(),
		Name:               name,
		Description:        descript,
		AnnualContribution: ac,
	}, nil

}
