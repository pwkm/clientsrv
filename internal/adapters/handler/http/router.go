package http

import (
	"client/internal/utils/env"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/gin-gonic/gin"
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

	// Swagger

	// Client group
	client := router.Group("/client")
	{
		client.POST("/new", clienthandler.RegisterClient)
		// client.GET("/:id", clienthandler.GetClientByID)
		// client.GET("/", clienthandler.ListClients)
		// client.GET("/delete/:id", clienthandler.DeleteClient)
	}

	// Product group
	// product := router.Group("/product")
	// {
	// 	// product.POST("/new", producthandler.CreateProduct)
	// 	// product.GET("/:id", clienthandler.GetClientByID)
	// 	// product.GET("/", clienthandler.ListClients)
	// 	// product.GET("/delete/:id", clienthandler.DeleteClient)
	// }

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
