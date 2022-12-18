package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductCode        string `json:"product_code"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductQty         int    `json:"product_qty"`
}

func (Product) TableName() string {
	return `products`
}
