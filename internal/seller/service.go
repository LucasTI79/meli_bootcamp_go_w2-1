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
	Get(ctx context.Context, id int) (domain.Seller, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Get(ctx context.Context, id int) (domain.Seller, error) {
	seller, err := s.repository.Get(ctx, id)
	if err != nil {
		return seller, ErrNotFound
	}
	return seller, nil
}
func NewService(r Repository) Service {
	return &service{repository: r}
}
