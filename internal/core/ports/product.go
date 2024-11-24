package ports

type IProductRepository interface {
	// SaveProduct(product *domain.Product) error
	// QueryClients() ([]*domain.Client, error)
	// QueryClientByID(id uuid.UUID) (*domain.Client, error)
	// DeleteClient(id uuid.UUID) error
}

// type IMessageStream interface {
// 	SendMessage(c *domain.Client) error
// }

type IProductService interface {
	// CreateProduct(name, description string, ac int) (uuid.UUID, error)

	// GetClients() ([]*domain.Client, error)
	// DeleteClient(id uuid.UUID) error
	// GetClientByID(id uuid.UUID) (*domain.Client, error)
}
