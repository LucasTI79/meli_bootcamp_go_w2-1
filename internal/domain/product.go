package domain

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
)

// Product represents an underlying URL with statistics on how it is used.
type Product struct {
	ID             int     `json:"id"`
	Description    string  `json:"description"`
	ExpirationRate float32 `json:"expiration_rate"`
	FreezingRate   float32 `json:"freezing_rate"`
	Height         float32 `json:"height"`
	Length         float32 `json:"length"`
	Netweight      float32 `json:"netweight"`
	ProductCode    string  `json:"product_code"`
	RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
	Width          float32 `json:"width"`
	ProductTypeID  int     `json:"product_type_id"`
	SellerID       int     `json:"seller_id"`
}

type UpdateProduct struct {
	ID             *int     `json:"id"`
	Description    *string  `json:"description"`
	ExpirationRate *float32 `json:"expiration_rate"`
	FreezingRate   *float32 `json:"freezing_rate"`
	Height         *float32 `json:"height"`
	Length         *float32 `json:"length"`
	Netweight      *float32 `json:"netweight"`
	ProductCode    *string  `json:"product_code"`
	RecomFreezTemp *float32 `json:"recommended_freezing_temperature"`
	Width          *float32 `json:"width"`
	ProductTypeID  *int     `json:"product_type_id"`
	SellerID       *int     `json:"seller_id"`
}

func (p *Product) Overlap(product UpdateProduct) {
	p.ID = helpers.Fill(product.ID, p.ID).(int)
	p.Description = helpers.Fill(product.Description, p.Description).(string)
	p.ExpirationRate = helpers.Fill(product.ExpirationRate, p.ExpirationRate).(float32)
	p.FreezingRate = helpers.Fill(product.FreezingRate, p.FreezingRate).(float32)
	p.Height = helpers.Fill(product.Height, p.Height).(float32)
	p.Length = helpers.Fill(product.Length, p.Length).(float32)
	p.Netweight = helpers.Fill(product.Netweight, p.Netweight).(float32)
	p.ProductCode = helpers.Fill(product.ProductCode, p.ProductCode).(string)
	p.RecomFreezTemp = helpers.Fill(product.RecomFreezTemp, p.RecomFreezTemp).(float32)
	p.Width = helpers.Fill(product.Width, p.Width).(float32)
	p.ProductTypeID = helpers.Fill(product.ProductTypeID, p.ProductTypeID).(int)
	p.SellerID = helpers.Fill(product.SellerID, p.SellerID).(int)
}
