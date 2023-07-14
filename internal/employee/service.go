package employee

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound = "empregado não encontrado com o id %d"
	WarehouseNotFound = "armazém não encontrado com o id '%d'"
	ResourceAlreadyExists = "um empregado com card number ID '%s' já existe"
)

type Service interface {
	GetAll() []domain.Employee
	Get(int) (*domain.Employee, error)
	Create(domain.Employee) (*domain.Employee, error)
	Update(int, domain.UpdateEmployee) (*domain.Employee, error)
	Delete(int) error
	CountInboundOrdersByAllEmployees() []domain.InboundOrdersByEmployee
	CountInboundOrdersByEmployee(id int) (*domain.InboundOrdersByEmployee, error)
}

type service struct {
	repository Repository
	warehouseRepository warehouse.Repository
}

func NewService(r Repository, w warehouse.Repository) Service {
	return &service{
		repository: r,
		warehouseRepository: w,
	}
}

func (s *service) GetAll() []domain.Employee {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (*domain.Employee, error) {
	employee := s.repository.Get(id)

	if employee == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return employee, nil
}

func (s *service) Create(employee domain.Employee) (*domain.Employee, error) {
	if s.repository.Exists(employee.CardNumberID) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, employee.CardNumberID)
	}

	w := s.warehouseRepository.Get(employee.WarehouseID)

	if w == nil {
		return nil, apperr.NewDependentResourceNotFound(WarehouseNotFound, employee.WarehouseID)
	}

	id := s.repository.Save(employee)
	created := s.repository.Get(id)

	return created, nil
}

func (s *service) Update(id int, employee domain.UpdateEmployee) (*domain.Employee, error) {
	employeeFound := s.repository.Get(id)

	if employeeFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}
	
	if employee.CardNumberID != nil {
		employeeCardNumber := *employee.CardNumberID 
		employeeCardNumberExists := s.repository.Exists(employeeCardNumber)

		if employeeCardNumberExists && employeeCardNumber != employeeFound.CardNumberID {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, employeeCardNumber)
		}
	}

	employeeFound.Overlap(employee)

	w := s.warehouseRepository.Get(employeeFound.WarehouseID)

	if w == nil {
		return nil, apperr.NewDependentResourceNotFound(WarehouseNotFound, employeeFound.WarehouseID)
	}

	s.repository.Update(*employeeFound)
	return s.repository.Get(id), nil

}

func (s *service) Delete(id int) error {
	employee := s.repository.Get(id)

	if employee == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(id)
	return nil
}

func (s *service) CountInboundOrdersByAllEmployees() []domain.InboundOrdersByEmployee {
	return s.repository.CountInboundOrdersByAllEmployees()
}

func (s *service) CountInboundOrdersByEmployee(id int) (*domain.InboundOrdersByEmployee, error) {
	employee := s.repository.Get(id)

	if employee == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return s.repository.CountInboundOrdersByEmployee(id), nil
}
