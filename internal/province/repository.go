package province

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetQuery = "SELECT id FROM provinces WHERE id=?"
)

// Repository encapsulates the storage of a Province.
type Repository interface {
	Get(id int) *domain.Province
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(id int) *domain.Province {
	row := r.db.QueryRow(GetQuery, id)
	s := domain.Province{}
	err := row.Scan(&s.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &s
}
