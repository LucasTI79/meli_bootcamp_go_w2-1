package product_type

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetQuery = "SELECT id FROM product_types WHERE id=?"
)

type Repository interface {
	Get(id int) *domain.ProductType
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(id int) *domain.ProductType {
	row := r.db.QueryRow(GetQuery, id)
	productType := domain.ProductType{}

	err := row.Scan(&productType.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &productType
}
