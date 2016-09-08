package models

import (
	"errors"
	"time"
)

type Amount struct {
	Subtotal uint `json:"subtotal"`
	Shipping uint `json:"shipping"`
	Total    uint `json:"total"`
}

type OrderInfo struct {
	CustomerName    string `sql:"customer_name"`
	CustomerAddress string `sql:"customer_address"`
	CustomerPhone   string `sql:"customer_phone"`
	CustomerEmail   string `sql:"customer_email"`
	CustomerNote    string `sql:"customer_note"`
	OrderCode       string `sql:"order_code"`
}

type Order struct {
	Id          uint   `json:"-"`
	AccessToken string `json:"access_token,omitempty"`
	Status      string `json:"-"`

	Items []OrderItem `json:"items" sql:"order_items"`

	OrderInfo `json:"-"`

	Amounts Amount `json:"amounts" sql:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	Id                       uint    `json:"id" sql:"id"`
	OrderId                  uint    `json:"-" sql:"REFERENCES Orders(id)"`
	ProductId                uint    `json:"product_id" sql:"REFERENCES products(id)"`
	ProductName              string  `json:"name" sql:"-"`
	ProductImageThumbnailUrl *string `json:"image_thumbnail_url" sql:"-"`
	ProductPrice             int     `json:"price" sql:"-"`
	Quantity                 int     `json:"quantity"`
}

func (Order *Order) UpdateItems(product_id *uint, item_id *uint, quantity int) error {
	for idx, item := range Order.Items {
		if product_id != nil {
			if item.ProductId == *product_id {
				Order.Items[idx].Quantity += quantity
				if Order.Items[idx].Quantity < 0 {
					Order.Items[idx].Quantity = 0
				}
				return nil
			}
		}
		if item_id != nil {
			if item.Id == *item_id {
				Order.Items[idx].Quantity = quantity
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
		Order.Items = append(Order.Items,
			OrderItem{
				ProductId: *product_id,
				Quantity:  quantity,
			})
	}

	return nil
}

func (Order *Order) CalculateAmount() {
	for _, item := range Order.Items {
		Order.Amounts.Subtotal += uint(item.ProductPrice) * uint(item.Quantity)
	}
	Order.Amounts.Shipping = 0
	Order.Amounts.Total = Order.Amounts.Shipping + Order.Amounts.Subtotal
}

func (Order *Order) EraseAccessToken() {
	Order.AccessToken = ""
}

func (Order *Order) RemoveZeroQuantityItems() {
	for idx, _ := range Order.Items {
		if Order.Items[idx].Quantity <= 0 {
			Order.Items = append(Order.Items[:idx], Order.Items[idx+1:]...)
			return
		}
	}
}
