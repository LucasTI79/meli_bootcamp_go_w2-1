package seller

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	return s.repository.GetAll(ctx)
}

func NewService(r Repository) Service {
	return &service{repository: r}
}
