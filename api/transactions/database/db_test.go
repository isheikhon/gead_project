package database

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	mock_db "transaction/server/transactions/mock"
)

func TestGetTransactions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "customerid", "productid", "totalprice", "quantity"}).
		AddRow("1", "orange", "1", "1", 20, 2).
		AddRow("1", "apple", "1", "1", 20, 2)

	query := "SELECT (.+) FROM transactions"
	mock.ExpectQuery(query).WillReturnRows(rows)

	transactions, err := mock_db.GetTransactions(context.Background())
	assert.NoError(t, err)
	assert.Len(t, transactions, 2)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

//func TestGetTransaction(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("failed to open sqlmock database: %v", err)
//	}
//	defer db.Close()
//
//	rows := sqlmock.NewRows([]string{"id", "customer_id", "product_id", "price", "quantity", "created_at"}).
//		AddRow("1", "orange", "1", "1", 20, 2, timestamppb.New(time.Now()))
//
//	query := "SELECT (.+) FROM transactions where id = $1"
//	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
//
//	transaction, err := database.MockTransactionData.GetTransactionById(context.Background(), 1)
//	assert.NoError(t, err)
//	assert.Equal(t, "1", transaction.ID)
//	assert.Equal(t, "orange", transaction.Name)
//	assert.Equal(t, "1", transaction.CustomerID)
//	assert.Equal(t, "1", transaction.ProductID)
//	assert.Equal(t, 20, transaction.TotalPrice)
//	assert.Equal(t, int32(2), transaction.Quantity)
//
//	assert.NoError(t, err)
//}

//func TestCreateTransaction(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("failed to open sqlmock database: %v", err)
//	}
//	defer db.Close()
//
//	newTransaction := transaction.Transaction{
//		ID:         "1",
//		Name:       "orange",
//		CustomerID: "1",
//		ProductID:  "1",
//		TotalPrice: 20,
//		Quantity:   2,
//		CreatedAt:  timestamppb.New(time.Now()),
//	}
//
//	query := "INSERT INTO transactions(.+)"
//	mock.ExpectQuery(query).
//		WithArgs(newTransaction.ID, newTransaction.Name, newTransaction.ProductID, newTransaction.TotalPrice, newTransaction.Quantity).
//		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
//
//	transaction, err := database.MockTransactionData.CreateTransaction(context.Background(), newTransaction)
//	assert.NoError(t, err)
//	assert.Equal(t, "1", transaction.ID)
//	assert.Equal(t, "orange", transaction.Name)
//	assert.Equal(t, "1", transaction.CustomerID)
//	assert.Equal(t, "1", transaction.ProductID)
//	assert.Equal(t, 20, transaction.TotalPrice)
//	assert.Equal(t, int32(2), transaction.Quantity)
//
//	err = mock.ExpectationsWereMet()
//	assert.NoError(t, err)
//}
