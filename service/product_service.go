package service

import (
	"productapp/domain"
	"productapp/persistence"
	"productapp/service/model"

	"github.com/pkg/errors"
)

type IProductService interface {
	Add(productCreate model.ProductCreate) error
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, price float32) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(store string) []domain.Product
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (productService *ProductService) Add(productCreate model.ProductCreate) error {

	validateErr := validateProductCreate(productCreate)
	if validateErr != nil {
		return validateErr
	}

	product := domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	}
	return productService.productRepository.AddProduct(product)
}

func (productService *ProductService) DeleteById(productId int64) error {
	return productService.productRepository.DeleteById(productId)
}

func (productService *ProductService) GetById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetById(productId)
}

func (productService *ProductService) UpdatePrice(productId int64, price float32) error {

	return productService.productRepository.UpdatePrice(productId, price)
}

func (productService *ProductService) GetAllProducts() []domain.Product {
	return productService.productRepository.GetAllProducts()
}

func (productService *ProductService) GetAllProductsByStore(store string) []domain.Product {
	return productService.productRepository.GetAllProductsByStore(store)
}

func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Name == "" {
		return errors.New("Product name is required")
	}
	if productCreate.Price <= 0 {
		return errors.New("Product price should be greater than 0")
	}
	if productCreate.Store == "" {
		return errors.New("Product store is required")
	}
	if productCreate.Discount > 70 {
		return errors.New("Product discount should be less than 70%")
	}

	return nil
}
