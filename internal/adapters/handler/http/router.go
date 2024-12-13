package http

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pwkm/clientsrv/internal/adapters/infra/monitoring"
	"github.com/pwkm/clientsrv/internal/utils/env"
)

// Router is a wrapper voor HTTP Router
type Router struct {
	*gin.Engine
}

// func NewRouter(env *env.Env, clienthandler ClientHandler, producthandler ProductHandler) (*Router, error) {
func NewRouter(env *env.Env, clienthandler ClientHandler) (*Router, error) {
	if env.GinMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// JWT Validator
	// keyFunc := func(ctx context.Context) (interface{}, error) {
	// 	return []byte("secret"), nil
	// }

	// jwtvalidator, _ := validator.New(keyFunc, validator.HS256, "http://localhost:8080/client", []string{"api:read"})
	// jwtMiddleware := jwtmiddleware.New(jwtvalidator.ValidateToken)
	// router.Use(adapter.Wrap(jwtMiddleware.CheckJWT))

	// midleware prometheus
	router.Use(monitoring.PrometheusMiddleware())

	// Register metrics
	router.GET("/metrics", monitoring.PrometheusHandler())

	// Client group
	client := router.Group("/client")
	{
		client.POST("/new", clienthandler.RegisterClient)
		client.GET("/:id", clienthandler.GetClientByID)
		client.GET("/", clienthandler.ListClients)
		client.DELETE("/:id", clienthandler.DeleteClient)
	}

	return &Router{
		router,
	}, nil
}

// Serve start the http Listener
func (r *Router) Serve(listenAddr string) error {
	fmt.Println("server starting")
	return r.Run(listenAddr)
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
