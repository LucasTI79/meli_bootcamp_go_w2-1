package product_batches

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

var (
	CreateQuery = "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	ExistsQuery = "SELECT id FROM product_batches WHERE batch_number = ?"
	SaveQuery   = "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	GetQuery    = "SELECT pb.section_id, s.section_number, COUNT(s.id) 'products_count' FROM product_batches pb JOIN products p ON p.id = pb.product_id JOIN sections s ON s.id = pb.section_id GROUP BY s.id"
	GetAllByID  = "SELECT pb.section_id, s.section_number, COUNT(s.id) 'products_count' FROM product_batches pb JOIN products p ON p.id = pb.product_id JOIN sections s ON s.id = pb.section_id WHERE s.id =? GROUP BY s.id"
	GetByID     = "SELECT * FROM product_batches WHERE id = ?"
)

type Repository interface {
	Exists(batchNumber int) bool
	Save(pb domain.ProductBatches) int
	CountProductsByAllSections() ([]domain.ProductsBySectionReport, error)
	CountProductsBySection(id int) ([]domain.ProductsBySectionReport, error)
	Get(id int) *domain.ProductBatches
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

func (r *repository) Save(pb domain.ProductBatches) int {
	stmt, err := r.db.Prepare(SaveQuery)
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

func (r *repository) CountProductsByAllSections() ([]domain.ProductsBySectionReport, error) {
	rows, err := r.db.Query(GetQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var productBatches []domain.ProductsBySectionReport
	for rows.Next() {
		var pb domain.ProductsBySectionReport
		err := rows.Scan(&pb.SectionID, &pb.SectionNumber, &pb.ProductsCount)
		if err != nil {
			panic(err)
		}
		productBatches = append(productBatches, pb)
	}
	return productBatches, nil
}

func (r *repository) CountProductsBySection(id int) ([]domain.ProductsBySectionReport, error) {
	rows, err := r.db.Query(GetAllByID, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var productBatches []domain.ProductsBySectionReport
	for rows.Next() {
		var pb domain.ProductsBySectionReport
		err := rows.Scan(&pb.SectionID, &pb.SectionNumber, &pb.ProductsCount)
		if err != nil {
			panic(err)
		}
		productBatches = append(productBatches, pb)
	}
	return productBatches, nil
}

func (r *repository) Get(id int) *domain.ProductBatches {
	row := r.db.QueryRow("SELECT * FROM product_batches WHERE id = ?", id)
	var pb domain.ProductBatches
	err := row.Scan(&pb.ID, &pb.BatchNumber, &pb.CurrentQuantity, &pb.CurrentTemperature, &pb.DueDate, &pb.InitialQuantity, &pb.ManufacturingDate, &pb.ManufacturingHour, &pb.MinimumTemperature, &pb.ProductID, &pb.SectionID)
	if err != nil {
		panic(err)
	}
	return &pb
}
