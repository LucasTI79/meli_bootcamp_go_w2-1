package warehouse

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

type Repository interface {
	GetAll() []domain.Warehouse
	Get(id int) *domain.Warehouse
	Exists(warehouseCode string) bool
	Save(w domain.Warehouse) int
	Update(w domain.Warehouse)
	Delete(id int)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() []domain.Warehouse {
	query := "SELECT * FROM warehouses"
	rows, err := r.db.Query(query)
	if err != nil {
		panic(err)
	}

	warehouses := make([]domain.Warehouse, 0)

	for rows.Next() {
		w := domain.Warehouse{}
		_ = rows.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
		warehouses = append(warehouses, w)
	}

	return warehouses
}

func (r *repository) Get(id int) *domain.Warehouse {
	query := "SELECT * FROM warehouses WHERE id=?;"
	row := r.db.QueryRow(query, id)
	w := domain.Warehouse{}
	err := row.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &w
}

func (r *repository) Exists(warehouseCode string) bool {
	query := "SELECT warehouse_code FROM warehouses WHERE warehouse_code=?;"
	row := r.db.QueryRow(query, warehouseCode)
	err := row.Scan(&warehouseCode)
	return err == nil
}

func (r *repository) Save(w domain.Warehouse) int {
	query := "INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(w.Address, w.Telephone, w.WarehouseCode, w.MinimumCapacity, w.MinimumTemperature)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) Update(w domain.Warehouse) {
	query := "UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(&w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature, &w.ID)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Delete(id int) {
	query := "DELETE FROM warehouses WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}
