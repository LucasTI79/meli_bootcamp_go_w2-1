package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Get(id int) *domain.ProductType {
	args := r.Called(id)
	return args.Get(0).(*domain.ProductType)
}
