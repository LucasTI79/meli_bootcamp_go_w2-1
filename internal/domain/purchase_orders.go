package domain

import "time"

type PurchaseOrders struct {
	ID				int  		`json:"id"`
	OrderNumber 	string 		`json:"order_number"`
	OrderDate  		time.Time   `json:"order_date"`
	TrackingCode   	string   	`json:"tracking_code"`
	BuyerID   		int    		`json:"buyer_id"`
	CarrierID   	int    		`json:"carrier_id"`
	ProductRecordID int    		`json:"product_record_id"`
	OrderStatusID  	int    		`json:"order_status_id"`
	WarehouseID   	int    		`json:"warehouse_id"`
}

