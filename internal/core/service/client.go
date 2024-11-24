package service

import (
	"client/internal/core/domain"
	"client/internal/core/ports"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

// ------------------------------
// Global ERROR Messages
// ------------------------------
var (
	err_client_creation_error = errors.New("client could not be created ")
	err_client_no_records     = errors.New("No records found")
)

// -----------------------------
// Client Service Structure
// -----------------------------
type ClientService struct {
	Repo ports.IClientRepository
	// Stream ports.IMessageStream
}

func NewClientService(repo ports.IClientRepository) *ClientService {
	return &ClientService{
		Repo: repo,
		// Stream: str,
	}
}

// -----------------------------
// Function: Register a new client
// -----------------------------
func (c *ClientService) RegisterClient(
	name, email, password string,
) (uuid.UUID, error) {
	// initiatie variabelen
	var id uuid.UUID
	var login *domain.Login
	var profile *domain.Profile

	// Create an unique ID
	id, err := uuid.NewV6()

	// Create new login entry
	login, err = domain.NewLogin(email, password, id)
	if err != nil {
		log.Printf("error creating login: %v", err)
		return uuid.Nil, err_client_creation_error
	}

	// Create a first profile
	profile, err = domain.NewProfile("", "", 0, "", email, time.Now())
	if err != nil {
		log.Printf("error creating profile: %v", err)
		return uuid.Nil, err_client_creation_error
	}

	// Create a client
	client, err := domain.NewClient(id, name, login, profile)
	if err != nil {
		log.Printf("error creating client: %v", err)
		return uuid.Nil, err_client_creation_error
	}

	// save client in the repository
	err = c.Repo.SaveClient(client)
	if err != nil {
		log.Printf("Client can not be saved: %v", err)
		return uuid.Nil, err_client_creation_error
	}

	// Stream client as a message
	// err = c.Stream.
	// 	SendMessage(client)
	// if err != nil {
	// 	log.Printf("client can not be streamed: %v", err)
	// 	return uuid.Nil, err_client_creation_error
	// }

	// Return the client ID
	return client.ID, nil
}

// -----------------------------
// Function: Get Clients
// -----------------------------
// func (c *ClientService) GetClients() ([]*domain.Client, error) {
// 	var clients []*domain.Client

// 	clients, err := c.Repo.QueryClients()
// 	if err != nil {
// 		log.Printf("QueryClient adapter function generates an error: %v", err)
// 		return nil, err_client_no_records
// 	}
// 	return clients, nil
// }

// // -----------------------------
// // Function: Get Client By ID
// // -----------------------------
// func (c *ClientService) GetClientByID(id uuid.UUID) (*domain.Client, error) {
// 	var client *domain.Client

// 	client, err := c.Repo.QueryClientByID(id)
// 	if err != nil {
// 		log.Printf("QueryClient adapter function generates an error: %v", err)
// 		return nil, err_client_creation_error
// 	}
// 	return client, nil
// }

// // -----------------------------
// // Function: Delete a Client
// // -----------------------------
// func (c *ClientService) DeleteClient(id uuid.UUID) error {

// 	err := c.Repo.DeleteClient(id)
// 	if err != nil {
// 		log.Printf("DeleteClient adapter function generates an error: %v", err)
// 		return err_client_creation_error
// 	}
// 	return nil
// }
