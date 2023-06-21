package domain

import "reflect"

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
	p.ID = fill(product.ID, p.ID).(int)
	p.Description = fill(product.Description, p.Description).(string)
	p.ExpirationRate = fill(product.ExpirationRate, p.ExpirationRate).(float32)
	p.FreezingRate = fill(product.FreezingRate, p.FreezingRate).(float32)
	p.Height = fill(product.Height, p.Height).(float32)
	p.Length = fill(product.Length, p.Length).(float32)
	p.Netweight = fill(product.Netweight, p.Netweight).(float32)
	p.ProductCode = fill(product.ProductCode, p.ProductCode).(string)
	p.RecomFreezTemp = fill(product.RecomFreezTemp, p.RecomFreezTemp).(float32)
	p.Width = fill(product.Width, p.Width).(float32)
	p.ProductTypeID = fill(product.ProductTypeID, p.ProductTypeID).(int)
	p.SellerID = fill(product.SellerID, p.SellerID).(int)
}

func fill(first interface{}, second interface{}) interface{} {
	valueOfFirst := reflect.ValueOf(first)
	if valueOfFirst.IsNil() {
		return second
	}

	return valueOfFirst.Elem().Interface()
}
