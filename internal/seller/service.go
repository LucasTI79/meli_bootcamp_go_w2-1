package seller

import (
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
)

type Service interface {
	GetAll() ([]domain.Seller, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}
