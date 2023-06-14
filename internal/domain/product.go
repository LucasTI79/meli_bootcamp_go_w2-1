package domain

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
	if product.ID != nil {
		p.ID = *product.ID
	}
	if product.Description != nil {
		p.Description = *product.Description
	}
	if product.ExpirationRate != nil {
		p.ExpirationRate = *product.ExpirationRate
	}
	if product.FreezingRate != nil {
		p.FreezingRate = *product.FreezingRate
	}
	if product.Height != nil {
		p.Height = *product.Height
	}
	if product.Length != nil {
		p.Length = *product.Length
	}
	if product.Netweight != nil {
		p.Netweight = *product.Netweight
	}
	if product.ProductCode != nil {
		p.ProductCode = *product.ProductCode
	}
	if product.RecomFreezTemp != nil {
		p.RecomFreezTemp = *product.RecomFreezTemp
	}
	if product.Width != nil {
		p.Width = *product.Width
	}
	if product.ProductTypeID != nil {
		p.ProductTypeID = *product.ProductTypeID
	}
	if product.SellerID != nil {
		p.SellerID = *product.SellerID
	}
}
