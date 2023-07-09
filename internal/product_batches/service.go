package product_batches

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

type Service interface {
	Create(pb domain.ProductBatches) (domain.ProductBatches, error)
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
