package domain

type ProductBatches struct {
	ID                 int     `json:"id"`
	BatchNumber        string  `json:"batch_number"`
	CurrentQuantity    int     `json:"current_quantity"`
	CurrentTemperature float64 `json:"current_temperature"`
	DueDate            string  `json:"due_date"`
	InitialQuantity    int     `json:"initial_quantity"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  string  `json:"manufacturing_date"`
	MinimumTemperature int     `json:"minimum_temperature"`
	ProductID          int     `json:"product_id"`
	SectionID          int     `json:"section_id"`
}
