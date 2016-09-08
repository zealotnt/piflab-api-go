package models

import (
	"errors"
)

type Amount struct {
	Subtotal uint `json:"subtotal"`
	Shipping uint `json:"shipping"`
	Total    uint `json:"total"`
}

type Cart struct {
	Id          uint   `json:"-"`
	AccessToken string `json:"access_token,omitempty"`
	Status      string `json:"-"`

	Items []CartItem `json:"items"`

	Amounts Amount `json:"amounts" sql:"-"`
}

type CartItem struct {
	Id                       uint    `json:"id" sql:"id"`
	CartId                   uint    `json:"-" sql:"REFERENCES carts(id)"`
	ProductId                uint    `json:"product_id" sql:"REFERENCES products(id)"`
	ProductName              string  `json:"name" sql:"-"`
	ProductImageThumbnailUrl *string `json:"image_thumbnail_url" sql:"-"`
	ProductPrice             int     `json:"price" sql:"-"`
	Quantity                 int     `json:"quantity"`
}

func (cart *Cart) UpdateItems(product_id *uint, item_id *uint, quantity int) error {
	for idx, item := range cart.Items {
		if product_id != nil {
			if item.ProductId == *product_id {
				cart.Items[idx].Quantity += quantity
				if cart.Items[idx].Quantity < 0 {
					cart.Items[idx].Quantity = 0
				}
				return nil
			}
		}
		if item_id != nil {
			if item.Id == *item_id {
				cart.Items[idx].Quantity = quantity
				return nil
			}
		}
	}

	if item_id != nil {
		return errors.New("Item ID not found")
	}

	if quantity < 0 {
		return errors.New("Quantity for new item should bigger than 0")
	}

	if product_id != nil {
		cart.Items = append(cart.Items,
			CartItem{
				ProductId: *product_id,
				Quantity:  quantity,
			})
	}

	return nil
}

func (cart *Cart) CalculateAmount() {
	for _, item := range cart.Items {
		cart.Amounts.Subtotal += uint(item.ProductPrice) * uint(item.Quantity)
	}
	cart.Amounts.Shipping = 0
	cart.Amounts.Total = cart.Amounts.Shipping + cart.Amounts.Subtotal
}

func (cart *Cart) EraseAccessToken() {
	cart.AccessToken = ""
}

func (cart *Cart) RemoveZeroQuantityItems() {
	for idx, _ := range cart.Items {
		if cart.Items[idx].Quantity <= 0 {
			cart.Items = append(cart.Items[:idx], cart.Items[idx+1:]...)
			return
		}
	}
}
