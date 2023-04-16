package utils

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func GenerateProfilePictureURL(imageBytes []byte, key string, types string) (string, error) {
	bucket := "se2-tuder"
	err := godotenv.Load()

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
		Bucket: aws.String(bucket),
		Key:    aws.String(types + "/" + key + ".jpg"),
		Body:   bytes.NewReader(imageBytes),
		ACL:    "public-read ",
	})
	if err != nil {
		fmt.Println("failed to upload image to S3", err)
		return "", fmt.Errorf("failed to upload image to S3: %v", err)
	}

	// Generate the S3 object URL
	objectURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s/%s.jpg", bucket, types, key)

	return objectURL, nil
}
