package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App
var firebaseAuthClient *auth.Client

// Init initialize firebase app
func init() {
	var err error
	opt := option.WithCredentialsFile("firebase.json")

	firebaseApp, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	firebaseAuthClient, err = firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
}

// VerifyToken verify firebase token
func VerifyToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := firebaseAuthClient.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return nil, err
	}

	log.Printf("Verified ID token: %v\n", token)
	return token, nil
}
