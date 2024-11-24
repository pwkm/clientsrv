package main

import (
	"client/internal/adapters/handler/http"
	"client/internal/adapters/infra/posgresdb"
	"client/internal/adapters/infra/posgresdb/repository"
	"client/internal/adapters/infra/rabbitmsg"
	"client/internal/core/domain"
	"client/internal/core/service"
	"client/internal/utils/env"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// Read Config file into environement container env
	env := env.NewEnv()

	// Setup a database
	db := posgresdb.Database(env)
	// Closing database connection
	// defer db.  .Close()

	// Database initialisation
	db.AutoMigrate(&domain.Client{}, &domain.Login{}, &domain.Profile{})

	// Setup RabbitMQ
	str, err := rabbitmsg.NewRabbitStream(env)
	if err != nil {
		log.Fatal("Error initiating Rabbit: ", err)
	}
	defer str.Channel.Close()
	defer str.Con.Close()

	// setup a Context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(env.ContextTimeout)*time.Second)
	defer cancel()

	// setup different services
	clientRepo := repository.NewClientRepository(db)
	clientservice := service.NewClientService(clientRepo, str)
	clientHandler := http.NewClientHandler(clientservice)

	// productRepo := repository.NewProductRepository(db)
	// productservice := service.NewProductService(productRepo, str)
	// productHandler := http.NewProductHandler(productservice)

	// Init Router
	router, err := http.NewRouter(
		env,
		*clientHandler,
		// *productHandler,
	)
	if err != nil {
		log.Fatal("Error initiating router ; ", err)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", env.ServerURL, env.ServerPort)

	err = router.Serve(listenAddr)
	if err != nil {
		log.Fatal("error starting up the http server ", err)
	}
}
