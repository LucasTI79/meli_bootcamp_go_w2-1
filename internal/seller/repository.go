package seller

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

// Repository encapsulates the storage of a Seller.
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
	query := "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers"
	rows, err := r.db.Query(query)
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
	query := "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers WHERE id=?;"
	row := r.db.QueryRow(query, id)
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
	query := "SELECT cid FROM sellers WHERE cid=?;"
	row := r.db.QueryRow(query, cid)
	err := row.Scan(&cid)
	return err == nil
}

func (r *repository) Save(s domain.Seller) int {
	query := "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
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
	query := "UPDATE sellers SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID, s.ID)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Delete(id int) {
	query := "DELETE FROM sellers WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}
