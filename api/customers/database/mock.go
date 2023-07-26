package database

import (
	"context"
	"transaction/server/proto/customer"
)

type MockTransactionData struct{}

func (m *MockTransactionData) GetCustomer(ctx context.Context) ([]customer.Customer, error) {
	return []customer.Customer{
		{
			ID:   "1",
			Name: "Suzan",
		}, {
			ID:   "2",
			Name: "Tephany",
		},
	}, nil
}

func (m *MockTransactionData) GetCustomerById(ctx context.Context, id int32) (customer.Customer, error) {
	return customer.Customer{
		ID:   "2",
		Name: "Stephan",
	}, nil
}

func (m *MockTransactionData) CreateCustomer(ctx context.Context, req customer.Customer) (customer.Customer, error) {
	req.ID = "90329"
	req.Name = "Hannah"
	return req, nil
}
