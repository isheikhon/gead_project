package database

import (
	"database/sql"
	_ "github.com/lib/pq" // <------------ here
	"log"
	"transaction/server/proto/analytics"
)

var db, err = sql.Open("postgres", "postgresql://root@localhost:26258/defaultdb?sslmode=disable")

func DBGetTotalSales() (*analytics.GetTotalSalesResponse, error) {
	ts := analytics.TotalSales{}
	query := "SELECT SUM(totalprice) FROM transaction"
	err := db.QueryRow(query).Scan(&ts.Sales)
	if err != nil {
		panic(err)
	}
	return &analytics.GetTotalSalesResponse{Sales: &analytics.TotalSales{Sales: ts.Sales}}, nil
}

func DBGetSalesByProduct() (*analytics.GetSalesByProductResponse, error) {
	log.Printf("Inserting a Row in to DB")
	var products []*analytics.Product
	//Inserting a Row in to DB.
	query := "SELECT productid,name,SUM(totalprice) AS Sales FROM transaction GROUP BY productid, name"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		pr := analytics.Product{}
		rows.Scan(&pr.Id, &pr.Name, &pr.Sales)
		products = append(products, &analytics.Product{Id: pr.Id, Name: pr.Name, Sales: pr.Sales})
	}
	if err != nil {
		panic(err)
	}
	return &analytics.GetSalesByProductResponse{Products: products}, nil
}

func DBGetTopCustomers() (*analytics.GetTopCustomersResponse, error) {
	log.Printf("Inserting a Row in to DB")
	var customers []*analytics.Customer
	//Inserting a Row in to DB.
	query := "SELECT transaction.customerid,SUM(transaction.totalprice), customer.name FROM transaction,customer GROUP BY transaction.customerid,customer.name LIMIT 5 |DESC;"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		cr := analytics.Customer{}
		rows.Scan(&cr.Id, &cr.Sales, &cr.Name)
		customers = append(customers, &analytics.Customer{Id: cr.Id, Name: cr.Name, Sales: cr.Sales})
	}
	if err != nil {
		panic(err)
	}
	return &analytics.GetTopCustomersResponse{Customers: customers}, nil
}
