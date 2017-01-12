package models

import (
	"regexp"
	"strconv"
	"time"

	. "github.com/o0khoiclub0o/piflab-store-api-go/services"
)

type Product struct {
	Id       uint    `json:"id,string"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	Provider string  `json:"provider"`
	Rating   float32 `json:"rating"`
	Status   string  `json:"status"`
	Detail   string  `json:"detail"`

	ImageData          []byte    `json:"-" sql:"-"`
	ImageThumbnailData []byte    `json:"-" sql:"-"`
	ImageDetailData    []byte    `json:"-" sql:"-"`
	Image              string    `json:"-"`
	NewImage           string    `json:"-" sql:"-"`
	ImageUpdatedAt     time.Time `json:"-"`
	ImageUrl           *string   `json:"image_url" sql:"-"`
	ImageThumbnailUrl  *string   `json:"image_thumbnail_url" sql:"-"`
	ImageDetailUrl     *string   `json:"image_detail_url" sql:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductSlice []Product

type PageUrl struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type ProductPage struct {
	Data   *ProductSlice `json:"data"`
	Paging PageUrl       `json:"paging"`
}

type ImageField int

const (
	IMAGE ImageField = iota
)

type ImageSize int

const (
	ORIGIN ImageSize = iota
	THUMBNAIL
	DETAIL
)

func getPage(offset uint, limit uint, total uint) PageUrl {
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
	next := "/products?offset=" + strconv.FormatUint(nextNum, 10) + "&limit=" + strconv.FormatUint(uint64(limit), 10)
	previous := "/products?offset=" + strconv.FormatUint(prevNum, 10) + "&limit=" + strconv.FormatUint(uint64(limit), 10)

	if uint64(total) <= nextNum {
		return PageUrl{
			Previous: &previous,
		}
	}
	if offset == 0 {
		return PageUrl{
			Next: &next,
		}
	}
	return PageUrl{
		Next:     &next,
		Previous: &previous,
	}

}

func (products ProductSlice) GetPaging(offset uint, limit uint, total uint) *ProductPage {
	return &ProductPage{
		Data:   &products,
		Paging: getPage(offset, limit, total),
	}
}

func (product *Product) GetImagePath(field ImageField, image ImageSize) string {
	var img_size string
	var extension string
	var img_field string
	var img_name string
	var img_updated_at string

	switch field {
	case IMAGE:
		img_field = "/image_"
		img_name = product.Image
		img_updated_at = strconv.FormatInt(product.ImageUpdatedAt.Unix(), 10)
	default:
		return ""
	}

	switch image {
	case ORIGIN:
		img_size = "origin_"
		re, _ := regexp.Compile(`.+(\..+$)`)
		if res := re.FindStringSubmatch(img_name); res != nil {
			extension = res[1]
		}
	case THUMBNAIL:
		img_size = "thumbnail_"
		extension = ".png"
	case DETAIL:
		img_size = "detail_"
		extension = ".png"
	default:
		return ""
	}

	if extension != "" {
		return "products/" + strconv.FormatUint(uint64(product.Id), 10) + img_field + img_size + img_updated_at + extension
	}

	return "products/" + strconv.FormatUint(uint64(product.Id), 10) + img_field + img_size + img_updated_at
}

func (product *Product) GetImageContentType(field ImageField, image ImageSize) string {
	var extension string
	var img_name string

	switch field {
	case IMAGE:
		img_name = product.Image
	default:
		return ""
	}

	switch image {
	case ORIGIN:
		re, _ := regexp.Compile(`.+\.(.+$)`)
		if res := re.FindStringSubmatch(img_name); res != nil {
			extension = res[1]
		} else {
			return "image"
		}
	case THUMBNAIL:
		extension = "png"
	case DETAIL:
		extension = "png"
	default:
		return ""
	}

	return "image/" + extension
}

func (product *Product) GetImageUrlType(field ImageField, image ImageSize) (string, error) {
	return (FileService{}).GetProtectedUrl(product.GetImagePath(field, image), 15)
}

func (product *Product) GetImageUrl() {
	imageSizeList := [3]ImageSize{ORIGIN, THUMBNAIL, DETAIL}
	urlResult := [3]string{}

	if product.Image == "" {
		return
	}

	for idx, _ := range imageSizeList {
		urlResult[idx], _ = product.GetImageUrlType(IMAGE, imageSizeList[idx])
	}
	product.ImageUrl = &urlResult[0]
	product.ImageThumbnailUrl = &urlResult[1]
	product.ImageDetailUrl = &urlResult[2]
}
