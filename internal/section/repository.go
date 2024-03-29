package section

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetAllQuery                     = "SELECT * FROM sections;"
	GetQuery                        = "SELECT * FROM sections WHERE id=?;"
	ExistsQuery                     = "SELECT section_number FROM sections WHERE section_number=?;"
	InsertQuery                     = "INSERT INTO sections(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	UpdateQuery                     = "UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, id_product_type=? WHERE id=?;"
	DeleteQuery                     = "DELETE FROM sections WHERE id=?"
	CountProductsByAllSectionsQuery = `SELECT s.id "section_id", s.section_number, COUNT(pb.product_id) "product_count" FROM sections s LEFT JOIN product_batches pb ON s.id = pb.section_id GROUP BY s.id`
	CountProductsBySectionQuery     = `SELECT s.id "section_id", s.section_number, COUNT(pb.product_id) "product_count" FROM sections s LEFT JOIN product_batches pb ON s.id = pb.section_id WHERE s.id=? GROUP BY s.id`
)

type Repository interface {
	GetAll() []domain.Section
	Get(id int) *domain.Section
	Exists(sectionNumber int) bool
	Save(sc domain.Section) int
	Update(s domain.Section)
	Delete(id int)
	CountProductsByAllSections() []domain.ProductsBySectionReport
	CountProductsBySection(id int) *domain.ProductsBySectionReport
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() []domain.Section {
	rows, err := r.db.Query(GetAllQuery)
	if err != nil {
		panic(err)
	}
	sections := make([]domain.Section, 0)

	for rows.Next() {
		s := domain.Section{}
		_ = rows.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
		sections = append(sections, s)
	}
	return sections
}

func (r *repository) Get(id int) *domain.Section {
	row := r.db.QueryRow(GetQuery, id)
	s := domain.Section{}
	err := row.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}
	return &s
}

func (r *repository) Exists(sectionNumber int) bool {
	row := r.db.QueryRow(ExistsQuery, sectionNumber)
	err := row.Scan(&sectionNumber)
	return err == nil
}

func (r *repository) Save(sc domain.Section) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(sc.SectionNumber, sc.CurrentTemperature, sc.MinimumTemperature, sc.CurrentCapacity, sc.MinimumCapacity, sc.MaximumCapacity, sc.WarehouseID, sc.ProductTypeID)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return int(id)
}

func (r *repository) Update(s domain.Section) {
	stmt, err := r.db.Prepare(UpdateQuery)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID, &s.ID)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Delete(id int) {
	stmt, err := r.db.Prepare(DeleteQuery)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}

func (r *repository) CountProductsByAllSections() []domain.ProductsBySectionReport {
	rows, err := r.db.Query(CountProductsByAllSectionsQuery)
	if err != nil {
		panic(err)
	}
	productBatches := make([]domain.ProductsBySectionReport, 0)
	for rows.Next() {
		pb := domain.ProductsBySectionReport{}
		_ = rows.Scan(&pb.SectionID, &pb.SectionNumber, &pb.ProductsCount)
		productBatches = append(productBatches, pb)
	}
	return productBatches
}

func (r *repository) CountProductsBySection(id int) *domain.ProductsBySectionReport {
	rows := r.db.QueryRow(CountProductsBySectionQuery, id)
	pb := domain.ProductsBySectionReport{}
	err := rows.Scan(&pb.SectionID, &pb.SectionNumber, &pb.ProductsCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}
	return &pb
}
