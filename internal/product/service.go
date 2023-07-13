package product

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_type"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "produto não encontrado com o id %d"
	ResourceAlreadyExists = "um produto com o código '%s' já existe"
	ProductTypeNotFound   = "tipo de produto não encontrado com o id %d"
	SellerNotFound        = "vendedor não encontrado com o id %d"
)

type Service interface {
	GetAll() []domain.Product
	Get(int) (*domain.Product, error)
	Create(domain.Product) (*domain.Product, error)
	Update(int, domain.UpdateProduct) (*domain.Product, error)
	Delete(int) error
	CountRecordsByAllProducts() []domain.RecordsByProductReport
	CountRecordsByProduct(id int) (*domain.RecordsByProductReport, error)
}

type service struct {
	repository            Repository
	productTypeRepository product_type.Repository
	sellerRepository      seller.Repository
}

func NewService(repository Repository, productTypeRepository product_type.Repository, sellerRerepository seller.Repository) Service {
	return &service{repository, productTypeRepository, sellerRerepository}
}

func (s *service) GetAll() []domain.Product {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (*domain.Product, error) {
	product := s.repository.Get(id)

	if product == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return product, nil
}

func (s *service) Create(product domain.Product) (*domain.Product, error) {
	if s.repository.Exists(product.ProductCode) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, product.ProductCode)
	}

	productTypeFound := s.productTypeRepository.Get(product.ProductTypeID)

	if productTypeFound == nil {
		return nil, apperr.NewDependentResourceNotFound(ProductTypeNotFound, product.ProductTypeID)
	}

	sellerFound := s.sellerRepository.Get(product.SellerID)

	if sellerFound == nil {
		return nil, apperr.NewDependentResourceNotFound(SellerNotFound, product.SellerID)
	}

	id := s.repository.Save(product)
	return s.repository.Get(id), nil
}

func (s *service) Update(id int, product domain.UpdateProduct) (*domain.Product, error) {
	productFound := s.repository.Get(id)

	if productFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if product.ProductCode != nil {
		productCode := *product.ProductCode
		productCodeExists := s.repository.Exists(productCode)

		if productCodeExists && productCode != productFound.ProductCode {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, productCode)
		}
	}

	productFound.Overlap(product)

	productTypeFound := s.productTypeRepository.Get(productFound.ProductTypeID)

	if productTypeFound == nil {
		return nil, apperr.NewDependentResourceNotFound(ProductTypeNotFound, product.ProductTypeID)
	}

	sellerFound := s.sellerRepository.Get(productFound.SellerID)

	if sellerFound == nil {
		return nil, apperr.NewDependentResourceNotFound(SellerNotFound, product.SellerID)
	}

	s.repository.Update(*productFound)
	return s.repository.Get(id), nil
}

func (s *service) Delete(id int) error {
	product := s.repository.Get(id)

	if product == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(id)
	return nil
}

func (s *service) CountRecordsByAllProducts() []domain.RecordsByProductReport {
	return s.repository.CountRecordsByAllProducts()
}

func (s *service) CountRecordsByProduct(id int) (*domain.RecordsByProductReport, error) {
	productFound := s.repository.Get(id)

	if productFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return s.repository.CountRecordsByProduct(id), nil
}
