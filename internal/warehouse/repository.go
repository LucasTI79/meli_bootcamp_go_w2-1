package warehouse

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetAllQuery = "SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id FROM warehouses"
	GetQuery    = "SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id FROM warehouses WHERE id=?"
	ExistsQuery = "SELECT warehouse_code FROM warehouses WHERE warehouse_code=?"
	InsertQuery = "INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id) VALUES (?, ?, ?, ?, ?,?)"
	UpdateQuery = "UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=?, locality_id=? WHERE id=?"
	DeleteQuery = "DELETE FROM warehouses WHERE id=?"
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
	rows, err := r.db.Query(GetAllQuery)
	if err != nil {
		panic(err)
	}

	warehouses := make([]domain.Warehouse, 0)

	for rows.Next() {
		w := domain.Warehouse{}
		_ = rows.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature, &w.LocalityID)
		warehouses = append(warehouses, w)
	}

	return warehouses
}

func (r *repository) Get(id int) *domain.Warehouse {
	row := r.db.QueryRow(GetQuery, id)
	w := domain.Warehouse{}
	err := row.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature, &w.LocalityID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &w
}

func (r *repository) Exists(warehouseCode string) bool {
	row := r.db.QueryRow(ExistsQuery, warehouseCode)
	err := row.Scan(&warehouseCode)
	return err == nil
}

func (r *repository) Save(w domain.Warehouse) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(w.Address, w.Telephone, w.WarehouseCode, w.MinimumCapacity, w.MinimumTemperature, w.LocalityID)
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
	stmt, err := r.db.Prepare(UpdateQuery)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(&w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature, &w.LocalityID, &w.ID)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Delete(id int) {
	stmt, err := r.db.Prepare(DeleteQuery)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}
