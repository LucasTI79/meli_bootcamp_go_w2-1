package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll(ctx context.Context) []domain.Section {
	args := r.Called()
	return args.Get(0).([]domain.Section)
}

func (r *Repository) Get(ctx context.Context, id int) *domain.Section {
	args := r.Called(id)
	return args.Get(0).(*domain.Section)
}

func (r *Repository) Exists(ctx context.Context, sectionNumber int) bool {
	args := r.Called(sectionNumber)
	return args.Get(0).(bool)
}

func (r *Repository) Save(ctx context.Context, section domain.Section) int {
	args := r.Called(section)
	return args.Get(0).(int)
}

func (r *Repository) Update(ctx context.Context, section domain.Section) {
	r.Called(section)
}

func (r *Repository) Delete(ctx context.Context, id int) {
	r.Called(id)
}