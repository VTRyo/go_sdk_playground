package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"os"
)

func main() {
	var bucket, key string
	var profile string
	flag.StringVar(&profile, "p", "default", "profile name")
	flag.StringVar(&bucket, "b", "sandbox", "bucket name")
	flag.StringVar(&key, "k", "test", "key name")
	flag.Parse()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile)) // context.TODO() こには空の可能性がある他のコンテキストがあるはずだが、適切な値がまだわからないため
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open("./tmp/information_schema.csv")
	if err != nil {
		fmt.Println("Failed to open file", err)
		return
	}
	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	}

	uploader := manager.NewUploader(client)
	_, err = uploader.Upload(context.TODO(), input)
	if err != nil {
		fmt.Println("Failed to upload file", err)
		return
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)

}
