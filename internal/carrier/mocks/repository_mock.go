package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Exists(cid string) bool {
	args := r.Called(cid)
	return args.Get(0).(bool)
}

func (r *Repository) Save(carrier domain.Carrier) int {
	args := r.Called(carrier)
	return args.Get(0).(int)
}

func (r *Repository) Get(id int) *domain.Carrier {
	args := r.Called(id)
	return args.Get(0).(*domain.Carrier)
}