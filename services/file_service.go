package services

import (
	"bytes"
	"errors"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileService struct {
}

func (service FileService) SaveFile(data []byte, key string) error {
	if data == nil {
		return errors.New("File content is required")
	} else if key == "" {
		return errors.New("Key is required")
	}

	uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))

	_, err := uploader.Upload(&s3manager.UploadInput{
		Body:        bytes.NewReader(data),
		Bucket:      aws.String((os.Getenv("S3_BUCKET_NAME"))),
		Key:         aws.String(key),
		ContentType: aws.String("image"),
	})
	return err
}

func (service FileService) GetFile(key string) ([]byte, error) {
	downloader := s3manager.NewDownloader(session.New(&aws.Config{
		Region: aws.String(os.Getenv("S3_REGION")),
	}))

	var aws_buff aws.WriteAtBuffer
	_, err := downloader.Download(&aws_buff,
		&s3.GetObjectInput{
			Bucket: aws.String((os.Getenv("S3_BUCKET_NAME"))),
			Key:    aws.String(key),
		})

	return aws_buff.Bytes(), err
}

func (service FileService) DeleteFile(key string) error {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))
	params := &s3.DeleteObjectsInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Delete: &s3.Delete{
			Objects: []*s3.ObjectIdentifier{
				{
					Key: aws.String(key),
				},
			},
			Quiet: aws.Bool(true),
		},
	}

	_, err := svc.DeleteObjects(params)

	return err
}

func (servie FileService) GetProtectedUrl(key string, minute int) (string, error) {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(key),
	})
	return req.Presign(time.Duration(minute) * time.Minute)
}
