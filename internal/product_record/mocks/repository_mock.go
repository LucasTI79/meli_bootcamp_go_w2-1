package mocks

import (
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Get(id int) *domain.ProductRecord {
	args := r.Called(id)
	return args.Get(0).(*domain.ProductRecord)
}
func (r *Repository) Exists(productId int, lastUpdateDate time.Time) bool {
	args := r.Called(productId, lastUpdateDate)
	return args.Get(0).(bool)
}
func (r *Repository) Save(locality domain.ProductRecord) int {
	args := r.Called(locality)
	return args.Get(0).(int)
}
