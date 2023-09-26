package main

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func main() {

	//firebaseApp
	opt := option.WithCredentialsFile("/Users/xiaocui/dev-env-374915-firebase-adminsdk-2skau-0308936127.json")
	config := &firebase.Config{

		ProjectID:     "dev-env-374915",
		StorageBucket: "gs://dev-env-374915.appspot.com/",
	}

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(bucket)

	//url, err := bucket.SignedURL("pytfpCmDmcVdaEUKG86eHRzFaf92/f882afd0e550faa62e629d8977f6e47.jpg", nil)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(url)

}
