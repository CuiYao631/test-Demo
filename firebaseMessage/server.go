package main

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

const token = "dSbEBCJsOUp2sUAjE8iKGl:APA91bHPkyaojyjKcVJBzep3yMv1PpHZP6x79Cd0hzCosJjArPVDUQTZkVtpiEYxleU0wCu0zDYexU1IMBX2_hoyrBbDBMfZAXMN8lOxamZos8mU6z-fGCITD_fQd3dBDLybMb8k58p2"

func main() {
	//firebaseApp
	opt := option.WithCredentialsFile("/Users/xiaocui/zhihe-tech-e-bike-tracking-firebase-adminsdk-ywtit-fb165d633c.json")
	config := &firebase.Config{

		ProjectID:     "zhihe-tech-e-bike-tracking",
		StorageBucket: "gs://dev-env-374915.appspot.com/",
	}

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	ctx := context.Background()

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	user, err := authClient.GetUser(ctx, "Nq9ppxBAG9VoIDvknneBeopR4gx1")
	if err != nil {
		log.Println(err)
	}
	log.Println(user.Email)

}
