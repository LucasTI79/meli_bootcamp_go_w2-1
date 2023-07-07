package locality

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetQuery    = "SELECT id, locality_name, province_id FROM localities WHERE id=?"
	ExistsQuery = "SELECT locality_name FROM localities WHERE locality_name=?"
	InsertQuery = "INSERT INTO localities (locality_name, province_id) VALUES (?, ?)"

	CountSellersByAllLocalitiesQuery = `SELECT s.locality_id, l.locality_name, count(s.id) "sellers_count"
		FROM localities l
		JOIN sellers s ON l.id = s.locality_id
		GROUP BY l.id`

	CountSellersByLocalityQuery = `SELECT s.locality_id, l.locality_name, count(s.id) "sellers_count"
		FROM localities l
		JOIN sellers s ON l.id = s.locality_id
		WHERE l.id=?
		GROUP BY l.id`
)

// Repository encapsulates the storage of a Locality.
type Repository interface {
	Get(id int) *domain.Locality
	Exists(localityName string) bool
	Save(l domain.Locality) int
	CountSellersByAllLocalities() []domain.SellersByLocalityReport
	CountSellersByLocality(id int) *domain.SellersByLocalityReport
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(id int) *domain.Locality {
	row := r.db.QueryRow(GetQuery, id)
	l := domain.Locality{}
	err := row.Scan(&l.ID, &l.LocalityName, &l.ProvinceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &l
}

func (r *repository) Exists(name string) bool {
	row := r.db.QueryRow(ExistsQuery, name)
	err := row.Scan(&name)
	return err == nil
}

func (r *repository) Save(l domain.Locality) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(l.LocalityName, l.ProvinceID)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) CountSellersByAllLocalities() []domain.SellersByLocalityReport {
	rows, err := r.db.Query(CountSellersByAllLocalitiesQuery)
	if err != nil {
		panic(err)
	}

	sellersByLocalities := make([]domain.SellersByLocalityReport, 0)

	for rows.Next() {
		s := domain.SellersByLocalityReport{}
		_ = rows.Scan(&s.ID, &s.LocalityName, &s.SellersCount)
		sellersByLocalities = append(sellersByLocalities, s)
	}

	return sellersByLocalities
}

func (r *repository) CountSellersByLocality(id int) *domain.SellersByLocalityReport {
	rows := r.db.QueryRow(CountSellersByLocalityQuery, id)
	s := domain.SellersByLocalityReport{}
	err := rows.Scan(&s.ID, &s.LocalityName, &s.SellersCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &s
}
