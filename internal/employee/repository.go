package employee

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetAllQuery = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees;"
	GetQuery = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id=?;"
	ExistsQuery = "SELECT card_number_id FROM employees WHERE card_number_id=?;"
	SaveQuery = "INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)"
	UpdateQuery = "UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?"
	DeleteQuery = "DELETE FROM employees WHERE id=?"

	CountInboundOrdersByAllEmployeesQuery = `SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, count(i.id) "inbound_orders_count"
	FROM employees e
	JOIN inbound_orders i ON e.id = i.employee_id
	GROUP BY e.id`
	CountInboundOrdersByEmployeeQuery = `SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, count(i.id) "inbound_orders_count"
	FROM employees e
	JOIN inbound_orders i ON e.id = i.employee_id
	WHERE e.id=?
	GROUP BY e.id`
)

type Repository interface {
	GetAll() []domain.Employee
	Get(id int) *domain.Employee
	Exists(cardNumberID string) bool
	Save(p domain.Employee) int
	Update(p domain.Employee)
	Delete(id int)
	CountInboundOrdersByAllEmployees() []domain.InboundOrdersByEmployee
	CountInboundOrdersByEmployee(id int) *domain.InboundOrdersByEmployee
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() []domain.Employee{
	rows, err := r.db.Query(GetAllQuery)
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

func (r *repository) Get(id int) *domain.Employee {
	row := r.db.QueryRow(GetQuery, id)
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

func (r *repository) Exists(cardNumberID string) bool {
	row := r.db.QueryRow(ExistsQuery, cardNumberID)
	err := row.Scan(&cardNumberID)
	
	return err == nil
}

func (r *repository) Save(e domain.Employee) int {
	stmt, err := r.db.Prepare(SaveQuery)
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

func (r *repository) Update(e domain.Employee) {
	stmt, err := r.db.Prepare(UpdateQuery)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(&e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID, &e.ID)
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

func (r *repository) CountInboundOrdersByAllEmployees() []domain.InboundOrdersByEmployee {
	rows, err := r.db.Query(CountInboundOrdersByAllEmployeesQuery)
	if err != nil {
		panic(err)
	}

	inboundOrdersByEmployees := make([]domain.InboundOrdersByEmployee, 0)

	for rows.Next() {
		e := domain.InboundOrdersByEmployee{}
		_ = rows.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID, &e.InboundOrdersCount)
		inboundOrdersByEmployees = append(inboundOrdersByEmployees, e)
	}

	return inboundOrdersByEmployees
}

func (r *repository) CountInboundOrdersByEmployee(id int) *domain.InboundOrdersByEmployee {
	rows := r.db.QueryRow(CountInboundOrdersByEmployeeQuery, id)
	e := domain.InboundOrdersByEmployee{}
	err := rows.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID, &e.InboundOrdersCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &e
}