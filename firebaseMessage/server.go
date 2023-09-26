package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

const token = "dSbEBCJsOUp2sUAjE8iKGl:APA91bHPkyaojyjKcVJBzep3yMv1PpHZP6x79Cd0hzCosJjArPVDUQTZkVtpiEYxleU0wCu0zDYexU1IMBX2_hoyrBbDBMfZAXMN8lOxamZos8mU6z-fGCITD_fQd3dBDLybMb8k58p2"

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

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "123123123",
			Body:  "Barcelona vs. Juventus",
		},

		//Token: token,
		Topic: "testTopic",
	}

	// Send a message to the devices subscribed to the provided topic.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

}
