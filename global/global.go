package global

import (
	"cloud.google.com/go/firestore"
	"context"
	fAuth "firebase.google.com/go/auth"
	"log"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"time"
)


var Firestore *firestore.Client
var Auth *fAuth.Client

func InitVariables() {
	ctx, _ := getContext()
	opt := option.WithCredentialsFile("env/serviceAccountKey.json")
	client, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	Firestore, err = client.Firestore(ctx)
	if err != nil {
		log.Fatal("Could not initialize firestore")
	}
	Auth, err = client.Auth(ctx)
	if err != nil {
		log.Fatal("Could not initialize firebase Auth")
	}
}

func getContext() (context.Context, context.CancelFunc) {
	timout := 5 * time.Minute
	return context.WithTimeout(context.Background(), timout)
}
