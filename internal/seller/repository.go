package seller

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

// Repository encapsulates the storage of a Seller.
type Repository interface {
	GetAll(ctx context.Context) []domain.Seller
	Get(ctx context.Context, id int) *domain.Seller
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s domain.Seller) int
	Update(ctx context.Context, s domain.Seller)
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

func (r *repository) GetAll(ctx context.Context) []domain.Seller {
	query := "SELECT * FROM sellers"
	rows, err := r.db.Query(query)
	if err != nil {
		panic(err)
	}

	sellers := make([]domain.Seller, 0)

	for rows.Next() {
		s := domain.Seller{}
		_ = rows.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone)
		sellers = append(sellers, s)
	}

	return sellers
}

func (r *repository) Get(ctx context.Context, id int) *domain.Seller {
	query := "SELECT * FROM sellers WHERE id=?;"
	row := r.db.QueryRow(query, id)
	s := domain.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &s
}

func (r *repository) Exists(ctx context.Context, cid int) bool {
	query := "SELECT cid FROM sellers WHERE cid=?;"
	row := r.db.QueryRow(query, cid)
	err := row.Scan(&cid)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Seller) int {
	query := "INSERT INTO sellers (cid, company_name, address, telephone) VALUES (?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) Update(ctx context.Context, s domain.Seller) {
	query := "UPDATE sellers SET cid=?, company_name=?, address=?, telephone=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.ID)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Delete(ctx context.Context, id int) {
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
