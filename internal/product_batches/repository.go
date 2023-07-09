package product_batches

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetQuery                              = "SELECT id, batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id FROM product_batches WHERE id=?"
	ExistsSectionQuery                    = "SELECT id FROM sections WHERE id=?"
	InsertQuery                           = "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	CountProductBatchesBySectionQuery     = "SELECT pb.section_id, s.section_number, COUNT(s.id) 'products_count' FROM product_batches pb JOIN products p ON p.id = pb.product_id JOIN sections s ON s.id = pb.section_id GROUP BY s.id"
	ProductBatchesBySectionProductsReport = "SELECT count(pb.id) as `products_count`, pb.section_id, s.section_number FROM product_batches pb JOIN sections s ON pb.section_id = s.id WHERE pb.section_id = ? GROUP BY pb.section_id"
)

type Repository interface {
	Save(pb domain.ProductBatches) int
	ProductBatchesBySectionProductsReport() []domain.ProductsBySectionReport
	CountProductBatchesBySection() domain.ProductsBySectionReport
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

func (r *repository) ProductBatchesBySectionProductsReport() []domain.ProductsBySectionReport {
	rows, err := r.db.Query(ProductBatchesBySectionProductsReport)
	if err != nil {
		panic(err)
	}
	var productsBySection []domain.ProductsBySectionReport
	for rows.Next() {
		var pbBySection domain.ProductsBySectionReport
		err = rows.Scan(&pbBySection.SectionID, &pbBySection.ProductsCount, &pbBySection.SectionID)
		if err != nil {
			panic(err)
		}
		productsBySection = append(productsBySection, pbBySection)
	}
	return productsBySection
}

func (r *repository) CountProductBatchesBySection() domain.ProductsBySectionReport {
	rows, err := r.db.Query(CountProductBatchesBySectionQuery)
	if err != nil {
		panic(err)
	}
	var productsBySection domain.ProductsBySectionReport
	for rows.Next() {
		err := rows.Scan(&productsBySection.SectionID, &productsBySection.ProductsCount, &productsBySection.SectionID)
		if err != nil {
			panic(err)
		}
	}
	return productsBySection
}
