package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
	"transaction/server/proto/transaction"
)

var (
	port = flag.Int("port", 50051, "The server port")
)
var db, err = sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

type server struct {
	transaction.UnimplementedTransactionServiceServer
}

func (s *server) AddTransaction(ctx context.Context, req *transaction.Transaction) (*transaction.TransactionResponse, error) {
	log.Printf("Connectin to DB")
	// Connect to the "company_db" database.
	log.Printf("Inserting a Row in to DB")
	//Inserting a Row in to DB.
	sqlStatement := `INSERT INTO transaction (id,name,totalprice,productid,customerid,quantity,created_at) VALUES ($1, $2, $3, $4,$5,$6,$7)`
	_, err = db.Exec(sqlStatement, req.ID, req.Name, 100, req.ProductID, req.CustomerID, req.Quantity, time.Now())
	if err != nil {
		panic(err)
	}
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return &transaction.TransactionResponse{Transaction: &transaction.Transaction{ID: req.ID, Name: req.Name, TotalPrice: 100, ProductID: req.ProductID, CustomerID: req.CustomerID, Quantity: req.Quantity, CreatedAt: req.CreatedAt}}, nil
}

func (s *server) GetAllTransaction(ctx context.Context, req *transaction.AllTransactionQueryRequest) (*transaction.TransactionListResponse, error) {
	var trans []*transaction.Transaction
	var created time.Time
	log.Print("Log, log")
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
	log.Print(trans)
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return &transaction.TransactionListResponse{Transactions: trans}, nil
}
func (s *server) FindTransactionByID(ctx context.Context, req *transaction.TransactionQueryRequest) (*transaction.TransactionResponse, error) {
	trans := transaction.Transaction{}
	var created time.Time
	query := "SELECT * FROM transaction where id = $1"
	err := db.QueryRow(query, req.ID).Scan(&trans.ID, &trans.Name, &trans.TotalPrice,
		&trans.ProductID, &trans.CustomerID, &trans.Quantity, &created)
	if err != nil {
		panic(err)
	}
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return &transaction.TransactionResponse{Transaction: &transaction.Transaction{ID: trans.ID, Name: trans.Name, TotalPrice: trans.TotalPrice, ProductID: trans.ProductID, CustomerID: trans.CustomerID, Quantity: trans.Quantity, CreatedAt: timestamppb.New(created)}}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	transaction.RegisterTransactionServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
