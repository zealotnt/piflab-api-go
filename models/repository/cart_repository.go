package repository

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"

	"strconv"
)

type CartRepository struct {
	*DB
}

func (repo CartRepository) GetCartItemsInfo(cart *Cart) {
	for idx, item := range cart.Items {
		product := &Product{}
		product.Id = item.ProductId
		repo.DB.Select("name, price, image").Find(&product)
		cart.Items[idx].ProductPrice = product.Price
		cart.Items[idx].ProductName = product.Name
		(*product).GetImageUrl()
		cart.Items[idx].ProductImageThumbnailUrl = product.ImageThumbnailUrl
	}
}

func (repo CartRepository) clearNullQuantity() {
	repo.DB.Delete(CartItem{}, "quantity=0")
}

func (repo CartRepository) createCart(cart *Cart) error {
	if err := repo.DB.Create(cart).Error; err != nil {
		return err
	}
	cart.AccessToken = strconv.FormatUint(uint64(cart.Id), 10)

	err := repo.DB.Save(cart).Error

	repo.GetCartItemsInfo(cart)

	return err
}

func (repo CartRepository) updateCart(cart *Cart) error {
	tx := repo.DB.Begin()

	if err := tx.Save(cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	repo.clearNullQuantity()

	// Don't return access_token when updating
	cart.AccessToken = ""

	return nil
}

func (repo CartRepository) GetCart(access_token string) (*Cart, error) {
	cart := &Cart{}
	items := &[]CartItem{}

	if err := repo.DB.Where("access_token = ?", access_token).Find(cart).Error; err != nil {
		return nil, err
	}

	if err := repo.DB.Where("cart_id = ?", cart.Id).Find(items).Error; err != nil {
		return nil, err
	}

	cart.Items = *items
	repo.GetCartItemsInfo(cart)

	return cart, nil
}

func (repo CartRepository) SaveCart(cart *Cart) error {
	if cart.AccessToken == "" {
		return repo.createCart(cart)
	}
	return repo.updateCart(cart)
}
