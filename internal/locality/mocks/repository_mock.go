package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Get(id int) *domain.Locality {
	args := r.Called(id)
	return args.Get(0).(*domain.Locality)
}
func (r *Repository) Exists(localityName string) bool {
	args := r.Called(localityName)
	return args.Get(0).(bool)
}
func (r *Repository) Save(locality domain.Locality) int {
	args := r.Called(locality)
	return args.Get(0).(int)
}
func (r *Repository) CountSellersByAllLocalities() []domain.SellersByLocalityReport {
	args := r.Called()
	return args.Get(0).([]domain.SellersByLocalityReport)
}
func (r *Repository) CountSellersByLocality(id int) *domain.SellersByLocalityReport {
	args := r.Called(id)
	return args.Get(0).(*domain.SellersByLocalityReport)
}
