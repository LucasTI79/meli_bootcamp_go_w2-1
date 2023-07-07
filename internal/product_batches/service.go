package productbatches

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ProductBatchNotFound  = "lote de produto não encontrado com o id %d"
	SectionNotFound       = "seção não encontrada com o id %d"
	ProductNotFound       = "produto não encontrado com o id %d"
	ResourceAlreadyExists = "um lote de produto com o número de lote '%s' já existe"
)

type Service interface {
	Exists(BatchNumber string) bool
	CheckSectionExists(id int) bool
	CheckProductExists(id int) bool
	Create(productBatch domain.ProductBatches) (*domain.ProductBatches, error)
}

type service struct {
	repository     Repository
	sectionService section.Service
	productService product.Service
}

func NewService(repository Repository, sectionService section.Service, productService product.Service) Service {
	return &service{repository, sectionService, productService}
}

func (s *service) Exists(BatchNumber string) bool {
	return s.repository.Exists(BatchNumber)
}

func (s *service) CheckSectionExists(id int) bool {
	return s.repository.CheckSectionExists(id)
}

func (s *service) CheckProductExists(id int) bool {
	return s.repository.CheckProductExists(id)
}

func (s *service) Create(productBatch domain.ProductBatches) (*domain.ProductBatches, error) {
	if s.repository.Exists(productBatch.BatchNumber) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, productBatch.BatchNumber)
	}

	if !s.CheckSectionExists(productBatch.SectionID) {
		return nil, apperr.NewDependentResourceNotFound(SectionNotFound, productBatch.SectionID)
	}

	if !s.CheckProductExists(productBatch.ProductID) {
		return nil, apperr.NewDependentResourceNotFound(ProductNotFound, productBatch.ProductID)
	}

	id := s.repository.Save(productBatch)
	return s.repository.Get(id), nil
}
