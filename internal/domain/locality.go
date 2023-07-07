package domain

type Locality struct {
	ID           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceID   int    `json:"province_id,omitempty"`
}

type SellersByLocalityReport struct {
	ID           *int    `json:"locality_id"`
	LocalityName *string `json:"locality_name"`
	SellersCount *int    `json:"sellers_count"`
}
