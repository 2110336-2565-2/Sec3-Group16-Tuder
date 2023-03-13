package utils

import (
	"context"
	"fmt"
	"bytes"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

func GenerateProfilePictureURL(imageBytes []byte, key string) (string, error) {
	bucket := "se2-tuder"

	// Load AWS session configuration
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to load AWS SDK configuration: %v", err)
	}

	// Create a new S3 client
	client := s3.NewFromConfig(cfg)

	// Set up a new S3 uploader
	uploader := manager.NewUploader(client)

	// Upload the image to S3
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   bytes.NewReader(imageBytes),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image to S3: %v", err)
	}

	// Generate the S3 object URL
	objectURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)

	return objectURL, nil
}
