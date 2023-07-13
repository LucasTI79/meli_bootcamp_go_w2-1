package section

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_type"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "produto não encontrado com o id %d"
	ResourceAlreadyExists = "um produto com o código '%d' já existe"
	WarehouseNotFound     = "armazem não encontrado com o id %d"
	ProductTypeNotFound   = "tipo do produto não encontrado com o id %d"
)

type Service interface {
	GetAll() []domain.Section
	Get(int) (*domain.Section, error)
	Create(sc domain.Section) (*domain.Section, error)
	Update(int, domain.UpdateSection) (*domain.Section, error)
	Delete(int) error
}
type service struct {
	repository            Repository
	warehouseRepository   warehouse.Repository
	productTypeRepository product_type.Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() []domain.Section {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (*domain.Section, error) {
	section := s.repository.Get(id)

	if section == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return section, nil
}

func (s *service) Create(sc domain.Section) (*domain.Section, error) {
	if s.repository.Exists(sc.SectionNumber) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, sc.SectionNumber)
	}

	productTypeById := s.productTypeRepository.Get(sc.ProductTypeID)
	if productTypeById == nil {
		return nil, apperr.NewDependentResourceNotFound(ProductTypeNotFound, sc.ProductTypeID)
	}

	warehouseById := s.warehouseRepository.Get(sc.WarehouseID)

	if warehouseById == nil {
		return nil, apperr.NewDependentResourceNotFound(WarehouseNotFound, sc.ProductTypeID)
	}

	id := s.repository.Save(sc)

	return s.repository.Get(id), nil
}

func (s *service) Update(id int, section domain.UpdateSection) (*domain.Section, error) {
	sectionFound := s.repository.Get(id)

	if sectionFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if section.SectionNumber != nil {
		sectionNumber := *section.SectionNumber
		sectionNumberExists := s.repository.Exists(sectionNumber)

		if sectionNumberExists && sectionNumber != sectionFound.SectionNumber {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, sectionNumber)
		}
	}

	productTypeById := s.productTypeRepository.Get(sectionFound.ProductTypeID)
	if productTypeById == nil {
		return nil, apperr.NewDependentResourceNotFound(ProductTypeNotFound, sectionFound.ProductTypeID)
	}

	warehouseById := s.warehouseRepository.Get(sectionFound.WarehouseID)

	if warehouseById == nil {
		return nil, apperr.NewDependentResourceNotFound(WarehouseNotFound, sectionFound.ProductTypeID)
	}

	sectionFound.Overlap(section)
	s.repository.Update(*sectionFound)
	return s.repository.Get(id), nil
}

func (s *service) Delete(id int) error {
	section := s.repository.Get(id)

	if section == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(id)
	return nil
}
