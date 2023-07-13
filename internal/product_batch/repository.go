package product_batch

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

var (
	InsertQuery = "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	ExistsQuery = "SELECT id FROM product_batches WHERE batch_number = ?"
	GetQuery    = "SELECT id, batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id FROM product_batches WHERE id = ?"
)

type Repository interface {
	Exists(batchNumber int) bool
	Save(pb domain.ProductBatch) int
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

func (r *repository) Exists(batchNumber int) bool {
	row := r.db.QueryRow(ExistsQuery, batchNumber)
	err := row.Scan(&batchNumber)
	return err == nil
}

func (r *repository) Save(pb domain.ProductBatch) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(pb.BatchNumber, pb.CurrentQuantity, pb.CurrentTemperature, pb.DueDate, pb.InitialQuantity, pb.ManufacturingDate, pb.ManufacturingHour, pb.MinimumTemperature, pb.ProductID, pb.SectionID)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return int(id)
}
func (r *repository) Get(id int) *domain.ProductBatch {
	row := r.db.QueryRow(GetQuery, id)
	var pb domain.ProductBatch
	err := row.Scan(&pb.ID, &pb.BatchNumber, &pb.CurrentQuantity, &pb.CurrentTemperature, &pb.DueDate, &pb.InitialQuantity, &pb.ManufacturingDate, &pb.ManufacturingHour, &pb.MinimumTemperature, &pb.ProductID, &pb.SectionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}
	return &pb
}
