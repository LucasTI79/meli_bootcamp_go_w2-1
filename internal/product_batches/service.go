package product_batches

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceAlreadyExists = "Batch number %d already exists"
)

type Service interface {
	Create(pb domain.ProductBatches) (*domain.ProductBatches, error)
	Exists(batchNumber int) (bool, error)
	CountProductsByAllSections() ([]domain.ProductsBySectionReport, error)
	CountProductsBySection(id int) ([]domain.ProductsBySectionReport, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(pb domain.ProductBatches) (*domain.ProductBatches, error) {
	if s.repository.Exists(pb.BatchNumber) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, pb.BatchNumber)
	}
	id := s.repository.Save(pb)
	pd := s.repository.Get(id)
	return pd, nil
}

func (s *service) Exists(batchNumber int) (bool, error) {
	return s.repository.Exists(batchNumber), nil
}

func (s *service) CountProductsByAllSections() ([]domain.ProductsBySectionReport, error) {
	productsBatches, err := s.repository.CountProductsByAllSections()
	if err != nil {
		return productsBatches, err
	}
	return productsBatches, nil
}

func (s *service) CountProductsBySection(id int) ([]domain.ProductsBySectionReport, error) {
	productsBatches, err := s.repository.CountProductsBySection(id)
	if err != nil {
		return productsBatches, err
	}
	return productsBatches, nil
}
