package domain

type Seller struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}
type CreateSeller struct {
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}
type UpdateSeller struct {
	ID          int     `json:"id"`
	CID         *int    `json:"cid,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	Address     *string `json:"address,omitempty"`
	Telephone   *string `json:"telephone,omitempty"`
}
