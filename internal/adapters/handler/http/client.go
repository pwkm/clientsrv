package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pwkm/clientsrv/internal/core/ports"
)

// Data transfer object NewUser
type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// -------------------------------
// ClientHandler
// -------------------------------
type ClientHandler struct {
	svc ports.IClientService
}

func NewClientHandler(clientService ports.IClientService) *ClientHandler {
	return &ClientHandler{
		svc: clientService,
	}
}

// --------------------------------
// Function:  Register Client
// --------------------------------
func (h *ClientHandler) RegisterClient(ctx *gin.Context) {
	var newUser NewUser
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Call service
	id, err := h.svc.RegisterClient(newUser.Name, newUser.Email, newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Reply status
	ctx.JSON(http.StatusOK, id)
}

// --------------------------------
// Function:  Get Client By ID
// --------------------------------
func (h *ClientHandler) GetClientByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Call service
	client, err := h.svc.GetClientByID(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Reply status
	ctx.JSON(http.StatusOK, client)
}

// --------------------------------
// Function:  List Clients
// --------------------------------
func (h *ClientHandler) ListClients(ctx *gin.Context) {
	// Call service
	clients, err := h.svc.GetClients()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Reply status
	ctx.JSON(http.StatusOK, clients)
}

// --------------------------------
// Function:  Delete Client
// --------------------------------
func (h *ClientHandler) DeleteClient(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// call service
	err = h.svc.DeleteClient(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Reply status
	ctx.JSON(http.StatusOK, "client deleted ")
}

// // --------------------------------
// // Function:  Update Client
// // --------------------------------
// func (h *ClientHandler) UpdateClient(ctx *gin.Context) {
// 	// id := ctx.Param("id")
// 	// uid, err := uuid.Parse(id)
// 	// if err != nil {
// 	// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 	// 		"error": err.Error(),
// 	// 	})
// 	// 	return
// 	// }
// 	// // call service
// 	// err = h.svc.DeleteClient(uid)
// 	// if err != nil {
// 	// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 	// 		"error": err.Error(),
// 	// 	})
// 	// 	return
// 	// }
// 	// Reply status
// 	ctx.JSON(http.StatusOK, "ok")

// }
