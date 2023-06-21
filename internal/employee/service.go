package employee

import (
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

type EmployeeRequest struct {
	Id             int    `json:"id"`
	Card_number_id string `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Warehouse_id   int    `json:"warehouse_id"`
}

type Service interface {
	GetAll() ([]domain.Employee, error)
	Get(id int) (domain.Employee, error)
	Exists(cardNumberID string) (string, error)
	Save(card_number_id, first_name, last_name string, warehouse_id int) (int, error)
	Update(EmployeeRequest) (EmployeeRequest, error)
	Delete(id int) error
}
type service struct {
	repository Repository
}

var (
	ErrNotFound = errors.New("Funcionário não encontrado.")
)

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Employee, error) {
	employees, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *service) Save(card_number_id, first_name, last_name string, warehouse_id int) (int, error) {

	employeeId, err := s.repository.Save(card_number_id, first_name, last_name, warehouse_id)
	if err != nil {
		return 0, err
	}

	return employeeId, nil
}

func (s *service) Update(employee EmployeeRequest) (EmployeeRequest, error) {
	err := s.repository.Update(employee)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return EmployeeRequest{}, ErrNotFound
		} else {
			return EmployeeRequest{}, err
		}
	}
	return employee, nil
}

func (s *service) Exists(cardNumberID string) (string, error) {
	cardNumber, err := s.repository.Exists(cardNumberID)

	return cardNumber, err
}

func (s *service) Get(id int) (domain.Employee, error) {
	employee, err := s.repository.Get(id)

	return employee, err

}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)

	return err
}
