package carrier

import (
	"database/sql"
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

type Repository interface {
	Get(id int) *domain.Carrier
	Save(c domain.Carrier) int
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(id int) *domain.Carrier {
	query := "SELECT * FROM carriers WHERE id=?"
	row := r.db.QueryRow(query, id)
	c := domain.Carrier{}
	err := row.Scan(&c.ID, &c.CID, &c.CompanyName, &c.Address, &c.Telephone, &c.LocalityID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &c
}

func (r *repository) Save(c domain.Carrier) int {
	query := "INSERT INTO carriers(cid,company_name,address,telephone,locality_id) VALUES (?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(c.CID, c.CompanyName)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}
