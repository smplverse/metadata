package db

type Metadata struct {
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

var clusteredOnes = []string{
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
