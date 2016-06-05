package services

import (
	"mime/multipart"
	"net/http"
)

type ImageService struct {
}

type ImageFile interface {
	Open() (multipart.File, error)
}

func (service ImageService) IsValidImage(image ImageFile) bool {
	fh, err := image.Open()
	if err != nil {
		return false
	}
	defer fh.Close()

	buff := make([]byte, 512)
	if _, err := fh.Read(buff); err != nil {
		return false
	}

	filetype := http.DetectContentType(buff)
	switch filetype {
	case "image/jpeg":
		fallthrough
	case "image/png":
		fallthrough
	case "image/gif":
		return true
	}

	return false
}
