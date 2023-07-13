package product_type

import (
	"database/sql"
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

type Repository interface {
	Get(id int) *domain.ProductType
}

const GetQuery = "SELECT * FROM product_type WHERE id=?"

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Get(id int) *domain.ProductType {
	row := r.db.QueryRow(GetQuery, id)
	pt := domain.ProductType{}
	err := row.Scan(&pt.ID, &pt.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &pt
}
