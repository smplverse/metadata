package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

const PROJECT_ID = "smplverse-metadata"

func Connect() {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatal()
	}

	m := client.Collection("Metadata")

	log.Print(m)
}
