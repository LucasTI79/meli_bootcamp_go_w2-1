package product_batches

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

type Service interface {
	Save(pb domain.ProductBatches) (int, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Save(pb domain.ProductBatches) (int, error) {
	productBatchID, err := s.repository.Save(pb)
	return productBatchID, err
}
