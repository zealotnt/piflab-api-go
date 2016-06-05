package services

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"mime/multipart"

	"github.com/disintegration/imaging"
)

type ImageService struct {
}

type ImageFile interface {
	Open() (multipart.File, error)
}

func (service ImageService) IsValidImage(file ImageFile) (bool, error) {
	fh, err := file.Open()
	if err != nil {
		return false, err
	}
	defer fh.Close()

	image, filetype, err := image.DecodeConfig(fh)
	if image.Width < 550 || image.Height < 550 || err != nil {
		if err != nil {
			return false, err
		}
		return false, errors.New("Image size is too small, Width/Height's minimum value should be 500")
	}

	switch filetype {
	case "jpeg":
		fallthrough
	case "png":
		fallthrough
	case "gif":
		return true, nil
	default:
		return false, errors.New("Image extension is invalid")
	}

	return true, nil
}

func (service ImageService) GetThumbnail(file ImageFile, size int) []byte {
	fh, err := file.Open()
	if err != nil {
		return nil
	}
	defer fh.Close()

	srcImage, _, err := image.Decode(fh)
	if err != nil {
		return nil
	}

	dstImage := imaging.Fit(srcImage, size, size, imaging.Lanczos)

	dataBytes := new(bytes.Buffer)

	if err := png.Encode(dataBytes, dstImage); err != nil {
		return nil
	}

	return dataBytes.Bytes()
}

func (service ImageService) GetDetail(file ImageFile, size int) []byte {
	fh, err := file.Open()
	if err != nil {
		return nil
	}
	defer fh.Close()

	srcImage, _, err := image.Decode(fh)
	if err != nil {
		return nil
	}

	dstImage := imaging.Resize(srcImage, size, 0, imaging.Lanczos)

	dataBytes := new(bytes.Buffer)

	if err := png.Encode(dataBytes, dstImage); err != nil {
		return nil
	}

	return dataBytes.Bytes()
}
