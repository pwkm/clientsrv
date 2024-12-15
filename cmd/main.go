package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pwkm/clientsrv/internal/adapters/handler/http"
	"github.com/pwkm/clientsrv/internal/adapters/infra/posgresdb"
	"github.com/pwkm/clientsrv/internal/adapters/infra/posgresdb/repository"
	"github.com/pwkm/clientsrv/internal/core/domain"
	"github.com/pwkm/clientsrv/internal/core/service"
	"github.com/pwkm/clientsrv/internal/utils/env"
)

func main() {
	// Read Config file into environement container env
	env := env.NewEnv()

	// Determine port for HTTP service.
	env.ServerPort = os.Getenv("PORT")
	if env.ServerPort == "" {
		env.ServerPort = "8080"
		log.Printf("defaulting to port %s", env.ServerPort)
	}

	// Setup a database
	db := posgresdb.Database(env)
	// Closing database connection
	// defer db.  .Close()

	// Database initialisation
	db.AutoMigrate(&domain.Client{}, &domain.Login{}, &domain.Profile{})

	// // Setup RabbitMQ
	// str, err := rabbitmsg.NewRabbitStream(env)
	// if err != nil {
	// 	log.Fatal("Error initiating Rabbit: ", err)
	// }
	// defer str.Channel.Close()
	// defer str.Con.Close()

	// setup a Context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(env.ContextTimeout)*time.Second)
	defer cancel()

	// setup different services
	clientRepo := repository.NewClientRepository(db)
	clientservice := service.NewClientService(clientRepo)
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
	// Use the FQHN for TLS
	certFile := "./server.crt"
	keyFile := "./server.key"

	// err = router.Serve(listenAddr)
	err = router.RunTLS(listenAddr, certFile, keyFile)
	if err != nil {
		log.Fatal("error starting up the http server ", err)
	}
}
