package domain

import "time"

type InboundOrder struct {
	ID int `json:"id"`
	OrderDate time.Time `json:"order_date"`
	OrderNumber int `json:"order_number"`
	EmployeeId int `json:"employee_id"`
	ProductBatchId int `json:"product_batch_id"`
	WarehouseId int `json:"warehouse_id"`
}