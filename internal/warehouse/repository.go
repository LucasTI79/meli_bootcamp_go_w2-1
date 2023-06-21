package warehouse

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	GetByCode(ctx context.Context, warehouseCode string) (domain.Warehouse, error)
	Exists(ctx context.Context, warehouseCode string) bool
	Save(ctx context.Context, w domain.Warehouse) (int, error)
	Update(ctx context.Context, w domain.Warehouse) error
	Delete(ctx context.Context, id int) error
}

var (
	ErrNotFound = errors.New("warehouse not found")
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetByCode(ctx context.Context, warehouseCode string) (domain.Warehouse, error) {
	query := "SELECT * FROM warehouses WHERE Warehouse_code=?;"
	row := r.db.QueryRow(query, warehouseCode)
	w := domain.Warehouse{}
	err := row.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Warehouse{}, ErrNotFound
		}
		return domain.Warehouse{}, err
	}

	return w, nil
}

func NewWarehouseRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	query := "SELECT * FROM warehouses"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	warehouses := make([]domain.Warehouse, 0)

	for rows.Next() {
		w := domain.Warehouse{}
		_ = rows.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
		warehouses = append(warehouses, w)
	}

	return warehouses, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	query := "SELECT * FROM warehouses WHERE id=?;"
	row := r.db.QueryRow(query, id)
	w := domain.Warehouse{}
	err := row.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Warehouse{}, ErrNotFound
		}
		return domain.Warehouse{}, err
	}

	return w, nil
}

func (r *repository) Exists(ctx context.Context, warehouseCode string) bool {
	query := "SELECT warehouse_code FROM warehouses WHERE warehouse_code=?;"
	row := r.db.QueryRowContext(ctx, query, warehouseCode)
	err := row.Scan(&warehouseCode)
	return err == nil
}

func (r *repository) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	query := "INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, w.Address, w.Telephone, w.WarehouseCode, w.MinimumCapacity, w.MinimumTemperature)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, w domain.Warehouse) error {
	query := "UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature, &w.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM warehouses WHERE id=?"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}
