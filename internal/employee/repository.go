package employee

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) []domain.Employee
	Get(ctx context.Context, id int) *domain.Employee
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, p domain.Employee) int
	Update(ctx context.Context, p domain.Employee)
	Delete(ctx context.Context, id int)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) []domain.Employee{
	query := "SELECT * FROM employees;"
	rows, err := r.db.Query(query)
	if err != nil {
		panic(err)
	}

	employees := make([]domain.Employee, 0)

	for rows.Next() {
		e := domain.Employee{}
		_ = rows.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
		employees = append(employees, e)
	}

	return employees
}

func (r *repository) Get(ctx context.Context, id int) *domain.Employee {
	query := "SELECT * FROM employees WHERE id=?;"
	row := r.db.QueryRow(query, id)
	e := domain.Employee{}
	err := row.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &e
}

func (r *repository) Exists(ctx context.Context, cardNumberID string) bool {
	query := "SELECT card_number_id FROM employees WHERE card_number_id=?;"
	row := r.db.QueryRow(query, cardNumberID)
	err := row.Scan(&cardNumberID)
	
	return err == nil
}

func (r *repository) Save(ctx context.Context, e domain.Employee) int {
	query := "INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(e.CardNumberID, e.FirstName, e.LastName, e.WarehouseID)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) Update(ctx context.Context, e domain.Employee) {
	query := "UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(&e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID, &e.ID)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Delete(ctx context.Context, id int){
	query := "DELETE FROM employees WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}