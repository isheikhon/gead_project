package database

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // <------------ here
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
	"transaction/server/proto/transaction"
)

var db, err = sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

func DBGetTransactionByID(req *transaction.TransactionQueryRequest) (*transaction.TransactionResponse, error) {
	trans := transaction.Transaction{}
	var created time.Time
	query := "SELECT * FROM transaction where id = $1"
	err := db.QueryRow(query, req.ID).Scan(&trans.ID, &trans.Name, &trans.TotalPrice,
		&trans.ProductID, &trans.CustomerID, &trans.Quantity, &created)
	if err != nil {
		panic(err)
	}
	return &transaction.TransactionResponse{Transaction: &transaction.Transaction{ID: trans.ID, Name: trans.Name, TotalPrice: trans.TotalPrice, ProductID: trans.ProductID, CustomerID: trans.CustomerID, Quantity: trans.Quantity, CreatedAt: timestamppb.New(created)}}, nil
}

func DBAddTransaction(req *transaction.Transaction) (*transaction.TransactionResponse, error) {
	log.Printf("Inserting a Row in to DB")
	var UUID = uuid.New().String()
	//Inserting a Row in to DB.
	sqlStatement := `INSERT INTO transaction (id,name,totalprice,productid,customerid,quantity,created_at) VALUES ($1, $2, $3, $4,$5,$6,$7)`
	_, err = db.Exec(sqlStatement, UUID, req.Name, req.TotalPrice*req.Quantity, req.ProductID, req.CustomerID, req.Quantity, time.Now())
	if err != nil {
		panic(err)
	}
	return &transaction.TransactionResponse{Transaction: &transaction.Transaction{ID: UUID, Name: req.Name, TotalPrice: req.TotalPrice * req.Quantity, ProductID: req.ProductID, CustomerID: req.CustomerID, Quantity: req.Quantity, CreatedAt: req.CreatedAt}}, nil

}

func DBGetAllTransaction(req *transaction.AllTransactionQueryRequest) (*transaction.TransactionListResponse, error) {
	var trans []*transaction.Transaction
	var created time.Time
	query := "SELECT * FROM transaction"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		tra := transaction.Transaction{}
		rows.Scan(&tra.ID, &tra.Name, &tra.TotalPrice,
			&tra.ProductID, &tra.CustomerID, &tra.Quantity, &created)
		trans = append(trans, &transaction.Transaction{ID: tra.ID, Name: tra.Name, TotalPrice: tra.TotalPrice, ProductID: tra.ProductID, CustomerID: tra.CustomerID, Quantity: tra.Quantity, CreatedAt: timestamppb.New(created)})
	}
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return &transaction.TransactionListResponse{Transactions: trans}, nil
}
