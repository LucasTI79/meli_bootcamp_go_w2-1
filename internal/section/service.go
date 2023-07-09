package section

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "produto não encontrado com o id %d"
	ResourceAlreadyExists = "um produto com o código '%d' já existe"
)

type Service interface {
	GetAll(context.Context) []domain.Section
	Get(context.Context, int) (*domain.Section, error)
	Create(ctx context.Context, sc domain.Section) (*domain.Section, error)
	Update(context.Context, int, domain.UpdateSection) (*domain.Section, error)
	Delete(context.Context, int) error
	ExistsSectionID(productID int) error
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(c context.Context) []domain.Section {
	return s.repository.GetAll(c)
}

func (s *service) Get(c context.Context, id int) (*domain.Section, error) {
	section := s.repository.Get(c, id)

	if section == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return section, nil
}

func (s *service) Create(ctx context.Context, sc domain.Section) (*domain.Section, error) {
	if s.repository.Exists(ctx, sc.SectionNumber) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, sc.SectionNumber)
	}

	id := s.repository.Save(ctx, sc)

	return s.repository.Get(ctx, id), nil
}

func (s *service) Update(ctx context.Context, id int, section domain.UpdateSection) (*domain.Section, error) {
	sectionFound := s.repository.Get(ctx, id)

	if sectionFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if section.SectionNumber != nil {
		sectionNumber := *section.SectionNumber
		sectionNumberExists := s.repository.Exists(ctx, sectionNumber)

		if sectionNumberExists && sectionNumber != sectionFound.SectionNumber {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, sectionNumber)
		}
	}

	sectionFound.Overlap(section)
	s.repository.Update(ctx, *sectionFound)
	return s.repository.Get(ctx, id), nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	section := s.repository.Get(ctx, id)

	if section == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(ctx, id)
	return nil
}

func (s *service) ExistsSectionID(productID int) error {
	sectionExists := s.repository.Exists(context.Background(), productID)
	if !sectionExists {
		return apperr.NewResourceNotFound(ResourceNotFound, productID)
	}
	return nil
}
