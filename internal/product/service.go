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
	Create(context.Context, domain.Product) (*domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
func (s *service) Create(ctx context.Context, product domain.Product) (*domain.Product, error) {
	if s.repository.Exists(ctx, product.ProductCode) {
		return nil, apperr.NewResourceAlreadyExists("product with code %s already exists", product.ProductCode)
	}

	id, err := s.repository.Save(ctx, product)

	if err != nil {
		return nil, err
	}

	created, err := s.repository.Get(ctx, id)

	return &created, nil
}
}
