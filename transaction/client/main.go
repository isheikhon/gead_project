package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"transaction/server/proto/transaction"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := transaction.NewTransactionServiceClient(conn)
	m, err := client.AddTransaction(context.Background(), &transaction.Transaction{ID: "17168167", Name: "GameBoy", TotalPrice: 132, ProductID: "ye786e8", CustomerID: "shbjhds67dg", Quantity: 1, CreatedAt: timestamppb.Now()})
	if err != nil {
		log.Fatalf("failed to get mesasge: %v", err)
	}
	log.Printf("book list: %v", m)
}
