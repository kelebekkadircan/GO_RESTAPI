package service

import (
	"productapp/domain"
	"productapp/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeProductRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeProductRepository.products = append(fakeProductRepository.products,
		domain.Product{
			Id:       int64(len(fakeProductRepository.products) + 1),
			Name:     product.Name,
			Price:    product.Price,
			Discount: product.Discount,
			Store:    product.Store,
		})
	return nil
}

func (fakeProductRepository *FakeProductRepository) DeleteById(productId int64) error {
	for i, product := range fakeProductRepository.products {
		if product.Id == productId {
			fakeProductRepository.products = append(fakeProductRepository.products[:i], fakeProductRepository.products[i+1:]...)
			return nil
		}
	}
	return nil
}

func (fakeProductRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	for _, product := range fakeProductRepository.products {
		if product.Id == productId {
			return product, nil
		}
	}
	return domain.Product{}, nil
}

func (fakeProductRepository *FakeProductRepository) UpdatePrice(productId int64, price float32) error {
	for i, product := range fakeProductRepository.products {
		if product.Id == productId {
			fakeProductRepository.products[i].Price = price
			return nil
		}
	}
	return nil
}

func (fakeProductRepository *FakeProductRepository) GetAllProducts() []domain.Product {

	return fakeProductRepository.products
}

func (fakeProductRepository *FakeProductRepository) GetAllProductsByStore(store string) []domain.Product {
	var productsByStore []domain.Product
	for _, product := range fakeProductRepository.products {
		if product.Store == store {
			productsByStore = append(productsByStore, product)
		}
	}
	return productsByStore
}
