package employee

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound = "empregado não encontrado com o id %d"
	ResourceAlreadyExists = "um empregado com card number ID '%s' já existe"
)

type Service interface {
	GetAll(context.Context) []domain.Employee
	Get(context.Context, int) (*domain.Employee, error)
	Create(context.Context, domain.Employee) (*domain.Employee, error)
	Update(context.Context, int, domain.UpdateEmployee) (*domain.Employee, error)
	Delete(context.Context, int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) []domain.Employee {
	return s.repository.GetAll(ctx)
}

func (s *service) Get(ctx context.Context, id int) (*domain.Employee, error) {
	employee := s.repository.Get(ctx, id)

	if employee == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return employee, nil
}

func (s *service) Create(ctx context.Context, employee domain.Employee) (*domain.Employee, error) {
	if s.repository.Exists(ctx, employee.CardNumberID) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, employee.CardNumberID)
	}

	id := s.repository.Save(ctx, employee)
	created := s.repository.Get(ctx, id)

	return created, nil
}

func (s *service) Update(ctx context.Context, id int, employee domain.UpdateEmployee) (*domain.Employee, error) {
	employeeFound := s.repository.Get(ctx, id)

	if employeeFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if employee.CardNumberID != nil {
		employeeCardNumber := *employee.CardNumberID 
		employeeCardNumberExists := s.repository.Exists(ctx, employeeCardNumber)

		if employeeCardNumberExists && employeeCardNumber != employeeFound.CardNumberID {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, employeeCardNumber)
		}
	}

	employeeFound.Overlap(employee)
	s.repository.Update(ctx, *employeeFound)
	updated := s.repository.Get(ctx, id)

	if updated == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return updated, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	employee := s.repository.Get(ctx, id)

	if employee == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(ctx, id)
	return nil
}