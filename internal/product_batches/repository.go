package product_batches

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	CreateQuery                       = "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	GetQuery                          = "SELECT id, batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id FROM product_batches WHERE id=?"
	ExistsQuery                       = "SELECT batch_number FROM product_batches WHERE batch_number=?"
	InsertQuery                       = "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	CountProductBatchesBySectionQuery = "SELECT pb.section_id, s.section_number, COUNT(s.id) 'products_count' FROM product_batches pb JOIN products p ON p.id = pb.product_id JOIN sections s ON s.id = pb.section_id GROUP BY s.id"
)

// Repository encapsulates the storage of a product_batches.
type Repository interface {
	Save(pb domain.ProductBatches) int
	Get(id int) *domain.ProductBatches
	Exists(BatchNumber string) bool
	CheckSectionExists(id int) bool
	CheckProductExists(id int) bool
	CountProductBatchesBySection() []domain.CountProductBatchesBySection
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(pb domain.ProductBatches) int {

	if !r.sectionExists(pb.SectionID) {
		panic("section does not exist")
	}

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

func (r *repository) Exists(BatchNumber string) bool {
	row := r.db.QueryRow(ExistsQuery, BatchNumber)
	err := row.Scan(&BatchNumber)
	return err == nil
}

func (r *repository) sectionExists(id int) bool {

	query := `SELECT id FROM sections WHERE id=?`
	var count int
	err := r.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func (r *repository) CheckSectionExists(id int) bool {
	query := `SELECT id FROM sections WHERE id=?`
	var count int
	err := r.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func (r *repository) CheckProductExists(id int) bool {
	query := `SELECT id FROM products WHERE id=?`
	var count int
	err := r.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
func (r *repository) CountProductBatchesBySection() []domain.CountProductBatchesBySection {
	rows, err := r.db.Query(CountProductBatchesBySectionQuery)
	if err != nil {
		panic(err)
	}
	productsBatcherBySections := make([]domain.CountProductBatchesBySection, 0)

	for rows.Next() {
		s := domain.CountProductBatchesBySection{}
		_ = rows.Scan(&s.SectionID, &s.SectionNumber, &s.ProductsCount)
		productsBatcherBySections = append(productsBatcherBySections, s)
	}
	return productsBatcherBySections
}
