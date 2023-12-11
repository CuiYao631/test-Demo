package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

//cloudflare
func main() {
	ctx := context.Background()
	var bucketName = "oetech"
	var accountId = "1ff790aacdc24f86de466ee7335b357f"
	var accessKeyId = "962947d12840d1d9249dafe44fed94a8"
	var accessKeySecret = "cbc07b7fbba1071dc440289bb7be04eff3aabd4bb4f16d5854fd5dff61e48797"

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
			HostnameImmutable: true,
			Source:            aws.EndpointSourceCustom,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	listObjectsOutput, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Println("bucketName err", err)
		log.Fatal(err)
	}
	log.Println(listObjectsOutput.Contents)

	for _, object := range listObjectsOutput.Contents {
		obj, _ := json.MarshalIndent(object, "", "\t")
		fmt.Println(string(obj))
	}

	presignClient := s3.NewPresignClient(client)

	presignResult, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("example.txt"),
	})

	if err != nil {
		panic("Couldn't get presigned URL for PutObject")
	}
	fmt.Printf("Presigned URL For object: %s\n", presignResult.URL)
}
