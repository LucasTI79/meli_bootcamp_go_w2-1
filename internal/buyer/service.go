package buyer

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
) 

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
)

type IService interface{
	GetAll(c *gin.Context) ([]domain.Buyer, error)
	
}

type service struct{
	repository IRepository
}

func NewService(r IRepository) IService {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(c *gin.Context) ([]domain.Buyer, error) {
	buyers, err := s.repository.GetAll(c)
	if err != nil {
		return nil, err
	}

	return buyers, nil
}