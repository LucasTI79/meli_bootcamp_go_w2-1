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
	Get(c *gin.Context, id int) (domain.Buyer, error)
	Save(c *gin.Context, b domain.Request) (int, error)
	Delete(c *gin.Context, id int) error
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

func (s *service) Get(c *gin.Context, id int) (domain.Buyer, error) {
	buyer, err := s.repository.Get(c, id)
	if err != nil {
		return domain.Buyer{}, err
	}

	return buyer, nil
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

func (s *service) Delete(c *gin.Context, id int) error {
	
	return s.repository.Delete(c, id)

}