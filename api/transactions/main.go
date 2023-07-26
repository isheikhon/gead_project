package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"transaction/server/proto/transaction"
	database_ "transaction/server/transactions/database"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	transaction.TransactionServiceServer
}

func (s *server) AddTransaction(ctx context.Context, req *transaction.Transaction) (*transaction.TransactionResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database_.DBAddTransaction(req)
}

func (s *server) GetAllTransaction(ctx context.Context, req *transaction.AllTransactionQueryRequest) (*transaction.TransactionListResponse, error) {
	return database_.DBGetAllTransaction(req)
}
func (s *server) FindTransactionByID(ctx context.Context, req *transaction.TransactionQueryRequest) (*transaction.TransactionResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database_.DBGetTransactionByID(req)
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
