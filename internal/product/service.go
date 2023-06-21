package product

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "produto não encontrado com o id %d"
	ResourceAlreadyExists = "um produto com o código '%s' já existe"
)

type Service interface {
	GetAll(context.Context) []domain.Product
	Get(context.Context, int) (*domain.Product, error)
	Create(context.Context, domain.Product) (*domain.Product, error)
	Update(context.Context, int, domain.UpdateProduct) (*domain.Product, error)
	Delete(context.Context, int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll(ctx context.Context) []domain.Product {
	return s.repository.GetAll(ctx)
}

func (s *service) Get(ctx context.Context, id int) (*domain.Product, error) {
	product := s.repository.Get(ctx, id)

	if product == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return product, nil
}

func (s *service) Create(ctx context.Context, product domain.Product) (*domain.Product, error) {
	if s.repository.Exists(ctx, product.ProductCode) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, product.ProductCode)
	}

	id := s.repository.Save(ctx, product)
	created := s.repository.Get(ctx, id)

	if created == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return created, nil
}

func (s *service) Update(ctx context.Context, id int, product domain.UpdateProduct) (*domain.Product, error) {
	productFound := s.repository.Get(ctx, id)

	if productFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if product.ProductCode != nil {
		productCode := *product.ProductCode
		productCodeExists := s.repository.Exists(ctx, productCode)

		if productCodeExists && productCode != productFound.ProductCode {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, productCode)
		}
	}

	productFound.Overlap(product)
	s.repository.Update(ctx, *productFound)
	updated := s.repository.Get(ctx, id)

	if updated == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return updated, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	product := s.repository.Get(ctx, id)

	if product == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(ctx, id)
	return nil
}
