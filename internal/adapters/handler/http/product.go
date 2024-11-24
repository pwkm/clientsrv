package http

// import (
// 	"client/internal/core/ports"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // Data transfer object NewUser
// type NewProduct struct {
// 	Name               string `json:"name"`
// 	Description        string `json:"description"`
// 	AnnualContribution int    `json:"annualcontribution"`
// }

// // -------------------------------
// // ClientHandler
// // -------------------------------
// type ProductHandler struct {
// 	svc ports.IProductService
// }

// func NewProductHandler(productService ports.IProductService) *ProductHandler {
// 	return &ProductHandler{
// 		svc: productService,
// 	}
// }

// // --------------------------------
// // Function:  Create Product
// // --------------------------------
// func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
// 	var NewProduct NewProduct
// 	if err := ctx.ShouldBindJSON(&NewProduct); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	// Call service
// 	id, err := h.svc.CreateProduct(NewProduct.Name, NewProduct.Description, NewProduct.AnnualContribution)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	// Reply status
// 	ctx.JSON(http.StatusOK, id)
// }
