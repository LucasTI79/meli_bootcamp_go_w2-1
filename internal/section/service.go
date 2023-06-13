package section

import (
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

var (
	ErrNotFound = errors.New("Seção não encontrada.")
)

type Service interface {
	GetAll() ([]domain.Section, error)
	Get(id int) (domain.Section, error)
	Exists(sectionNumber int) (int, error)
	Save(section_number, current_temperature, minimum_temperature, current_capacity, maximum_capacity, warehouse_id, id_product_type int) (int, error)
	Update(domain.Section) error
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Section, error) {
	sections, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return sections, nil
}

func (s *service) Save(section_number, current_temperature, minimum_temperature, current_capacity, maximum_capacity, warehouse_id, id_product_type int) (int, error) {

	sectionId, err := s.repository.Save(section_number, current_temperature, minimum_temperature, current_capacity, maximum_capacity, warehouse_id, id_product_type)
	if err != nil {
		return 0, err
	}

	return sectionId, nil
}

func (s *service) Update(domain.Section) error {
	err := s.repository.Update(domain.Section{})

	return err
}

func (s *service) Exists(section_number int) (int, error) {
	sectionNumber, err := s.repository.Exists(section_number)

	return sectionNumber, err
}

func (s *service) Get(id int) (domain.Section, error) {
	section, err := s.repository.Get(id)

	return section, err

}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	return err
}
