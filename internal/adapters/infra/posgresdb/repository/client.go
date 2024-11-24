package repository

import (
	"client/internal/core/domain"
	"log"

	"gorm.io/gorm"
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

	// // Save login
	// insertStmt = `insert into login(id, login, password) values($1, $2, $3)`
	// _, err = cr.db.Exec(insertStmt, client.ID, client.Login.Email, client.Login.Password)
	// if err != nil {
	// 	log.Printf("error saving login: %v", err)
	// 	return err
	// }

	// // Save profile
	// insertStmt = `insert into profile(id, street, number, postcode, community, email, birthday) values($1, $2, $3, $4, $5, $6, $7)`
	// _, err = cr.db.Exec(insertStmt, client.ID, client.Profile.Street, client.Profile.Number, client.Profile.Postcode, client.Profile.Community, client.Profile.Email, client.Profile.Birthday)
	// if err != nil {
	// 	log.Printf("error saving profile: %v", err)
	// 	return err
	// }

	// commit transactions
	result = tx.Commit()
	return result.Error
}

// // ------------------------------------
// // Func: Query Clients
// // ------------------------------------
// func (cr *ClientRepository) QueryClients() ([]*domain.Client, error) {
// 	var clients []*domain.Client
// 	var id uuid.UUID
// 	var id2 uuid.UUID
// 	var name string

// 	// Query all clients
// 	querystm1 := `SELECT id, name FROM client`
// 	result, err := cr.db.Query(querystm1)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer result.Close()

// 	for result.Next() {
// 		// Query client
// 		err = result.Scan(&id, &name)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Query login
// 		var login = new(domain.Login)
// 		querystm2 := `SELECT id, login, password FROM login WHERE id=$1;`
// 		row2 := cr.db.QueryRow(querystm2, id)
// 		err := row2.Scan(&id2, &login.Email, &login.Password)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if err == sql.ErrNoRows {
// 			return nil, errors.New("no rows were found")
// 		}

// 		// Query profile
// 		var profile = new(domain.Profile)
// 		querystm3 := `SELECT id, street, number, postcode, community, email, birthday FROM profile WHERE id=$1`
// 		row3 := cr.db.QueryRow(querystm3, id)
// 		err = row3.Scan(&id2, &profile.Street, &profile.Number, &profile.Postcode, &profile.Community, &profile.Email, &profile.Birthday)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if err == sql.ErrNoRows {
// 			return nil, errors.New("no rows were found")
// 		}

// 		// Create a client struct
// 		client := domain.Client{
// 			ID:      id,
// 			Name:    name,
// 			Login:   login,
// 			Profile: profile,
// 		}
// 		// append the client to the array
// 		clients = append(clients, &client)
// 	}

// 	return clients, nil
// }

// // ------------------------------------
// // Func: Retrieve Client By ID
// // ------------------------------------
// func (cr *ClientRepository) QueryClientByID(id uuid.UUID) (*domain.Client, error) {
// 	var client = new(domain.Client)
// 	var id2 uuid.UUID

// 	// Query the client
// 	querystm := `SELECT id, name FROM client WHERE id=$1`
// 	result := cr.db.QueryRow(querystm, id)
// 	err := result.Scan(&client.ID, &client.Name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Query login
// 	var login = new(domain.Login)
// 	querystm2 := `SELECT id, login, password FROM login WHERE id=$1;`
// 	row2 := cr.db.QueryRow(querystm2, id)
// 	err = row2.Scan(&id2, &login.Email, &login.Password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Query profile
// 	var profile = new(domain.Profile)
// 	querystm3 := `SELECT id, street, number, postcode, community, email, birthday FROM profile WHERE id=$1`
// 	row3 := cr.db.QueryRow(querystm3, id)
// 	err = row3.Scan(&id2, &profile.Street, &profile.Number, &profile.Postcode, &profile.Community, &profile.Email, &profile.Birthday)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create a client struct
// 	client.Login = login
// 	client.Profile = profile
// 	return client, nil
// }

// // ------------------------------------
// // Func: Delete Client
// // ------------------------------------
// func (cr *ClientRepository) DeleteClient(id uuid.UUID) error {

// 	// Delete client
// 	deleteStmt := `DELETE FROM client WHERE id = $1`
// 	_, err := cr.db.Exec(deleteStmt, id)
// 	if err != nil {
// 		log.Printf("error saving client: %v", err)
// 		return err
// 	}

// 	deleteStmt = `DELETE FROM login WHERE id = $1`
// 	_, err = cr.db.Exec(deleteStmt, id)
// 	if err != nil {
// 		log.Printf("error saving client: %v", err)
// 		return err
// 	}

// 	deleteStmt = `DELETE FROM profile WHERE id = $1`
// 	_, err = cr.db.Exec(deleteStmt, id)
// 	if err != nil {
// 		log.Printf("error saving client: %v", err)
// 		return err
// 	}

// 	return nil
// }
