package mock

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"transaction/server/proto/transaction"
)

type MockTransactionData struct{}

func GetTransactions(ctx context.Context) ([]transaction.Transaction, error) {
	return []transaction.Transaction{
		{
			ID:         "1",
			Name:       "orange",
			CustomerID: "2",
			ProductID:  "2",
			TotalPrice: 100,
			Quantity:   15,
			CreatedAt:  timestamppb.New(time.Now()),
		}, {
			ID:         "2",
			Name:       "orange",
			CustomerID: "3",
			ProductID:  "3",
			TotalPrice: 130,
			Quantity:   1,
			CreatedAt:  timestamppb.New(time.Now()),
		},
	}, nil
}

func (m *MockTransactionData) GetTransactionById(ctx context.Context, id int32) (transaction.Transaction, error) {
	return transaction.Transaction{
		ID:         "2",
		Name:       "orange",
		CustomerID: "3",
		ProductID:  "3",
		TotalPrice: 130,
		Quantity:   1,
		CreatedAt:  timestamppb.New(time.Now()),
	}, nil
}

func (m *MockTransactionData) CreateTransaction(ctx context.Context, req transaction.Transaction) (transaction.Transaction, error) {
	req.ID = "90329"
	req.Name = "orange"
	req.TotalPrice = 250
	req.ProductID = "432"
	req.CustomerID = "4343"
	req.Quantity = 4
	req.CreatedAt = timestamppb.New(time.Now())
	return req, nil
}
