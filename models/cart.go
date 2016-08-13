package models

import (
	"errors"
)

type Cart struct {
	Id          uint   `json:"-"`
	AccessToken string `json:"access_token,omitempty"`
	Status      string `json:"-"`

	Items []CartItem `json:"items"`

	Amounts uint `json:"amounts" sql:"-"`
}

type CartItem struct {
	Id                       uint    `json:"item_id" sql:"id"`
	CartId                   uint    `json:"-" sql:"REFERENCES carts(id)"`
	ProductId                uint    `json:"product_id" sql:"REFERENCES products(id)"`
	ProductName              string  `json:"name" sql:"-"`
	ProductImageThumbnailUrl *string `json:"image_thumbnail_url" sql:"-"`
	ProductPrice             int     `json:"price" sql:"-"`
	Quantity                 int     `json:"quantity"`
}

func (cart *Cart) UpdateItems(product_id uint, quantity int) error {
	for idx, item := range cart.Items {
		if item.ProductId == product_id {
			cart.Items[idx].Quantity += quantity
			if cart.Items[idx].Quantity < 0 {
				cart.Items[idx].Quantity = 0
			}
			return nil
		}
	}

	if quantity < 0 {
		return errors.New("Quantity for new item should bigger than 0")
	}

	cart.Items = append(cart.Items,
		CartItem{
			ProductId: product_id,
			Quantity:  quantity,
		})
	return nil
}

func (cart *Cart) CalculateAmount() {
	for _, item := range cart.Items {
		cart.Amounts += uint(item.ProductPrice) * uint(item.Quantity)
	}
}

func (cart *Cart) RemoveZeroQuantityItems() {
	for idx, _ := range cart.Items {
		if cart.Items[idx].Quantity <= 0 {
			cart.Items = append(cart.Items[:idx], cart.Items[idx+1:]...)
			return
		}
	}
}
