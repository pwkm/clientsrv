package service

// import (
// 	"client2/internal/core/domain"
// 	"client2/internal/core/ports"
// 	"errors"
// 	"log"

// 	"github.com/google/uuid"
// )

// // ------------------------------
// // Global ERROR Messages
// // ------------------------------
// var (
// 	err_product_name_empty     = errors.New("productname could not be empty")
// 	err_product_creation_error = errors.New("product could not be created ")
// 	err_product_no_records     = errors.New("No records found")
// 	err_product_desc_empty     = errors.New("productdescripttion could not be empty")
// )

// // -----------------------------
// // Product Service Structure
// // -----------------------------
// type ProductService struct {
// 	Repo   ports.IProductRepository
// 	Stream ports.IMessageStream
// }

// func NewProductService(repo ports.IProductRepository, str ports.IMessageStream) *ProductService {
// 	return &ProductService{
// 		Repo:   repo,
// 		Stream: str,
// 	}
// }

// // ---------------------------------
// // function: create a new product
// // ---------------------------------

// func (p *ProductService) CreateProduct(name, description string, ac int) (uuid.UUID, error) {
// 	// validate information
// 	if name == "" {
// 		return uuid.Nil, err_product_name_empty
// 	}

// 	if description == "" {
// 		return uuid.Nil, err_product_desc_empty
// 	}

// 	// Create a product
// 	product, err := domain.NewProduct(name, description, ac)
// 	if err != nil {
// 		log.Println("can not create a product: ", err)
// 		return uuid.Nil, err_product_creation_error
// 	}
// 	// Save product
// 	err = p.Repo.SaveProduct(product)
// 	if err != nil {
// 		log.Println("can not save a product: ", err)
// 		return uuid.Nil, err_product_creation_error
// 	}

// 	return product.ID, nil
// }
