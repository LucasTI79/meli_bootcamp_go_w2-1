package seller

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "vendedor não encontrado com o id %d"
	ResourceAlreadyExists = "um vededor com o CID '%d' já existe"
)

type Service interface {
	GetAll(ctx context.Context) []domain.Seller
	Get(ctx context.Context, id int) (*domain.Seller, error)
	Create(ctx context.Context, seller domain.Seller) (*domain.Seller, error)
	Update(ctx context.Context, id int, seller domain.UpdateSeller) (*domain.Seller, error)
	Delete(ctc context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll(ctx context.Context) []domain.Seller {
	return s.repository.GetAll(ctx)
}

func (s *service) Get(ctx context.Context, id int) (*domain.Seller, error) {
	seller := s.repository.Get(ctx, id)

	if seller == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return seller, nil
}

func (s *service) Create(ctx context.Context, seller domain.Seller) (*domain.Seller, error) {
	if s.repository.Exists(ctx, seller.CID) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, seller.CID)
	}

	id := s.repository.Save(ctx, seller)
	return s.repository.Get(ctx, id), nil
}

func (s *service) Update(ctx context.Context, id int, seller domain.UpdateSeller) (*domain.Seller, error) {
	sellerFound := s.repository.Get(ctx, id)

	if sellerFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if seller.CID != nil {
		sellerCID := *seller.CID
		sellerCodeExists := s.repository.Exists(ctx, sellerCID)

		if sellerCodeExists && sellerCID != sellerFound.CID {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, sellerCID)
		}
	}

	sellerFound.Overlap(seller)
	s.repository.Update(ctx, *sellerFound)
	return s.repository.Get(ctx, id), nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	seller := s.repository.Get(ctx, id)

	if seller == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(ctx, id)
	return nil
}
