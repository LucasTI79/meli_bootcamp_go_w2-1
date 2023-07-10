package product_batch

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	getQuery = "SELECT id FROM product_batches WHERE id =?"
)

type Repository interface {
	Get(id int) *domain.ProductBatch
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(id int) *domain.ProductBatch {
	row := r.db.QueryRow(getQuery, id)
	p := domain.ProductBatch{}
	err := row.Scan(&p.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &p
}