package database

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // <------------ here
	"log"
	"transaction/server/proto/product"
)

var db, err = sql.Open("postgres", "postgresql://root@localhost:26258/defaultdb?sslmode=disable")

func DBGetProductByID(req *product.FindProductByIDRequest) (*product.FindProductByIDResponse, error) {
	pr := product.Product{}
	query := "SELECT * FROM product where id = $1"
	err := db.QueryRow(query, req.Id).Scan(&pr.Id, &pr.Name, &pr.Price)
	if err != nil {
		panic(err)
	}
	return &product.FindProductByIDResponse{Product: &product.Product{Id: pr.Id, Name: pr.Name, Price: pr.Price}}, nil
}

func DBAddProduct(req *product.Product) (*product.AddProductResponse, error) {
	log.Printf("Inserting a Row in to DB")
	var UUID = uuid.New().String()
	//Inserting a Row in to DB.
	sqlStatement := `INSERT INTO product (id,name,price) VALUES ($1, $2,$3)`
	_, err = db.Exec(sqlStatement, UUID, req.Name, req.Price)
	if err != nil {
		panic(err)
	}
	return &product.AddProductResponse{Product: &product.Product{Id: UUID, Name: req.Name, Price: req.Price}}, nil
}

func DBGetAllProducts(req *product.GetAllProductsRequest) (*product.GetAllProductsResponse, error) {
	var products []*product.Product
	var db, err = sql.Open("postgres", "postgresql://root@localhost:26258/defaultdb?sslmode=disable")
	log.Printf("Received DB: %v", err)

	query := "SELECT * FROM product"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		pr := product.Product{}
		rows.Scan(&pr.Id, &pr.Name, &pr.Price)
		products = append(products, &product.Product{Id: pr.Id, Name: pr.Name, Price: pr.Price})
	}
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return &product.GetAllProductsResponse{Products: products}, nil
}
