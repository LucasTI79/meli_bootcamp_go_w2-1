package domain

import "github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"

type Warehouse struct {
	ID                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimumCapacity    int    `json:"minimum_capacity"`
	MinimumTemperature int    `json:"minimum_temperature"`
<<<<<<< HEAD
	LocalityId         int    `json:"locality_id"`
=======
	LocalityID         int    `json:"locality_id"`
>>>>>>> 72a09cc (adjust section handler)
}

func (w *Warehouse) Overlap(updateWarehouse UpdateWarehouse) {
	w.ID = helpers.Fill(updateWarehouse.ID, w.ID).(int)
	w.Address = helpers.Fill(updateWarehouse.Address, w.Address).(string)
	w.Telephone = helpers.Fill(updateWarehouse.Telephone, w.Telephone).(string)
	w.WarehouseCode = helpers.Fill(updateWarehouse.WarehouseCode, w.WarehouseCode).(string)
	w.MinimumCapacity = helpers.Fill(updateWarehouse.MinimumCapacity, w.MinimumCapacity).(int)
	w.MinimumTemperature = helpers.Fill(updateWarehouse.MinimumTemperature, w.MinimumTemperature).(int)
	w.LocalityID = helpers.Fill(updateWarehouse.LocalityID, w.LocalityID).(int)
}

type UpdateWarehouse struct {
	ID                 *int    `json:"id"`
	Address            *string `json:"address"`
	Telephone          *string `json:"telephone"`
	WarehouseCode      *string `json:"warehouse_code"`
	MinimumCapacity    *int    `json:"minimum_capacity"`
	MinimumTemperature *int    `json:"minimum_temperature"`
	LocalityID         *int    `json:"locality_id"`
}
