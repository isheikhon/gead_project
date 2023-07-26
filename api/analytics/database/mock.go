package database

import (
	"context"
	"transaction/server/proto/analytics"
)

type MockAnalyticsData struct{}

func (m *MockAnalyticsData) GetTotalSales(ctx context.Context) (*int32, error) {
	var a int32 = 999
	return &a, nil
}

//func (m *MockAnalyticsData) GetSalesByProduct(ctx context.Context, product_id int32) (*float64, error) {
//	var a float64 = 998
//	return &a, nil
//}

func (m *MockAnalyticsData) GetTop5Customers(ctx context.Context) ([]analytics.Customer, error) {
	customers := []analytics.Customer{
		{Id: "1", Name: "Carlos", Sales: 950},
		{Id: "2", Name: "Stephan", Sales: 450},
		{Id: "3", Name: "Hannah", Sales: 900},
		{Id: "4", Name: "Suzan", Sales: 260},
		{Id: "5", Name: "Sean", Sales: 1200},
	}
	return customers, nil
}
