package order_status

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetQuery = "SELECT id FROM order_status WHERE id=?"
)

type Repository interface {
	Get(id int) *domain.OrderStatus
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(id int) *domain.OrderStatus {
	row := r.db.QueryRow(GetQuery, id)
	order := domain.OrderStatus{}
	err := row.Scan(&order.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &order
}
