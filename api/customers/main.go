package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	database_ "transaction/server/customers/database"
	"transaction/server/proto/customer"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	customer.UnimplementedCustomerServiceServer
}

func (s *server) AddCustomer(ctx context.Context, req *customer.Customer) (*customer.AddCustomerResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database_.DBAddCustomer(req)
}
func (s *server) GetAllCustomers(ctx context.Context, req *customer.GetAllCustomersRequest) (*customer.GetAllCustomersResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database_.DBGetAllCustomer(req)
}

func (s *server) FindCustomerByID(ctx context.Context, req *customer.FindCustomerByIDRequest) (*customer.FindCustomerByIDResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database_.DBGetCustomerByID(req)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	customer.RegisterCustomerServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
