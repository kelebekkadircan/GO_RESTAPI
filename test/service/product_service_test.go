package service

import (
	"os"
	"productapp/domain"
	"productapp/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productionService service.IProductService

func TestMain(m *testing.M) {

	initialProducts := []domain.Product{
		{
			Id:       1,
			Name:     "Product 1",
			Price:    100,
			Discount: 10,
			Store:    "Store 1",
		},
		{
			Id:       2,
			Name:     "Product 2",
			Price:    200,
			Discount: 20,
			Store:    "Store 2",
		},
	}

	fakeProductRepository := NewFakeProductRepository(initialProducts)
	productionService = service.NewProductService(fakeProductRepository)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestAdd(t *testing.T) {

	t.Run("Test Add Product", func(t *testing.T) {
		actualProducts := productionService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))
	})
}
