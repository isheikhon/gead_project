package database

import (
	"context"
	"transaction/server/proto/product"
)

type MockTransactionData struct{}

func (m *MockTransactionData) GetProducts(ctx context.Context) ([]product.Product, error) {
	return []product.Product{
		{
			ID:    "1",
			Name:  "orange",
			Price: 12,
		}, {
			ID:    "2",
			Name:  "apple",
			Price: 12,
		},
	}, nil
}

func (m *MockTransactionData) GetProductById(ctx context.Context, id int32) (product.Product, error) {
	return product.Product{
		ID:    "2",
		Name:  "apple",
		Price: 12,
	}, nil
}

func (m *MockTransactionData) CreateProduct(ctx context.Context, req product.Product) (product.Product, error) {
	req.ID = "90329"
	req.Name = "Banana"
	req.Price = 6
	return req, nil
}
