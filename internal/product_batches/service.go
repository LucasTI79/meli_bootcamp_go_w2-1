package product_batches

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
)

const (
	ResourceAlreadyExists = "Batch number %d already exists"
)

type Service interface {
	Create(pb domain.ProductBatches) (*domain.ProductBatches, error)
}
type service struct {
	repository        Repository
	productRepository product.Repository
	sectionRepository section.Repository
}

func NewService(repository Repository, productRepository product.Repository, sectionRepository section.Repository) Service {
	return &service{
		repository,
		productRepository,
		sectionRepository,
	}
}

func (s *service) Create(pb domain.ProductBatches) (*domain.ProductBatches, error) {
	if s.repository.Exists(pb.BatchNumber) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, pb.BatchNumber)
	}
	productFound := s.productRepository.Get(context.TODO(), pb.ProductID)
	if productFound == nil {
		return nil, apperr.NewResourceNotFound("Product with id %d does not exist", pb.ProductID)
	}
	sectionFound := s.sectionRepository.Get(pb.SectionID)
	if sectionFound == nil {
		return nil, apperr.NewResourceNotFound("Section with id %d does not exist", pb.SectionID)
	}
	id := s.repository.Save(pb)

	return s.repository.Get(id), nil
}
