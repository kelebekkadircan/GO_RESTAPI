package infrastructer

import (
	"context"
	"fmt"
	"os"
	"productapp/common/postgresql"
	"productapp/domain"
	"productapp/persistence"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

/*
func TestAdd(t *testing.T) {

t.Run("TestAdd", func(t *testing.T) {

actual := Add(1, 2)
expected := 3

assert.Equal(t, expected, actual)

})

}

func Add(x int, y int) int {

return x + y
}
*/

var productRepository persistence.IProductRepository

var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {

	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	productRepository = persistence.NewProductRepository(dbPool)
	fmt.Println("Before running the test cases")
	exitCode := m.Run()
	fmt.Println("AFTER running the test cases")
	os.Exit(exitCode)

}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}
func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {

	setup(ctx, dbPool)

	fmt.Println("TestGetAllProducts")

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}

	t.Run("GetAllProducts", func(t *testing.T) {
		actual := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actual))
		assert.Equal(t, expectedProducts, actual)
	})

	clear(ctx, dbPool)

}

func TestGetAllProductsByStore(t *testing.T) {
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}
	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}

func TestAddProductRepository(t *testing.T) {

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
	}

	newProduct := domain.Product{
		Name:     "AirFryer",
		Price:    3000.0,
		Discount: 22.0,
		Store:    "ABC TECH",
	}

	t.Run("AddProduct", func(t *testing.T) {
		productRepository.AddProduct(newProduct)
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}

func TestGetProductById(t *testing.T) {

	setup(ctx, dbPool)

	t.Run("GetProductById", func(t *testing.T) {
		actualProduct, _ := productRepository.GetById(1)
		assert.Equal(t, domain.Product{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		}, actualProduct)
	})

	clear(ctx, dbPool)
}

func TestDeleteById(t *testing.T) {

	setup(ctx, dbPool)

	t.Run("GetProductById", func(t *testing.T) {
		productRepository.DeleteById(1)
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 3, len(actualProducts))
	})

	clear(ctx, dbPool)
}

func TestUpdatePrice(t *testing.T) {

	setup(ctx, dbPool)

	t.Run("UpdatePrice", func(t *testing.T) {
		productBeforeUpdate, _ := productRepository.GetById(1)
		assert.Equal(t, domain.Product{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		}, productBeforeUpdate)
		err := productRepository.UpdatePrice(1, 4000.0)
		if err != nil {
			fmt.Println("Error while updating price")
		}
		productAfterUpdate, _ := productRepository.GetById(1)
		assert.Equal(t, domain.Product{
			Id:       1,
			Name:     "AirFryer",
			Price:    4000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		}, productAfterUpdate)

	})

	clear(ctx, dbPool)

}
