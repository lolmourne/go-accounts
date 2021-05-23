package s3

import (
	"bytes"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/lolmourne/go-accounts/model"
)

type IS3 interface {
	Put(data []byte, filename string, mime string) error
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

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		log.Fatal("Cannot connect AWS S3 ", err)
	}

	return &S3Resource{
		uploadClient: s3manager.NewUploader(sess),
		bucketName:   cfg.S3Cred.BucketName,
	}
}

func (s3 *S3Resource) Put(data []byte, filename string, mime string) error {

	upInput := &s3manager.UploadInput{
		Bucket:      aws.String(s3.bucketName),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(mime),
	}

	_, err := s3.uploadClient.Upload(upInput)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
