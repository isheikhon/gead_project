package database

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // <------------ here
	"log"
	"transaction/server/proto/product"
)

var db, err = sql.Open("postgres", "postgresql://root@localhost:26258/defaultdb?sslmode=disable")

func DBGetProductByID(req *product.ProductQueryRequest) (*product.ProductResponse, error) {
	pr := product.Product{}
	query := "SELECT * FROM product where id = $1"
	err := db.QueryRow(query, req.ID).Scan(&pr.ID, &pr.Name, &pr.Price)
	if err != nil {
		panic(err)
	}
	return &product.ProductResponse{Product: &product.Product{ID: pr.ID, Name: pr.Name, Price: pr.Price}}, nil
}

func DBAddProduct(req *product.Product) (*product.ProductResponse, error) {
	log.Printf("Inserting a Row in to DB")
	var UUID = uuid.New().String()
	//Inserting a Row in to DB.
	sqlStatement := `INSERT INTO product (id,name,price) VALUES ($1, $2,$3)`
	_, err = db.Exec(sqlStatement, UUID, req.Name, req.Price)
	if err != nil {
		panic(err)
	}
	return &product.ProductResponse{Product: &product.Product{ID: UUID, Name: req.Name, Price: req.Price}}, nil
}

func DBGetAllProducts(req *product.AllProductQueryRequest) (*product.ProductListResponse, error) {
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
		rows.Scan(&pr.ID, &pr.Name, &pr.Price)
		products = append(products, &product.Product{ID: pr.ID, Name: pr.Name, Price: pr.Price})
	}
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return &product.ProductListResponse{Products: products}, nil
}
