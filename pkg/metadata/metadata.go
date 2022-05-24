package metadata

const DESCRIPTION = "SMPLverse is a collection of synthetic face data from the computational infrastructure of the metaverse, assigned to minters using facial recognition."

const PLACEHOLDER_IMAGE = "ipfs://QmYypT49WH7rYTL2jXpfoNH2DAMHe9VM7pwwEjUVr45XK1"

type Metadata struct {
	entries map[string]Entry
}

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

func (m *Metadata) Get(tokenId string) *Entry {
	if entry, ok := m.entries[tokenId]; ok {
		return &entry
	}
	return &Entry{
		TokenId:     tokenId,
		Name:        "UNCLAIMED SMPL",
		Description: DESCRIPTION,
		ExternalUrl: "",
		Image:       PLACEHOLDER_IMAGE,
		Attributes:  []Attribute{},
	}
}

func (m *Metadata) Add(tokenId string, entry Entry) {}
