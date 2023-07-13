package inbound_orders

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
)

const (
	InsertQuery = "INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?, ?, ?, ?, ?)"
	GetQuery    = "SELECT id, order_date, order_number, employee_id, product_batch_id, warehouse_id FROM inbound_orders WHERE id=?"
	ExistsQuery = "SELECT order_number FROM inbound_orders WHERE order_number=?"
)
type Repository interface {
	Get(id int) *domain.InboundOrder
	Save(i domain.InboundOrder) int
	Exists(orderNumber string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(id int) *domain.InboundOrder {
	row := r.db.QueryRow(GetQuery, id)
	i := domain.InboundOrder{}
	var OrderDate string
	err := row.Scan(&i.ID, &OrderDate, &i.OrderNumber, &i.EmployeeId, &i.ProductBatchId, &i.WarehouseId)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}
	
	i.OrderDate = helpers.ToDateTime(OrderDate)

	return &i
}

func (r *repository) Save(i domain.InboundOrder) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(i.OrderDate, i.OrderNumber, i.EmployeeId, i.ProductBatchId, i.WarehouseId)

	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) Exists(orderNumber string) bool {
	row := r.db.QueryRow(ExistsQuery, orderNumber)
	err := row.Scan(&orderNumber)
	return err == nil
}