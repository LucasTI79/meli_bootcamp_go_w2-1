package product_batches

import (
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

type ProductBatches struct {
	ID                 int       `json:"id"`
	BatchNumber        int       `json:"batch_number"`
	CurrentQuantity    int       `json:"current_quantity"`
	CurrentTemperature float64   `json:"current_temperature"`
	DueDate            time.Time `json:"due_date"`
	InitialQuantity    int       `json:"initial_quantity"`
	ManufacturingDate  time.Time `json:"manufacturing_date"`
	ManufacturingHour  string    `json:"manufacturing_hour"`
	MinimumTemperature int       `json:"minimum_temperature"`
	ProductID          int       `json:"product_id"`
	SectionID          int       `json:"section_id"`
}

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
	existsProductBatchNumber := s.repository.Exists(pb.BatchNumber)
	if existsProductBatchNumber {
		return nil, apperr.NewResourceNotFound("Invalid batch number")
	}
	id := s.repository.Save(pb)

	return s.repository.Get(id), nil
}

func (s *service) Exists(batchNumber int) (bool, error) {
	return s.repository.Exists(batchNumber), nil
}

func (s *service) CountProductsByAllSections() ([]domain.ProductsBySectionReport, error) {
	productsbratches, err := s.repository.CountProductsByAllSections()
	if err != nil {
		return productsbratches, err
	}
	return productsbratches, nil
}

func (s *service) CountProductsBySection(id int) ([]domain.ProductsBySectionReport, error) {
	productsbratches, err := s.repository.CountProductsBySection(id)
	if err != nil {
		return productsbratches, err
	}
	return productsbratches, nil
}
