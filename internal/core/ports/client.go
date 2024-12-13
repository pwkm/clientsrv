package ports

import (
	"github.com/google/uuid"
	"github.com/pwkm/clientsrv/internal/core/domain"
)

type IClientRepository interface {
	SaveClient(client *domain.Client) error
	QueryClients() ([]*domain.Client, error)
	QueryClientByID(id uuid.UUID) (*domain.Client, error)
	DeleteClient(id uuid.UUID) error
}

type IMessageStream interface {
	SendMessage(c *domain.Client) error
}

type IClientService interface {
	RegisterClient(name, email, password string) (uuid.UUID, error)
	GetClients() ([]*domain.Client, error)
	DeleteClient(id uuid.UUID) error
	GetClientByID(id uuid.UUID) (*domain.Client, error)
}
