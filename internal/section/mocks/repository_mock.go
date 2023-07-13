package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll() []domain.Section {
	args := r.Called()
	return args.Get(0).([]domain.Section)
}

func (r *Repository) Get(id int) *domain.Section {
	args := r.Called(id)
	return args.Get(0).(*domain.Section)
}

func (r *Repository) Exists(sectionNumber int) bool {
	args := r.Called(sectionNumber)
	return args.Get(0).(bool)
}

func (r *Repository) Save(section domain.Section) int {
	args := r.Called(section)
	return args.Get(0).(int)
}

func (r *Repository) Update(section domain.Section) {
	r.Called(section)
}

func (r *Repository) Delete(id int) {
	r.Called(id)
}
func (r *Repository) CountProductsByAllSections() []domain.ProductsBySectionReport {
	args := r.Called()
	return args.Get(0).([]domain.ProductsBySectionReport)
}
func (r *Repository) CountProductsBySection(id int) *domain.ProductsBySectionReport {
	args := r.Called(id)
	return args.Get(0).(*domain.ProductsBySectionReport)
}
