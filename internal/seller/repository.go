package seller

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetAllQuery = "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers"
	GetQuery    = "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers WHERE id=?"
	ExistsQuery = "SELECT cid FROM sellers WHERE cid=?"
	InsertQuery = "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	UpdateQuery = "UPDATE sellers SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?"
	DeleteQuery = "DELETE FROM sellers WHERE id=?"
)

type Repository interface {
	GetAll() []domain.Seller
	Get(id int) *domain.Seller
	Exists(cid int) bool
	Save(s domain.Seller) int
	Update(s domain.Seller)
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

func (r *repository) GetAll() []domain.Seller {
	rows, err := r.db.Query(GetAllQuery)
	if err != nil {
		panic(err)
	}

	sellers := make([]domain.Seller, 0)

	for rows.Next() {
		s := domain.Seller{}
		_ = rows.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
		sellers = append(sellers, s)
	}

	return sellers
}

func (r *repository) Get(id int) *domain.Seller {
	row := r.db.QueryRow(GetQuery, id)
	s := domain.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &s
}

func (r *repository) Exists(cid int) bool {
	row := r.db.QueryRow(ExistsQuery, cid)
	err := row.Scan(&cid)
	return err == nil
}

func (r *repository) Save(s domain.Seller) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) Update(s domain.Seller) {
	stmt, err := r.db.Prepare(UpdateQuery)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID, s.ID)
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
