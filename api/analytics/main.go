package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"transaction/server/analytics/database"
	"transaction/server/proto/analytics"
)

var (
	port = flag.Int("port", 50054, "The server port")
)

type server struct {
	analytics.UnimplementedCustomerServiceServer
}

func (s *server) GetTotalSales(ctx context.Context, req *analytics.GetTotalSalesRequest) (*analytics.TotalSalesResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database.DBGetTotalSales()
}

func (s *server) GetSalesByProduct(ctx context.Context, req *analytics.GetSalesByProductRequest) (*analytics.ProductSalesResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database.DBGetSalesByProduct()
}

func (s *server) GetTopCustomers(ctx context.Context, req *analytics.GetTopCustomersRequest) (*analytics.TopCustomerResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database.DBGetTopCustomers()
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	analytics.RegisterCustomerServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
