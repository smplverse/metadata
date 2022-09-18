package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/smplverse/metadata/data"
)

func serve(metadata data.Metadata, port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf, err := json.Marshal(metadata)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		return err
	}
	return nil
}

var ctx = context.Background()

func main() {
	log.Print("fetching metadata from google storage")
	metadata, err := data.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	port := "8080"
	log.Print("starting server on port 8080")
	err = serve(metadata, port)
	if err != nil {
		log.Fatal(err)
	}
}
