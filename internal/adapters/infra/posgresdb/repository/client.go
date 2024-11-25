package repository

import (
	"client/internal/core/domain"
	"errors"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ERROR_NO_CLIENT_FOUND = errors.New("no rows were found")
)

// ------------------------------------
// Client Repository for CRUD ACTIONS
// ------------------------------------
type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		db,
	}
}

// ------------------------------------
// Function: Save Client
// ------------------------------------
func (cr *ClientRepository) SaveClient(client *domain.Client) error {
	// begin a transaction
	tx := cr.db.Begin()
	defer func() {
		tx.Rollback()
	}()

	// Save client
	result := tx.Create(&client)
	if result.Error != nil {
		log.Printf("error saving login: %v", result.Error)
		return result.Error
	}

	// commit transactions
	result = tx.Commit()
	return result.Error
}

// ------------------------------------
// Func: Query Clients
// ------------------------------------
func (cr *ClientRepository) QueryClients() ([]*domain.Client, error) {
	var clients []*domain.Client

	// Query all clients
	result := cr.db.Preload("Profile").Preload("Login").Find(&clients)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ERROR_NO_CLIENT_FOUND
	}

	return clients, nil
}

// ------------------------------------
// Func: Retrieve Client By ID
// ------------------------------------
func (cr *ClientRepository) QueryClientByID(id uuid.UUID) (*domain.Client, error) {
	var client = new(domain.Client)

	// Query the client
	result := cr.db.Preload("Login").Preload("Profile").First(&client, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ERROR_NO_CLIENT_FOUND
	}
	return client, nil
}

// ------------------------------------
// Func: Delete Client
// ------------------------------------
func (cr *ClientRepository) DeleteClient(id uuid.UUID) error {

	// Delete client
	result := cr.db.Delete(&Client{}, id)
	if result.Error != nil {
		log.Printf("error saving client: %v", err)
		return err
	}
	if result.RowsAffected == 0 {
		return ERROR_NO_CLIENT_FOUND
	}

	return nil
}
