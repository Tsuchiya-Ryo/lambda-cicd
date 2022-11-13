package adaptor

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Region     string
	BucketName string
}

type S3Repository struct {
	Client *s3.Client
}

func NewS3Repository() *S3Repository {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg)
	return &S3Repository{
		Client: client,
	}
}

func (s *S3Repository) PutObject(bucket string, key string, body io.Reader) error {
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	}
	_, err := s.Client.PutObject(context.TODO(), input)
	if err != nil {
		return err
	}
	return nil
}
