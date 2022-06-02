package metadata

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/piotrostr/metadata/pkg/config"
	"github.com/piotrostr/metadata/pkg/db"
)

var ctx = context.Background()

type Metadata struct {
	rdb  *redis.Client
	base config.Base
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
	rdb := db.Client()

	base, err := config.Get()
	if err != nil {
		log.Println("Error getting config: ", err)
		log.Println("Using default")
	}

	// if there is no config will return zeros for the given fields
	return &Metadata{
		rdb:  rdb,
		base: base,
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
		entry = &Entry{
			TokenId:     "#",
			Name:        m.base.Name,
			Description: m.base.Description,
			ExternalUrl: m.base.ExternalUrl,
			Image:       m.base.Image,
			Attributes:  []Attribute(nil),
		}
		err = nil
		return
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
