package database

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // <------------ here
	"log"
	"transaction/server/proto/customer"
)

var db, err = sql.Open("postgres", "postgresql://root@localhost:26258/defaultdb?sslmode=disable")

func DBGetCustomerByID(req *customer.CustomerQueryRequest) (*customer.CustomerResponse, error) {
	cr := customer.Customer{}
	query := "SELECT * FROM customer where id = $1"
	err := db.QueryRow(query, req.ID).Scan(&cr.ID, &cr.Name)
	if err != nil {
		panic(err)
	}
	return &customer.CustomerResponse{Customer: &customer.Customer{ID: cr.ID, Name: cr.Name}}, nil
}

func DBAddCustomer(req *customer.Customer) (*customer.CustomerResponse, error) {
	log.Printf("Inserting a Row in to DB")
	var UUID = uuid.New().String()
	//Inserting a Row in to DB.
	sqlStatement := `INSERT INTO customer (id,name) VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, UUID, req.Name)
	if err != nil {
		panic(err)
	}
	return &customer.CustomerResponse{Customer: &customer.Customer{ID: UUID, Name: req.Name}}, nil
}

func DBGetAllCustomer(req *customer.AllCustomersQueryRequest) (*customer.CustomerListResponse, error) {
	var customers []*customer.Customer
	query := "SELECT * FROM customer"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		cr := customer.Customer{}
		rows.Scan(&cr.ID, &cr.Name)
		customers = append(customers, &customer.Customer{ID: cr.ID, Name: cr.Name})
	}
	log.Printf("Received request: %v", req.ProtoReflect().Descriptor().FullName())
	return &customer.CustomerListResponse{Customers: customers}, nil
}
