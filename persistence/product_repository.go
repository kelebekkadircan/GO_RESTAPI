package persistence

import (
	"context"
	"fmt"
	"productapp/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product // Get all products
	// GetProductById(id int64) domain.Product // Get product by id
	GetAllProductsByStore(store string) []domain.Product // Get all products by store

}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {

	return &ProductRepository{
		dbPool: dbPool,
	}

}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products") // Query to get all products

	if err != nil {
		// log.Error("Error while getting products from database %v", err)
		// log.Error("Error while getting products from database %v", err)
		fmt.Printf("Error while getting products from database %v", err)
	}

	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})

	}
	return products

}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()

	getProductsByStoreNameSql := `Select * from products where store = $1`

	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)

	if err != nil {
		// log.Error("Error while getting products from database %v", err)
		fmt.Printf("Error while getting products from database %v", err)
	}

	// return extractProductsFromRows(productRows)

	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})

	}
	return products

}

// func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
// 	var products = []domain.Product{}
// 	var id int64
// 	var name string
// 	var price float32
// 	var discount float32
// 	var store string

// 	for productRows.Next() {
// 		productRows.Scan(&id, &name, &price, &discount, &store)
// 		products = append(products, domain.Product{
// 			Id:       id,
// 			Name:     name,
// 			Price:    price,
// 			Discount: discount,
// 			Store:    store,
// 		})
// 	}
// 	return products
// }
