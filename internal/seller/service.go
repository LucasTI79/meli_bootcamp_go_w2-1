package seller

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

// Errors
var (
	ErrNotFound      = errors.New("seller not found")
	ErrAlreadyExists = errors.New("already exists")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Save(ctx context.Context, seller domain.CreateSeller) (int, error)
	Update(ctx context.Context, seller domain.UpdateSeller) (domain.UpdateSeller, error)
	Delete(ctc context.Context, id int) error
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
func (s *service) Save(ctx context.Context, seller domain.CreateSeller) (int, error) {
	if s.repository.Exists(ctx, seller.CID) {
		return 0, ErrAlreadyExists
	}
	id, err := s.repository.Save(ctx, seller)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *service) Update(ctx context.Context, seller domain.UpdateSeller) (domain.UpdateSeller, error) {
	if seller.CID != nil {
		if s.repository.Exists(ctx, *seller.CID) {
			return domain.UpdateSeller{}, ErrAlreadyExists
		}
	}
	err := s.repository.Update(ctx, seller)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return domain.UpdateSeller{}, ErrNotFound
		} else {
			return domain.UpdateSeller{}, err
		}
	}
	return seller, nil

}
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func NewService(r Repository) Service {
	return &service{repository: r}
}
