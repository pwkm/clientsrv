package repository

import (
	"client/internal/core/domain"
	"database/sql"
	"log"
)

// ------------------------------------
// Client Repository for CRUD ACTIONS
// ------------------------------------
type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db,
	}
}

// ------------------------------------
// Function: Save Client
// ------------------------------------
func (cr *ProductRepository) SaveProduct(product *domain.Product) error {
	// Save product
	insertStmt := `INSERT INTO product(id, name, description, annualcontribution) VALUES($1, $2, $3, $4)`
	_, err := cr.db.Exec(insertStmt, product.ID, product.Name, product.Description, product.AnnualContribution)
	if err != nil {
		log.Printf("error saving product: %v", err)
		return err
	}

	return nil
}
