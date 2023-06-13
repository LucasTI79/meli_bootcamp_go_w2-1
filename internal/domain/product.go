package domain

type Optional struct {
	Value    interface{}
	HasValue bool
}

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

func (p Product) ToProductOptional() ProductOptional {
	return ProductOptional{
		ID:             Optional{Value: p.ID, HasValue: true},
		Description:    Optional{Value: p.Description, HasValue: true},
		ExpirationRate: Optional{Value: p.ExpirationRate, HasValue: true},
		FreezingRate:   Optional{Value: p.FreezingRate, HasValue: true},
		Height:         Optional{Value: p.Height, HasValue: true},
		Length:         Optional{Value: p.Length, HasValue: true},
		Netweight:      Optional{Value: p.Netweight, HasValue: true},
		ProductCode:    Optional{Value: p.ProductCode, HasValue: true},
		RecomFreezTemp: Optional{Value: p.RecomFreezTemp, HasValue: true},
		Width:          Optional{Value: p.Width, HasValue: true},
		ProductTypeID:  Optional{Value: p.ProductTypeID, HasValue: true},
		SellerID:       Optional{Value: p.SellerID, HasValue: true},
	}
}

func (p Product) Overlap(productOptional ProductOptional) Product {
	var product Product = p

	if productOptional.ID.HasValue {
		product.ID = *productOptional.ID.Value.(*int)
	}

	if productOptional.Description.HasValue {
		product.Description = *productOptional.Description.Value.(*string)
	}

	if productOptional.ExpirationRate.HasValue {
		product.ExpirationRate = *productOptional.ExpirationRate.Value.(*float32)
	}

	if productOptional.FreezingRate.HasValue {
		product.FreezingRate = *productOptional.FreezingRate.Value.(*float32)
	}

	if productOptional.Height.HasValue {
		product.Height = *productOptional.Height.Value.(*float32)
	}

	if productOptional.Length.HasValue {
		product.Length = *productOptional.Length.Value.(*float32)
	}

	if productOptional.Netweight.HasValue {
		product.Netweight = *productOptional.Netweight.Value.(*float32)
	}

	if productOptional.ProductCode.HasValue {
		product.ProductCode = *productOptional.ProductCode.Value.(*string)
	}

	if productOptional.RecomFreezTemp.HasValue {
		product.RecomFreezTemp = *productOptional.RecomFreezTemp.Value.(*float32)
	}

	if productOptional.Width.HasValue {
		product.Width = *productOptional.Width.Value.(*float32)
	}

	if productOptional.ProductTypeID.HasValue {
		product.ProductTypeID = *productOptional.ProductTypeID.Value.(*int)
	}

	if productOptional.SellerID.HasValue {
		product.SellerID = *productOptional.SellerID.Value.(*int)
	}

	return product
}

type ProductOptional struct {
	ID             Optional
	Description    Optional
	ExpirationRate Optional
	FreezingRate   Optional
	Height         Optional
	Length         Optional
	Netweight      Optional
	ProductCode    Optional
	RecomFreezTemp Optional
	Width          Optional
	ProductTypeID  Optional
	SellerID       Optional
}

func (p ProductOptional) ToProduct() Product {
	return Product{
		ID:             p.ID.Value.(int),
		Description:    p.Description.Value.(string),
		ExpirationRate: p.ExpirationRate.Value.(float32),
		FreezingRate:   p.FreezingRate.Value.(float32),
		Height:         p.Height.Value.(float32),
		Length:         p.Length.Value.(float32),
		Netweight:      p.Netweight.Value.(float32),
		ProductCode:    p.ProductCode.Value.(string),
		RecomFreezTemp: p.RecomFreezTemp.Value.(float32),
		Width:          p.Width.Value.(float32),
		ProductTypeID:  p.ProductTypeID.Value.(int),
		SellerID:       p.SellerID.Value.(int),
	}
}
