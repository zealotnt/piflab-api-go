package models

import (
	"errors"
	"strconv"
	"time"
)

type Amount struct {
	Subtotal uint `json:"subtotal"`
	Shipping uint `json:"shipping"`
	Total    uint `json:"total"`
}

type OrderInfo struct {
	OrderCode       string `json:"id" sql:"order_code"`
	CustomerName    string `json:"name" sql:"customer_name"`
	CustomerAddress string `json:"address" sql:"customer_address"`
	CustomerPhone   string `json:"phone" sql:"customer_phone"`
	CustomerEmail   string `json:"email" sql:"customer_email"`
	CustomerNote    string `json:"note" sql:"customer_note"`
}

type CheckoutReturn struct {
	Items     []OrderItem `json:"items"`
	Amounts   Amount      `json:"amounts" sql:"-"`
	OrderInfo OrderInfo   `json:"customer" sql:"-"`
	Status    string      `json:"status"`
}

type OrderSlice []Order

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

type OrderUrl struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type OrderPage struct {
	Data   *OrderSlice `json:"data"`
	Paging OrderUrl    `json:"paging"`
}

func getOrderPage(offset uint, limit uint, total uint) OrderUrl {
	prevNum := uint64(offset - limit)
	nextNum := uint64(offset + limit)
	if offset < limit {
		prevNum = 0
	}
	if total <= offset {
		if total > limit {
			prevNum = uint64(total - limit)
		} else {
			prevNum = 0
		}
	}
	next := "/orders?offset=" + strconv.FormatUint(nextNum, 10) + "&limit=" + strconv.FormatUint(uint64(limit), 10)
	previous := "/orders?offset=" + strconv.FormatUint(prevNum, 10) + "&limit=" + strconv.FormatUint(uint64(limit), 10)

	if uint64(total) <= nextNum {
		return OrderUrl{
			Previous: &previous,
		}
	}
	if offset == 0 {
		return OrderUrl{
			Next: &next,
		}
	}
	return OrderUrl{
		Next:     &next,
		Previous: &previous,
	}

}

func (orders OrderSlice) GetPaging(offset uint, limit uint, total uint) *OrderPage {
	return &OrderPage{
		Data:   &orders,
		Paging: getOrderPage(offset, limit, total),
	}
}

func (order *Order) UpdateItems(product_id *uint, item_id *uint, quantity int) error {
	for idx, item := range order.Items {
		if product_id != nil {
			if item.ProductId == *product_id {
				order.Items[idx].Quantity += quantity
				if order.Items[idx].Quantity < 0 {
					order.Items[idx].Quantity = 0
				}
				return nil
			}
		}
		if item_id != nil {
			if item.Id == *item_id {
				order.Items[idx].Quantity = quantity
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
		order.Items = append(order.Items,
			OrderItem{
				ProductId: *product_id,
				Quantity:  quantity,
			})
	}

	return nil
}

func (order *Order) CalculateAmount() {
	for _, item := range order.Items {
		order.Amounts.Subtotal += uint(item.ProductPrice) * uint(item.Quantity)
	}
	order.Amounts.Shipping = 0
	order.Amounts.Total = order.Amounts.Shipping + order.Amounts.Subtotal
}

func (order *Order) EraseAccessToken() {
	order.AccessToken = ""
}

func (order *Order) RemoveZeroQuantityItems() {
	for idx, _ := range order.Items {
		if order.Items[idx].Quantity <= 0 {
			order.Items = append(order.Items[:idx], order.Items[idx+1:]...)
			return
		}
	}
}

func (order *Order) ReturnCheckoutRequest() CheckoutReturn {
	ret := new(CheckoutReturn)
	ret.Items = order.Items
	ret.Amounts = order.Amounts
	ret.OrderInfo = order.OrderInfo
	ret.Status = order.Status
	return *ret
}
