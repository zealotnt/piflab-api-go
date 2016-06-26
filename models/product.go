package models

import (
	"regexp"
	"strconv"
	"time"

	. "github.com/o0khoiclub0o/piflab-store-api-go/services"
)

type Product struct {
	Id                 uint      `json:"id"`
	Name               string    `json:"name"`
	Price              int       `json:"price"`
	Provider           string    `json:"provider"`
	Rating             float32   `json:"rating"`
	Status             string    `json:"status"`
	Detail             string    `json:"detail"`
	ImageData          []byte    `json:"-" sql:"-"`
	ImageThumbnailData []byte    `json:"-" sql:"-"`
	ImageDetailData    []byte    `json:"-" sql:"-"`
	Image              string    `json:"-"`
	ImageUpdatedAt     time.Time `json:"-"`
	ImageUrl           *string   `json:"image_url" sql:"-"`
	ImageThumbnailUrl  *string   `json:"image_thumbnail_url" sql:"-"`
	ImageDetailUrl     *string   `json:"image_detail_url" sql:"-"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
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

type ImageType int

const (
	ORIGIN ImageType = iota
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
	next := "/products/offset=" + strconv.FormatUint(nextNum, 10) + "&limit=" + strconv.FormatUint(uint64(limit), 10)
	previous := "/products/offset=" + strconv.FormatUint(prevNum, 10) + "&limit=" + strconv.FormatUint(uint64(limit), 10)

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

func (product *Product) GetImagePath(image ImageType) string {
	var prefix string
	var extension string

	switch image {
	case ORIGIN:
		prefix = "/origin."
		re, _ := regexp.Compile(`.+(\..+$)`)
		if res := re.FindStringSubmatch(product.Image); res != nil {
			extension = res[1]
		}
	case THUMBNAIL:
		prefix = "/thumbnail."
		extension = ".png"
	case DETAIL:
		prefix = "/detail."
		extension = ".png"
	default:
		return ""
	}

	if extension != "" {
		return "products/" + strconv.FormatUint(uint64(product.Id), 10) + prefix + strconv.FormatInt(product.ImageUpdatedAt.Unix(), 10) + extension
	}

	return "products/" + strconv.FormatUint(uint64(product.Id), 10) + prefix + strconv.FormatInt(product.ImageUpdatedAt.Unix(), 10)

}

func (product *Product) GetImageUrlType(image ImageType) (string, error) {
	return (FileService{}).GetProtectedUrl(product.GetImagePath(image), 15)
}

func (product *Product) GetImageUrl() error {
	imageTypeList := [3]ImageType{ORIGIN, THUMBNAIL, DETAIL}
	urlResult := [3]string{}

	for idx, _ := range imageTypeList {
		var err error
		if urlResult[idx], err = product.GetImageUrlType(imageTypeList[idx]); err != nil {
			return err
		}
	}
	product.ImageUrl = &urlResult[0]
	product.ImageThumbnailUrl = &urlResult[1]
	product.ImageDetailUrl = &urlResult[2]

	return nil
}
