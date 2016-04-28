package models

import (
	"regexp"
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
	Image          string    `json:"-"`
	ImageUpdatedAt time.Time `json:"-"`
	ImageUrl       string    `json:"image_url" sql:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (product *Product) GetImagePath() string {
	re, _ := regexp.Compile(`.+(\..+$)`)
	extension := re.FindStringSubmatch(product.Image)
	if extension != nil {
		return "products/" + strconv.FormatUint(uint64(product.Id), 10) + "/origin." + strconv.FormatInt(product.ImageUpdatedAt.Unix(), 10) + extension[1]
	}

	return "products/" + strconv.FormatUint(uint64(product.Id), 10) + "/origin." + strconv.FormatInt(product.ImageUpdatedAt.Unix(), 10)

}

func (product *Product) GetImageUrl() (string, error) {
	return (FileService{}).GetProtectedUrl(product.GetImagePath(), 15)
}
