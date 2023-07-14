package domain

import "github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"

type Seller struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}

type UpdateSeller struct {
	ID          *int    `json:"id"`
	CID         *int    `json:"cid,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	Address     *string `json:"address,omitempty"`
	Telephone   *string `json:"telephone,omitempty"`
	LocalityID  *int    `json:"locality_id,omitempty"`
}

func (p *Seller) Overlap(product UpdateSeller) {
	p.ID = helpers.Fill(product.ID, p.ID).(int)
	p.CID = helpers.Fill(product.CID, p.CID).(int)
	p.CompanyName = helpers.Fill(product.CompanyName, p.CompanyName).(string)
	p.CompanyName = helpers.Fill(product.CompanyName, p.CompanyName).(string)
	p.Address = helpers.Fill(product.Address, p.Address).(string)
	p.Telephone = helpers.Fill(product.Telephone, p.Telephone).(string)
	p.LocalityID = helpers.Fill(product.LocalityID, p.LocalityID).(int)
}
