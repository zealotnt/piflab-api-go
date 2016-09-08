package repository

import (
	"github.com/icrowley/fake"
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"

	"errors"
	"math/rand"
	"time"
)

type OrderRepository struct {
	*DB
}

func (repo OrderRepository) generateAccessToken(order *Order) error {
	rand.Seed(time.Now().UTC().UnixNano())

try_gen_other_value:
	order.AccessToken = fake.CharactersN(32)

	temp_order := &Order{}
	if err := repo.DB.Where("access_token = ?", order.AccessToken).Find(temp_order).Error; err != nil {
		// Check if err is not found -> access_token is unique
		if err.Error() == "record not found" {
			return nil
		}

		// Otherwise, this is database operation error
		return errors.New("Database error")
	}

	// duplicate, try again
	goto try_gen_other_value
}

func (repo OrderRepository) getOrderItemsInfo(order *Order) {
	for idx, item := range order.Items {
		product := &Product{}
		product.Id = item.ProductId
		repo.DB.Select("name, price, image, image_updated_at").Find(&product)
		order.Items[idx].ProductPrice = product.Price
		order.Items[idx].ProductName = product.Name
		(*product).GetImageUrl()
		order.Items[idx].ProductImageThumbnailUrl = product.ImageThumbnailUrl
	}
}

func (repo OrderRepository) clearNullQuantity() {
	repo.DB.Delete(OrderItem{}, "quantity=0")
}

func (repo OrderRepository) createOrder(order *Order) error {
	if err := repo.generateAccessToken(order); err != nil {
		return err
	}

	if err := repo.DB.Create(order).Error; err != nil {
		return err
	}

	repo.getOrderItemsInfo(order)

	return nil
}

func (repo OrderRepository) updateOrder(order *Order) error {
	tx := repo.DB.Begin()

	if err := tx.Save(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	repo.clearNullQuantity()

	// Don't return access_token when updating
	order.EraseAccessToken()

	repo.getOrderItemsInfo(order)

	return nil
}

func (repo OrderRepository) GetOrder(access_token string) (*Order, error) {
	order := &Order{}
	items := &[]OrderItem{}

	// find a order by its access_token
	if err := repo.DB.Where("access_token = ?", access_token).Find(order).Error; err != nil {
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

func (repo OrderRepository) SaveOrder(order *Order) error {
	if order.AccessToken == "" {
		return repo.createOrder(order)
	}
	return repo.updateOrder(order)
}

func (repo OrderRepository) DeleteOrderItem(order *Order, item_id uint) error {
	item := OrderItem{}

	// use order.Id to find its OrderItem data (order.Id is its forein key)
	if err := repo.DB.Where("id = ? AND order_id = ?", item_id, order.Id).Find(&item).Error; err != nil {
		if err.Error() == "record not found" {
			return errors.New("Not found Item Id in Order")
		}

		return err
	}

	repo.DB.Delete(&item)

	// use order.Id to find its OrderItem data (order.Id is its forein key)
	items := &[]OrderItem{}
	repo.DB.Where("order_id = ?", order.Id).Find(items)
	order.Items = *items

	repo.getOrderItemsInfo(order)

	return nil
}

func (repo OrderRepository) CheckoutOrder(order *Order) error {

	return nil
}
