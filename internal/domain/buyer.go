package domain

import "github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"

type Buyer struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type UpdateBuyer struct {
	ID           *int    `json:"id"`
	CardNumberID *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
}

type PurchasesByBuyerReport struct {
	ID             int    `json:"id"`
	CardNumberID   string `json:"card_number_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	PurchasesCount int    `json:"purchase_orders_count"`
}

func (p *Buyer) Overlap(product UpdateBuyer) {
	p.ID = helpers.Fill(product.ID, p.ID).(int)
	p.CardNumberID = helpers.Fill(product.CardNumberID, p.CardNumberID).(string)
	p.FirstName = helpers.Fill(product.FirstName, p.FirstName).(string)
	p.LastName = helpers.Fill(product.LastName, p.LastName).(string)
}
