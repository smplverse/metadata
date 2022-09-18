package data

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"cloud.google.com/go/storage"
)

type Metadata map[string]MetadataEntry

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

func Get(ctx context.Context) (Metadata, error) {
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

	buf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	log.Println(string(buf[:30]))
	var metadata Metadata
	err = json.Unmarshal(buf, &metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}
