package productbatches

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

// Repository encapsulates the storage of a product_batches.
type Repository interface {
	Create(pb domain.ProductBatches) int
	Get(id int) *domain.ProductBatches
	Exists(name string) bool
	Update(id int, pb domain.ProductBatches) *domain.ProductBatches
	Delete(id int) *domain.ProductBatches
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(pb domain.ProductBatches) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(pb.BatchNumber)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) Get(id int) *domain.ProductBatches {
	row := r.db.QueryRow(GetQuery, id)
	pb := domain.ProductBatches{}
	err := row.Scan(&pb.ID, &pb.BatchNumber, &pb.CurrentQuantity, &pb.CurrentTemperature, &pb.DueDate, &pb.InitialQuantity, &pb.ManufacturingDate, &pb.ManufacturingHour, &pb.MinimumTemperature, &pb.ProductID, &pb.SectionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &pb
}

func (r *repository) Exists(name string) bool {
	row := r.db.QueryRow(ExistsQuery, name)
	err := row.Scan(&name)
	return err == nil
}

func (r *repository) Update(id int, pb domain.ProductBatches) *domain.ProductBatches {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(pb.BatchNumber)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return &pb
}

func (r *repository) Delete(id int) *domain.ProductBatches {
	row := r.db.QueryRow(GetQuery, id)
	pb := domain.ProductBatches{}
	err := row.Scan(&pb.ID, &pb.BatchNumber, &pb.CurrentQuantity, &pb.CurrentTemperature, &pb.DueDate, &pb.InitialQuantity, &pb.ManufacturingDate, &pb.ManufacturingHour, &pb.MinimumTemperature, &pb.ProductID, &pb.SectionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}
	return &pb
}
