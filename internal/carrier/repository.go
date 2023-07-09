package carrier

import (
	"database/sql"
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetQuery    = "SELECT * FROM carriers WHERE id=?"
	ExistsQuery = "SELECT cid FROM carriers WHERE cid=?"
	InsertQuery = "INSERT INTO carriers(cid,company_name,address,telephone,locality_id) VALUES (?,?,?,?,?)"
)

type Repository interface {
	Get(id int) *domain.Carrier
	Save(c domain.Carrier) int
	Exists(cid string) bool
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
	row := r.db.QueryRow(GetQuery, id)
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
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(c.CID, c.CompanyName, c.Address, c.Telephone, c.LocalityID)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) Exists(cid string) bool {
	row := r.db.QueryRow(ExistsQuery, cid)
	err := row.Scan(&cid)
	return err == nil
}
