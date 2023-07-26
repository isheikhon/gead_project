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

func DBGetTransactionByID(req *transaction.FindTransactionByIDRequest) (*transaction.FindTransactionByIDResponse, error) {
	trans := transaction.Transaction{}
	var created time.Time
	query := "SELECT * FROM transaction where id = $1"
	err := db.QueryRow(query, req.Id).Scan(&trans.Id, &trans.Name, &trans.TotalPrice,
		&trans.ProductId, &trans.CustomerId, &trans.Quantity, &created)
	if err != nil {
		panic(err)
	}
	return &transaction.FindTransactionByIDResponse{Transaction: &transaction.Transaction{Id: trans.Id, Name: trans.Name, TotalPrice: trans.TotalPrice, ProductId: trans.ProductId, CustomerId: trans.CustomerId, Quantity: trans.Quantity, CreatedAt: timestamppb.New(created)}}, nil
}

func DBAddTransaction(req *transaction.Transaction) (*transaction.AddTransactionResponse, error) {
	log.Printf("Inserting a Row in to DB")
	var UUID = uuid.New().String()
	//Inserting a Row in to DB.
	sqlStatement := `INSERT INTO transaction (id,name,totalprice,productid,customerid,quantity,created_at) VALUES ($1, $2, $3, $4,$5,$6,$7)`
	_, err = db.Exec(sqlStatement, UUID, req.Name, req.TotalPrice*req.Quantity, req.ProductId, req.CustomerId, req.Quantity, time.Now())
	if err != nil {
		panic(err)
	}
	return &transaction.AddTransactionResponse{Transaction: &transaction.Transaction{Id: UUID, Name: req.Name, TotalPrice: req.TotalPrice * req.Quantity, ProductId: req.ProductId, CustomerId: req.CustomerId, Quantity: req.Quantity, CreatedAt: req.CreatedAt}}, nil

}

func DBGetAllTransaction(req *transaction.GetAllTransactionRequest) (*transaction.GetAllTransactionResponse, error) {
	var trans []*transaction.Transaction
	var created time.Time
	query := "SELECT * FROM transaction"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		tra := transaction.Transaction{}
		rows.Scan(&tra.Id, &tra.Name, &tra.TotalPrice,
			&tra.ProductId, &tra.CustomerId, &tra.Quantity, &created)
		trans = append(trans, &transaction.Transaction{Id: tra.Id, Name: tra.Name, TotalPrice: tra.TotalPrice, ProductId: tra.ProductId, CustomerId: tra.CustomerId, Quantity: tra.Quantity, CreatedAt: timestamppb.New(created)})
	}
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return &transaction.GetAllTransactionResponse{Transactions: trans}, nil
}
