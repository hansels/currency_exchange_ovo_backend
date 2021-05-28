package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/hansels/currency_exchange_ovo_backend/config"
	"google.golang.org/api/option"
	"log"
)

var App *firebase.App

type UploadImageData struct {
	Ctx      context.Context
	FileName string
	File     []byte
}

func init() {
	opt := option.WithCredentialsFile("./files/firebase/firebase.json")
	cfg := &firebase.Config{
		ProjectID:     config.FirebaseProjectId,
		StorageBucket: config.FirebaseStorageBucket,
	}

	var err error
	App, err = firebase.NewApp(context.Background(), cfg, opt)
	if err != nil {
		log.Fatalf("error initializing firebase: %v\n", err)
	}

	log.Println("Firebase Initialization Success 🔥")
}

func InitFirestore() *firestore.Client {
	fs, err := App.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing firestore: %v\n", err)
	}
	log.Println("Firestore Roll-on 🔥🔥🔥")
	return fs
}
