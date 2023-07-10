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
	Create(pb domain.ProductBatches) (domain.ProductBatches, error)
	Exists(batchNumber int) (bool, error)
	Get() ([]domain.ProductBatches, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(pb domain.ProductBatches) (domain.ProductBatches, error) {
	existsProductBatchNumber := s.repository.Exists(pb.BatchNumber)
	if existsProductBatchNumber {
		return domain.ProductBatches{}, apperr.NewResourceNotFound("Invalid batch number")
	}

	productBatch := domain.ProductBatches{
		BatchNumber:        pb.BatchNumber,
		CurrentQuantity:    pb.CurrentQuantity,
		CurrentTemperature: pb.CurrentTemperature,
		DueDate:            pb.DueDate,
		InitialQuantity:    pb.InitialQuantity,
		ManufacturingDate:  pb.ManufacturingDate,
		ManufacturingHour:  pb.ManufacturingHour,
		MinimumTemperature: pb.MinimumTemperature,
		ProductID:          pb.ProductID,
		SectionID:          pb.SectionID,
	}

	prodBatches, err := s.repository.Save(productBatch)
	if err != nil {
		return domain.ProductBatches{}, err
	}
	productBatch.ID = prodBatches
	return productBatch, nil
}

func (s *service) Exists(batchNumber int) (bool, error) {
	return s.repository.Exists(batchNumber), nil
}

func (s *service) Get() ([]domain.ProductBatches, error) {
	productsbratches, err := s.repository.Get()
	if err != nil {
		return productsbratches, err
	}
	return productsbratches, nil
}
