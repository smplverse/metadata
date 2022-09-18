package main

import (
	"context"
	"log"

	"github.com/smplverse/metadata/data"
	"github.com/smplverse/metadata/server"
)

var ctx = context.Background()

func main() {
	log.Print("fetching metadata from google storage")
	metadata, err := data.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	port := "8080"
	log.Print("starting server on port 8080")
	err = server.Serve(metadata, port)
	if err != nil {
		log.Fatal(err)
	}
}
