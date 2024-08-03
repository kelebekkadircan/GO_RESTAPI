package persistence

import (
	"context"
	"errors"
	"fmt"
	"productapp/domain"
	"productapp/persistence/common"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product                    // Get all products
	GetAllProductsByStore(store string) []domain.Product // Get all products by store
	AddProduct(product domain.Product) error             // Add a product
	GetById(id int64) (domain.Product, error)            // Get product by id
	DeleteById(id int64) error                           // Delete product by id
	UpdatePrice(id int64, price float32) error           // Update price of product by id
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

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insert_sql := `INSERT INTO products (name, price, discount, store) VALUES ($1, $2, $3, $4)`

	addNew, err := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		fmt.Printf("Error while adding product to database %v", err)
		return err
	}

	log.Info(fmt.Printf("Product added successfully %v", addNew))

	return nil

}

func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()

	getByIdSql := `Select * from products where id = $1`

	productRow := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)

	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	scanErr := productRow.Scan(&id, &name, &price, &discount, &store)

	if scanErr != nil && scanErr.Error() == common.NOTFOUND {
		return domain.Product{}, errors.New(fmt.Sprintf("Product with id %d not found", productId))

	}

	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil

}

func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()

	_, getErr := productRepository.GetById(productId)

	if getErr != nil {
		return errors.New("Product not found")
	}

	deleteByIdSql := `Delete from products where id = $1`

	_, err := productRepository.dbPool.Exec(ctx, deleteByIdSql, productId)

	if err != nil {
		fmt.Printf("Error while deleting product from database %v", err)
		return errors.New("Error while deleting product")
	}

	log.Info(fmt.Printf("Product deleted successfully %d", productId))

	return nil
}

func (productRepository *ProductRepository) UpdatePrice(productId int64, price float32) error {
	ctx := context.Background()

	_, getErr := productRepository.GetById(productId)

	if getErr != nil {
		return errors.New("Product not found")
	}

	updatePriceSql := `Update products set price = $1 where id = $2`

	_, err := productRepository.dbPool.Exec(ctx, updatePriceSql, price, productId)

	if err != nil {
		fmt.Printf("Error while updating price of product %v", err)
		return errors.New("Error while updating price of product")
	}

	log.Info(fmt.Printf("Price updated successfully %d", productId))

	return nil
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
