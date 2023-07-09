package section

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	SectionExists = "SELECT id FROM sections WHERE id=?"
)

type Repository interface {
	GetAll(ctx context.Context) []domain.Section
	Get(ctx context.Context, id int) *domain.Section
	Exists(ctx context.Context, sectionNumber int) bool
	Save(ctx context.Context, sc domain.Section) int
	Update(ctx context.Context, s domain.Section)
	Delete(ctx context.Context, id int)
	ExistSectionID(sectionID int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) []domain.Section {
	query := "SELECT * FROM sections;"
	rows, err := r.db.Query(query)
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

func (r *repository) Get(ctx context.Context, id int) *domain.Section {
	query := "SELECT * FROM sections WHERE id=?;"
	row := r.db.QueryRow(query, id)
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

func (r *repository) Exists(ctx context.Context, sectionNumber int) bool {
	query := "SELECT section_number FROM sections WHERE section_number=?;"
	row := r.db.QueryRow(query, sectionNumber)
	err := row.Scan(&sectionNumber)
	return err == nil
}

func (r *repository) Save(ctx context.Context, sc domain.Section) int {
	query := "INSERT INTO sections(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	stmt, err := r.db.Prepare(query)
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

func (r *repository) Update(ctx context.Context, s domain.Section) {
	query := "UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, id_product_type=? WHERE id=?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID, &s.ID)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Delete(ctx context.Context, id int) {
	query := "DELETE FROM sections WHERE id=?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}

func (r *repository) ExistSectionID(sectionID int) bool {
	row := r.db.QueryRow(SectionExists, sectionID)
	err := row.Scan(&sectionID)
	return err == nil
}
