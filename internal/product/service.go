package product

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	apperr "github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/errors"
)

type Service interface {
	GetAll(context.Context) ([]domain.Product, error)
	Get(context.Context, int) (domain.Product, error)
	Create(context.Context, domain.Product) (*domain.Product, error)
	Update(context.Context, int, domain.ProductOptional) (*domain.Product, error)
	Delete(context.Context, int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.Get(ctx, id)
}

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

func (s *service) Update(ctx context.Context, id int, product domain.ProductOptional) (*domain.Product, error) {
	productFound, err := s.repository.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	if product.ProductCode.HasValue {
		productCode := product.ProductCode.Value.(*string)
		productCodeExists := s.repository.Exists(ctx, *productCode)

		if productCodeExists && *productCode != productFound.ProductCode {
			return nil, apperr.NewResourceAlreadyExists("product with code '%s' already exists", *productCode)
		}
	}

	productFound = productFound.Overlap(product)

	err = s.repository.Update(ctx, productFound)

	if err != nil {
		return nil, err
	}

	productResponse, _ := s.repository.Get(ctx, id)
	return &productResponse, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
