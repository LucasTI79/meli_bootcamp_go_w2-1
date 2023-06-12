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
	Save(c *gin.Context, b domain.Request) (int, error)
	
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

func (s *service) Save(c *gin.Context, b domain.Request) (int, error) {
	exists := s.repository.Exists(c, b.CardNumberID)

	if !exists {
		id, err := s.repository.Save(c, b)
		if err != nil {
			return 0, err
		}
		return id, nil 
	} else {
		return 0, errors.New("Nao é possível cadastrar um comprador com Card Number repetido.")
	}

}
