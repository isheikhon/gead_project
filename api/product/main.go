package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	database_ "transaction/server/product/database"
	"transaction/server/proto/product"
)

var (
	port = flag.Int("port", 50053, "The server port")
)

type server struct {
	product.UnimplementedProductServiceServer
}

func (s *server) AddProduct(ctx context.Context, req *product.Product) (*product.AddProductResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database_.DBAddProduct(req)
}

func (s *server) GetAllProducts(ctx context.Context, req *product.GetAllProductsRequest) (*product.GetAllProductsResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database_.DBGetAllProducts(req)
}

func (s *server) FindProductByID(ctx context.Context, req *product.FindProductByIDRequest) (*product.FindProductByIDResponse, error) {
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return database_.DBGetProductByID(req)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	product.RegisterProductServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
