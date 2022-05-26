package metadata

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/piotrostr/metadata/pkg/db"
)

const DESCRIPTION = "SMPLverse is a collection of synthetic face data from the computational infrastructure of the metaverse, assigned to minters using facial recognition."

const PLACEHOLDER_IMAGE = "ipfs://QmYypT49WH7rYTL2jXpfoNH2DAMHe9VM7pwwEjUVr45XK1"

var ctx = context.Background()

var BlankEntry = Entry{
	TokenId:     "#",
	Name:        "UNCLAIMED SMPL",
	Description: DESCRIPTION,
	ExternalUrl: "",
	Image:       PLACEHOLDER_IMAGE,
	Attributes:  []Attribute(nil),
}

type Metadata struct {
	rdb *redis.Client
}

// TODO use default values for description, image and extUrl
type Entry struct {
	TokenId     string      `json:"token_id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	ExternalUrl string      `json:"external_url,omitempty"`
	Image       string      `json:"image,omitempty"`
	Attributes  []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	TraitType string `json:"trait_type,omitempty"`
	Value     string `json:"value,omitempty"`
}

var _ = []string{
	"037544",
	"069701",
	"099370",
	"093321",
	"051039",
	"046594",
	"059759",
	"074727",
	"083824",
	"037661",
	"059324",
}

func New() *Metadata {
	return &Metadata{
		rdb: db.Client(),
	}
}

func ValidateMetadataEntry(metadataEntry Entry) bool {
	invalidTokenId := metadataEntry.TokenId == ""
	invalidImage := metadataEntry.Image == ""
	invalidAttrs := len(metadataEntry.Attributes) == 0

	if invalidTokenId || invalidImage || invalidAttrs {
		return false
	}

	return true
}

func (m *Metadata) Get(tokenId string) (entry *Entry, err error) {
	entryString, err := m.rdb.Get(ctx, tokenId).Result()
	if err == redis.Nil {
		entry = &BlankEntry
		return entry, nil
	}

	var entryObj Entry
	err = json.Unmarshal([]byte(entryString), &entryObj)
	if err != nil {
		return
	}

	entry = &entryObj
	return
}

func (m *Metadata) Add(tokenId string, entry Entry) (err error) {
	valid := ValidateMetadataEntry(entry)
	if !valid {
		err = errors.New("invalid metadata entry")
		return
	}
	entryString, _ := json.Marshal(entry)
	err = m.rdb.Set(ctx, tokenId, string(entryString), 0 /* expiration time */).Err()
	return
}
