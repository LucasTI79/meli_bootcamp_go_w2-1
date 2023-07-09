package product_record

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
)

const (
	ProductRecordNotFound = "registro de produto não encontrado com o id %d"
	ProductNotFound       = "produto não encontrado com o id %d"
	ResourceAlreadyExists = "um registro de produto com o id de produto '%d' e última data de atualização `%s` já existe"
)

type Service interface {
	Create(record domain.ProductRecord) (*domain.ProductRecord, error)
	CountSellersByAllProductRecords() []domain.RecordsByProductReport
	CountSellersByProductRecord(id int) (*domain.RecordsByProductReport, error)
}

type service struct {
	repository        Repository
	productRepository product.Repository
}

func NewService(repository Repository, productRepository product.Repository) Service {
	return &service{repository, productRepository}
}

func (s *service) Create(record domain.ProductRecord) (*domain.ProductRecord, error) {
	if s.repository.Exists(record.ProductID, record.LastUpdateDate) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, record.ProductID, helpers.ToFormattedDateTime(record.LastUpdateDate))
	}

	productFound := s.productRepository.Get(record.ProductID)

	if productFound == nil {
		return nil, apperr.NewDependentResourceNotFound(ProductNotFound, record.ProductID)
	}

	id := s.repository.Save(record)
	return s.repository.Get(id), nil
}

func (s *service) CountSellersByAllProductRecords() []domain.RecordsByProductReport {
	return s.repository.CountRecordsByAllProducts()
}

func (s *service) CountSellersByProductRecord(id int) (*domain.RecordsByProductReport, error) {
	record := s.repository.Get(id)

	if record == nil {
		return nil, apperr.NewResourceNotFound(ProductRecordNotFound, id)
	}

	return s.repository.CountRecordsByProduct(id), nil
}
