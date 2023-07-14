package product_batch

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
)

const (
	ResourceAlreadyExists = "batch number %d already exists"
	ProductNotFound       = "product not found with id %d"
	SectionNotFound       = "section not found with id %d"
)

type Service interface {
	Create(pb domain.ProductBatch) (*domain.ProductBatch, error)
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

func (s *service) Create(pb domain.ProductBatch) (*domain.ProductBatch, error) {
	if s.repository.Exists(pb.BatchNumber) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, pb.BatchNumber)
	}

	productFound := s.productRepository.Get(pb.ProductID)

	if productFound == nil {
		return nil, apperr.NewDependentResourceNotFound(ProductNotFound, pb.ProductID)
	}

	sectionFound := s.sectionRepository.Get(pb.SectionID)

	if sectionFound == nil {
		return nil, apperr.NewDependentResourceNotFound(SectionNotFound, pb.SectionID)
	}
	id := s.repository.Save(pb)

	return s.repository.Get(id), nil
}
