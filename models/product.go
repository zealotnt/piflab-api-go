package models

import (
	"errors"
	"strconv"
	"time"

	. "github.com/o0khoiclub0o/piflab-store-api-go/services"
)

type Product struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	Price          int       `json:"price"`
	Provider       string    `json:"provider"`
	Rating         float32   `json:"rating"`
	Status         string    `json:"status"`
	Detail         string    `json:"detail"`
	ImageData      []byte    `json:"-" sql:"-"`
	Image          string    `json:"image"`	
	ImageUpdatedAt time.Time `json:"image_updated_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`	
}

func (product *Product) GetImagePath() string {
	return "products/" + strconv.FormatUint(uint64(product.Id), 10) + "/origin." + strconv.FormatInt(product.ImageUpdatedAt.Unix(), 10)
}

func (product *Product) GetImageUrl(params ...int) (string, error) {
	if params == nil {
		return "", errors.New("Public URL is not support yet")
	}
	return (FileService{}).GetProtectedUrl(product.GetImagePath(), params[0])
}
