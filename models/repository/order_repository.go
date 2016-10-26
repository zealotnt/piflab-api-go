package repository

import (
	"github.com/icrowley/fake"
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"

	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type OrderRepository struct {
	*App
}

func (repo OrderRepository) generateOrderCode(order *Order) error {
	rand.Seed(time.Now().UTC().UnixNano())

try_gen_other_value:
	order.OrderInfo.OrderCode = fake.CharactersN(32)

	temp_order := &Order{}
	if err := repo.DB.Where("code = ?", order.OrderInfo.OrderCode).Find(temp_order).Error; err != nil {
		// Check if err is not found -> code is unique
		if err.Error() == "record not found" {
			return nil
		}

		// Otherwise, this is database operation error
		return errors.New("Database error")
	}

	// duplicate, try again
	goto try_gen_other_value
}

func (repo OrderRepository) getOrderItemsInfo(order *Order) error {
	for idx, item := range order.Items {
		product := &Product{}
		var err error

		product.Id = item.ProductId
		product, err = (ProductRepository{repo.App}).FindById(product.Id)
		if err != nil {
			return fmt.Errorf("Product %v", err)
		}

		order.Items[idx].ProductPrice = product.Price
		order.Items[idx].ProductName = product.Name
		order.Items[idx].ProductImageThumbnailUrl = product.ImageThumbnailUrl
		return nil
	}

	return nil
}

func (repo OrderRepository) clearNullQuantity() {
	repo.DB.Delete(OrderItem{}, "quantity=0")
}

func (repo OrderRepository) createOrder(order *Order) error {
	type CreateCartForm struct {
		Product_Id uint   `json:"product_id"`
		Quantity   int    `json:"quantity"`
		Name       string `json:"name"`
		Price      uint   `json:"price"`
	}
	form := new(CreateCartForm)

	if err := repo.getOrderItemsInfo(order); err != nil {
		return err
	}

	form.Product_Id = order.Items[0].ProductId
	form.Quantity = order.Items[0].Quantity
	form.Price = uint(order.Items[0].ProductPrice)
	form.Name = order.Items[0].ProductName
	response, body := repo.App.HttpRequest("PUT",
		repo.ORDER_SERVICE+"/cart/items",
		form)
	if response.Status != "200 OK" {
		return ParseError(body)
	}

	if err := json.Unmarshal([]byte(body), order); err != nil {
		return err
	}
	repo.getOrderItemsInfo(order)

	return nil
}

func (repo OrderRepository) updateOrder(order *Order) error {
	type UpdateCartItemForm struct {
		AccessToken string `json:"access_token"`
		Quantity    int    `json:"quantity"`
		Name        string `json:"name"`
		Price       uint   `json:"price"`
	}
	type CreateCartItemForm struct {
		AccessToken string `json:"access_token"`
		Product_Id  uint   `json:"product_id"`
		Quantity    int    `json:"quantity"`
		Name        string `json:"name"`
		Price       uint   `json:"price"`
	}

	if err := repo.getOrderItemsInfo(order); err != nil {
		return err
	}

	if order.ItemUpdateId == 0 {
		form := new(CreateCartItemForm)

		if order.ItemUpdateNew == true {
			// the brand-new item, update with {product_id, quantity}
			num_of_item := len(order.Items)
			form.Quantity = order.ItemUpdateQuantity
			form.AccessToken = order.AccessToken
			form.Product_Id = order.Items[num_of_item-1].ProductId
			form.Quantity = order.Items[num_of_item-1].Quantity
			form.Price = uint(order.Items[num_of_item-1].ProductPrice)
			form.Name = order.Items[num_of_item-1].ProductName
		} else {
			// update quantity base on offset {product_id, quantity}
			form.Quantity = order.ItemUpdateQuantity
			form.AccessToken = order.AccessToken
			form.Product_Id = uint(order.Items[order.ItemUpdateIdx].ProductId)
			form.Price = uint(order.Items[order.ItemUpdateIdx].ProductPrice)
			form.Name = order.Items[order.ItemUpdateIdx].ProductName
		}

		response, body := repo.App.HttpRequest("PUT",
			repo.ORDER_SERVICE+"/cart/items",
			form)
		if response.Status != "200 OK" {
			return ParseError(body)
		}

		if err := json.Unmarshal([]byte(body), order); err != nil {
			return err
		}
	} else {
		form := new(UpdateCartItemForm)

		// update quantity base on offset {item_id, quantity}
		form.Quantity = order.ItemUpdateQuantity
		form.AccessToken = order.AccessToken
		form.Price = uint(order.Items[order.ItemUpdateIdx].ProductPrice)
		form.Name = order.Items[order.ItemUpdateIdx].ProductName

		response, body := repo.App.HttpRequest("PUT",
			repo.ORDER_SERVICE+"/cart/items/"+strconv.Itoa(order.ItemUpdateId),
			form)
		if response.Status != "200 OK" {
			return ParseError(body)
		}

		if err := json.Unmarshal([]byte(body), order); err != nil {
			return err
		}
	}

	repo.getOrderItemsInfo(order)

	return nil
}

func (repo OrderRepository) FindByOrderId(order_code string) (*Order, error) {
	order := &Order{}
	response, body := repo.App.HttpRequest("GET",
		repo.ORDER_SERVICE+"/orders/"+order_code,
		nil)
	if response.Status != "200 OK" {
		return nil, ParseError(body)
	}

	if err := json.Unmarshal([]byte(body), order); err != nil {
		return nil, err
	}

	return order, nil
}

func (repo OrderRepository) GetOrderByOrdercode(order_code string) (*Order, error) {
	order := &Order{}
	items := &[]OrderItem{}

	// find a order by its order_code
	if err := repo.DB.Where("code = ?", order_code).Find(order).Error; err != nil {
		return nil, err
	}

	// use order.Id to find its OrderItem data (order.Id is its forein key)
	if err := repo.DB.Where("order_id = ?", order.Id).Find(items).Error; err != nil {
		return nil, err
	}

	// use the order.Items to update products information
	order.Items = *items
	repo.getOrderItemsInfo(order)

	return order, nil
}

func (repo OrderRepository) GetOrder(access_token string) (*Order, error) {
	order := &Order{}
	response, body := repo.App.HttpRequest("GET", repo.ORDER_SERVICE+"/cart?access_token="+access_token, "")
	if response.Status != "200 OK" {
		return nil, ParseError(body)
	}

	if err := json.Unmarshal([]byte(body), &order); err != nil {
		return nil, err
	}
	repo.getOrderItemsInfo(order)

	return order, nil
}

func (repo OrderRepository) SaveOrder(order *Order) error {
	if order.AccessToken == "" {
		return repo.createOrder(order)
	}
	return repo.updateOrder(order)
}

func (repo OrderRepository) DeleteOrderItem(order *Order, item_id uint) error {
	response, body := repo.App.HttpRequest("DELETE", repo.ORDER_SERVICE+"/cart/items/"+strconv.Itoa(int(item_id))+"?access_token="+order.AccessToken, "")
	if response.Status != "200 OK" {
		return ParseError(body)
	}

	if err := json.Unmarshal([]byte(body), &order); err != nil {
		return err
	}

	return nil
}

func (repo OrderRepository) CountOrders() (uint, error) {
	count := uint(0)

	err := repo.DB.Table("orders").Count(&count).Error

	return count, err
}

func (repo OrderRepository) GetPage(offset uint, limit uint, status string, sort_field string, sort_order string, search string) (*OrderSlice, uint, error) {
	orders := &OrderSlice{}
	items := &[]OrderItem{}
	var err error
	var where_param string

	if status == "" {
		where_param = "status!='cart'"
	} else {
		where_param = "status='" + status + "'"
	}

	if search != "" {
		where_param += " AND LOWER(customer_name) LIKE  '%" + strings.ToLower(search) + "%'"
	}

	err = repo.DB.Order(sort_field + " " + sort_order).Offset(int(offset)).Where(where_param).Limit(int(limit)).Find(orders).Error

	for idx, order := range *orders {
		// use order.Id to find its OrderItem data (order.Id is its forein key)
		if err := repo.DB.Where("order_id = ?", order.Id).Find(items).Error; err != nil {
			return nil, 0, err
		}
		// use the order.Items to update products information
		(*orders)[idx].Items = *items
		repo.getOrderItemsInfo(&(*orders)[idx])
		(*orders)[idx].CalculateAmount()
	}

	count, _ := repo.CountOrders()

	return orders, count, err
}

func (repo OrderRepository) CheckoutOrder(order *Order) error {
	type CheckoutCartForm struct {
		AccessToken     string `json:"access_token"`
		CustomerName    string `json:"name"`
		CustomerAddress string `json:"address"`
		CustomerPhone   string `json:"phone"`
		CustomerEmail   string `json:"email"`
		CustomerNote    string `json:"note"`
	}

	form := new(CheckoutCartForm)

	form.AccessToken = order.AccessToken
	form.CustomerName = order.OrderInfo.CustomerName
	form.CustomerAddress = order.OrderInfo.CustomerAddress
	form.CustomerPhone = order.OrderInfo.CustomerPhone
	form.CustomerEmail = order.OrderInfo.CustomerEmail
	form.CustomerNote = order.OrderInfo.CustomerNote

	response, body := repo.App.HttpRequest("POST",
		repo.ORDER_SERVICE+"/cart/checkout",
		form)
	if response.Status != "200 OK" {
		return ParseError(body)
	}

	if err := json.Unmarshal([]byte(body), order); err != nil {
		return err
	}

	return nil
}
