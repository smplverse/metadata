package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
)

var ctx = context.Background()

type Metadata []MetadataEntry

type MetadataEntry struct {
	TokenID     string      `json:"token_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	ExternalURL string      `json:"external_url"`
	IPFSURL     string      `json:"ipfs_url"`
	Attributes  []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType string
	Value     string
}

func getMetadata() (Metadata, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := client.Bucket("smplverse").Object("metadata.json")

	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	buf := make([]byte, reader.Size())
	if _, err := reader.Read(buf); err != nil {
		return nil, err
	}

	var metadata Metadata
	err = json.Unmarshal(buf, &metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	metadata, err := getMetadata()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	buf, err := json.Marshal(metadata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	port := 8080
	log.Printf("listening on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
