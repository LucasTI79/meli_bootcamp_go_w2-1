package product

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("product not found")
)

type Service interface {
	GetAll(context.Context) ([]domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repository.GetAll(ctx)
	return products, err
}
