package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"time"
)

//cloudflare
func main() {
	ctx := context.Background()
	var bucketName = "oetech"
	var accountId = "1ff790aacdc24f86de466ee7335b357f//123"
	var accessKeyId = "962947d12840d1d9249dafe44fed94a8//123"
	var accessKeySecret = "cbc07b7fbba1071dc440289bb7be04eff3aabd4bb4f16d5854fd5dff61e48797//123"

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
	currentTime := time.Now()

	// 格式化为RFC3339格式的时间戳字符串
	timeStamp := currentTime.Format("20060102T150405Z")

	fmt.Println(timeStamp)

	client := s3.NewFromConfig(cfg)

	//object, err := client.PutObject(ctx, &s3.PutObjectInput{
	//	Bucket: aws.String(bucketName),
	//	Key:    aws.String("example.txt"),
	//	Body:   bytes.NewReader([]byte("Hello World!")),
	//})
	//if err != nil {
	//	return
	//}
	//log.Println("object", object)

	//listObjectsOutput, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
	//	Bucket: &bucketName,
	//})
	//if err != nil {
	//	log.Println("bucketName err", err)
	//	log.Fatal(err)
	//}
	//log.Println(listObjectsOutput.Contents)
	//
	//for _, object := range listObjectsOutput.Contents {
	//	obj, _ := json.MarshalIndent(object, "", "\t")
	//	fmt.Println(string(obj))
	//}
	//
	presignClient := s3.NewPresignClient(client)

	object, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("smudge-the-viral-cat.webp"),
	})
	if err != nil {
		return
	}
	log.Println("object", object)

	//
	//presignResult, err := presignClient.PresignUploadPart(context.TODO(), &s3.UploadPartInput{
	//	Bucket: aws.String(bucketName),
	//	Key:    aws.String("example.txt"),
	//})
	//
	//if err != nil {
	//	log.Println("presignClient err", err)
	//
	//}
	//log.Println(presignResult.SignedHeader.Get("X-Amz-Date"))

	//log.Println("presignResult", presignResult)
	//fmt.Printf("Presigned URL For object: %s\n", presignResult.URL)
}
