package purchase_orders

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetQuery    = "SELECT id, order_number, order_date, tracking_code, buyer_id, carrier_id, product_record_id, order_status_id, warehouse_id FROM purchase_orders WHERE id=?"
	ExistsQuery = "SELECT order_number FROM purchase_orders WHERE order_number=?"
	InsertQuery = "INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, carrier_id, product_record_id, order_status_id, warehouse_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
)

type Repository interface {
	Get(id int) *domain.PurchaseOrders
	Exists(orderNumber string) bool
	Save(purchaseOrder domain.PurchaseOrders) int
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(id int) *domain.PurchaseOrders {
	row := r.db.QueryRow(GetQuery, id)
	po := domain.PurchaseOrders{}
	err := row.Scan(&po.ID, &po.OrderNumber, &po.OrderDate, &po.TrackingCode, &po.BuyerID, &po.CarrierID, &po.ProductRecordID, &po.OrderStatusID, &po.WarehouseID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &po
}

func (r *repository) Exists(orderNumber string) bool {
	row := r.db.QueryRow(ExistsQuery, orderNumber)
	err := row.Scan(&orderNumber)
	return err == nil
}

func (r *repository) Save(po domain.PurchaseOrders) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(po.OrderNumber, po.OrderDate, po.TrackingCode, po.BuyerID, po.CarrierID, po.ProductRecordID, po.OrderStatusID, po.WarehouseID)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

