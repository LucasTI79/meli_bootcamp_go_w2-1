package domain

import (
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
)

const (
	invalidDateFormat = "invalid date format"
)

type Section struct {
	ID                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseID        int `json:"warehouse_id"`
	ProductTypeID      int `json:"product_type_id"`
}

type UpdateSection struct {
	ID                 *int `json:"id"`
	SectionNumber      *int `json:"section_number"`
	CurrentTemperature *int `json:"current_temperature"`
	MinimumTemperature *int `json:"minimum_temperature"`
	CurrentCapacity    *int `json:"current_capacity"`
	MinimumCapacity    *int `json:"minimum_capacity"`
	MaximumCapacity    *int `json:"maximum_capacity"`
	WarehouseID        *int `json:"warehouse_id"`
	ProductTypeID      *int `json:"product_type_id"`
}

func (s *Section) Overlap(section UpdateSection) {
	s.ID = helpers.Fill(section.ID, s.ID).(int)
	s.SectionNumber = helpers.Fill(section.SectionNumber, s.SectionNumber).(int)
	s.CurrentTemperature = helpers.Fill(section.CurrentTemperature, s.CurrentTemperature).(int)
	s.MinimumTemperature = helpers.Fill(section.MinimumTemperature, s.MinimumTemperature).(int)
	s.CurrentCapacity = helpers.Fill(section.CurrentCapacity, s.CurrentCapacity).(int)
	s.MinimumCapacity = helpers.Fill(section.MinimumCapacity, s.MinimumCapacity).(int)
	s.MaximumCapacity = helpers.Fill(section.MaximumCapacity, s.MaximumCapacity).(int)
	s.WarehouseID = helpers.Fill(section.WarehouseID, s.WarehouseID).(int)
	s.ProductTypeID = helpers.Fill(section.ProductTypeID, s.ProductTypeID).(int)
}

type ProductsBySectionReport struct {
	SectionID     int `json:"section_id"`
	SectionNumber int `json:"section_number"`
	ProductsCount int `json:"products_count"`
}

const dataModel = "2023-07-07 17:54:09"

func (pb *ProductBatches) Validate() error {
	_, err := time.Parse(dataModel, pb.ManufacturingDate)
	if err != nil {
		return apperr.NewResourceNotFound(invalidDateFormat, pb.ManufacturingDate)
	}
	_, err = time.Parse(dataModel, pb.DueDate)
	if err != nil {
		return apperr.NewResourceNotFound(invalidDateFormat, pb.DueDate)
	}
	return nil
}
