package s3

import (
	"bytes"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3manager"
	"github.com/lolmourne/go-accounts/model"
	uuid "github.com/satori/go.uuid"
)

type IS3 interface {
	Put(data []byte) error
}

type S3Resource struct {
	uploadClient *s3manager.Uploader
	bucketName   string
}

func NewS3Resource(cfg model.Config) IS3 {
	awsCfg := &aws.Config{
		Region: aws.String("ap-southeast-1"),
	}
	awsCfg.WithCredentials(credentials.NewCredentials(&credentials.StaticProvider{
		Value: credentials.Value{
			AccessKeyID:     cfg.S3Cred.AccessID,
			SecretAccessKey: cfg.S3Cred.Secret,
		},
	}))

	sess := session.New(awsCfg)

	return &S3Resource{
		uploadClient: s3manager.NewDownloader(sess),
		bucketName:   cfg.S3Cred.BucketName,
	}
}

func (s3 *S3Resource) Put(data []byte) error {
	upInput := &s3manager.UploadInput{
		Bucket: aws.String(s3.bucketName),
		Key:    aws.String(GenerateUUIDv4()),
		Body:   bytes.NewReader(data),
	}

	_, err := s3.uploadClient.Upload(upInput)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GenerateUUIDv4() string {
	return uuid.Must(uuid.NewV4()).String()
}
