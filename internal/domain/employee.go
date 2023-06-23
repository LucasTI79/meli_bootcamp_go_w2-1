package domain

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
)

type Employee struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}

type UpdateEmployee struct {
	ID           *int    `json:"id"`
	CardNumberID *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	WarehouseID  *int    `json:"warehouse_id"`
}

func (e *Employee) Overlap(employee UpdateEmployee) {
	e.ID = helpers.Fill(employee.ID, e.ID).(int)
	e.CardNumberID = helpers.Fill(employee.CardNumberID, e.CardNumberID).(string)
	e.FirstName = helpers.Fill(employee.FirstName, e.FirstName).(string)
	e.LastName = helpers.Fill(employee.LastName, e.LastName).(string)
	e.WarehouseID = helpers.Fill(employee.WarehouseID, e.WarehouseID).(int)
}