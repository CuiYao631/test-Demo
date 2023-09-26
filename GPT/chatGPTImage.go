package main

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	key := "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"
	c := openai.NewClient(key)
	ctx := context.Background()

	// Sample image by link
	reqUrl := openai.ImageRequest{
		Prompt:         "A cat with a dog's head",
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              3,
	}

	respUrl, err := c.CreateImage(ctx, reqUrl)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}
	//fmt.Println(respUrl.Data[0].URL)
	for _, v := range respUrl.Data {
		fmt.Println(v.URL)
	}

	// Example image as base64
	//reqBase64 := openai.ImageRequest{
	//	Prompt:         "蓝天、白云和大海",
	//	Size:           openai.CreateImageSize256x256,
	//	ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	//	N:              1,
	//}
	//
	//respBase64, err := c.CreateImage(ctx, reqBase64)
	//if err != nil {
	//	fmt.Printf("Image creation error: %v\n", err)
	//	return
	//}
	//
	//imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	//if err != nil {
	//	fmt.Printf("Base64 decode error: %v\n", err)
	//	return
	//}
	//
	//r := bytes.NewReader(imgBytes)
	//imgData, err := png.Decode(r)
	//if err != nil {
	//	fmt.Printf("PNG decode error: %v\n", err)
	//	return
	//}
	//
	//file, err := os.Create("example.png")
	//if err != nil {
	//	fmt.Printf("File creation error: %v\n", err)
	//	return
	//}
	//defer file.Close()
	//
	//if err := png.Encode(file, imgData); err != nil {
	//	fmt.Printf("PNG encode error: %v\n", err)
	//	return
	//}
	//
	//fmt.Println("The image was saved as example.png")
}
