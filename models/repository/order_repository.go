package repository

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"

	"encoding/json"
	"fmt"
	"strconv"
)

type OrderRepository struct {
	*App
}

func (repo OrderRepository) getOrderItemsInfo(order_items []OrderItem, get_product_price_name bool) error {
	for idx, item := range order_items {
		product := &Product{}
		var err error

		product.Id = item.ProductId
		product, err = (ProductRepository{repo.App}).FindById(product.Id)
		if err != nil {
			return fmt.Errorf("Product %v", err)
		}

		// This option is for cart/checkout
		// + when cart, we will update the product price and name whenever there is a change
		// + when checkout, we will not fetch the product price and name, it is stored in the order's db table
		if get_product_price_name == true {
			order_items[idx].ProductPrice = product.Price
			order_items[idx].ProductName = product.Name
		}

		order_items[idx].ProductImageThumbnailUrl = product.ImageThumbnailUrl
	}

	return nil
}

func (repo OrderRepository) createOrder(order *Order) error {
	type CreateCartForm struct {
		Product_Id uint   `json:"product_id"`
		Quantity   int    `json:"quantity"`
		Name       string `json:"name"`
		Price      uint   `json:"price"`
	}
	form := new(CreateCartForm)

	if err := repo.getOrderItemsInfo(order.Items, true); err != nil {
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
	repo.getOrderItemsInfo(order.Items, true)

	return nil
}

func (repo OrderRepository) updateOrder(order *Order) error {
	type UpdateCartItemForm struct {
		AccessToken string `json:"access_token"`
		Quantity    int    `json:"quantity"`
		Name        string `json:"name"`
		Price       uint   `json:"price"`
	}
	type UpdateCheckoutStatusForm struct {
		Status string `json:"status"`
	}
	type CreateCartItemForm struct {
		AccessToken string `json:"access_token"`
		Product_Id  uint   `json:"product_id"`
		Quantity    int    `json:"quantity"`
		Name        string `json:"name"`
		Price       uint   `json:"price"`
	}

	if err := repo.getOrderItemsInfo(order.Items, true); err != nil {
		return err
	}

	if order.ItemUpdateId == 0 {
		form := new(CreateCartItemForm)

		if order.StatusUpdated == true {
			form := new(UpdateCheckoutStatusForm)

			form.Status = order.Status
			response, body := repo.App.HttpRequest("PUT",
				repo.ORDER_SERVICE+"/orders/"+order.OrderCodeRet,
				form)
			if response.Status != "200 OK" {
				return ParseError(body)
			}

			if err := json.Unmarshal([]byte(body), order); err != nil {
				return err
			}
		} else {
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

	repo.getOrderItemsInfo(order.Items, true)

	return nil
}

func (repo OrderRepository) FindByOrderCode(order_code string) (*Order, error) {
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

	// Update the image url (not the price & name)
	repo.getOrderItemsInfo(order.Items, false)

	return order, nil
}

func (repo OrderRepository) GetOrder(access_token string) (*Order, error) {
	order := &Order{}
	response, body := repo.App.HttpRequest("GET", repo.ORDER_SERVICE+"/cart?access_token="+access_token, nil)
	if response.Status != "200 OK" {
		return nil, ParseError(body)
	}

	if err := json.Unmarshal([]byte(body), &order); err != nil {
		return nil, err
	}
	repo.getOrderItemsInfo(order.Items, true)

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

func (repo OrderRepository) GetCheckoutPage(offset uint, limit uint, status string, sort_field string, sort_order string, search string) (*OrderPage, error) {
	order_page := &OrderPage{}
	var err error
	var query_param string

	query_param += "?offset=" + strconv.Itoa(int(offset))
	query_param += "&limit=" + strconv.Itoa(int(limit))
	if status != "" {
		query_param += "&status" + status
	}
	query_param += "&sort=" + sort_field + "|" + sort_order
	if search != "" {
		query_param += "&q=" + search
	}

	response, body := repo.App.HttpRequest("GET", repo.ORDER_SERVICE+"/orders"+query_param, nil)
	if response.Status != "200 OK" {
		return nil, ParseError(body)
	}

	if err := json.Unmarshal([]byte(body), &order_page); err != nil {
		return nil, err
	}

	// Update the image url (not the price & name)
	for idx, _ := range *order_page.Data {
		repo.getOrderItemsInfo((*order_page.Data)[idx].Items, false)
	}

	return order_page, err
}

func (repo OrderRepository) CheckoutOrder(order *Order) error {
	type GatewayCheckoutCartForm struct {
		AccessToken     string `json:"access_token"`
		CustomerName    string `json:"name"`
		CustomerAddress string `json:"address"`
		CustomerPhone   string `json:"phone"`
		CustomerEmail   string `json:"email"`
		CustomerNote    string `json:"note"`
	}

	form := new(GatewayCheckoutCartForm)

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
