package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUnclaimed(t *testing.T) {
	m := New()
	m.rdb.Del(ctx, "4")
	entry, err := m.Get("4")
	assert.Nil(t, err)
	assert.Equal(t, entry, &BlankEntry)
}

func TestGetInvalid(t *testing.T) {
	m := New()
	m.rdb.Set(ctx, "6", "invalid]", 0)
	entry, err := m.Get("6")
	assert.NotNil(t, err)
	assert.Nil(t, entry)
}

func TestGetClaimed(t *testing.T) {
	m := New()
	want := &Entry{
		TokenId:     "6",
		Name:        "Name",
		Description: "Description",
		ExternalUrl: "ExternalUrl",
		Image:       "Image",
	}
	entryString, err := json.Marshal(want)
	if err != nil {
		t.Error(err)
	}
	m.rdb.Set(ctx, "5", entryString, 0)
	got, err := m.Get("5")
	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestAddValid(t *testing.T) {
	m := New()
	want := Entry{
		TokenId:     "6",
		Name:        "Name",
		Description: "Description",
		ExternalUrl: "ExternalUrl",
		Image:       "Image",
		Attributes: []Attribute{
			{TraitType: "TraitType", Value: "Value"},
		},
	}
	err := m.Add("6", want)
	assert.Nil(t, err)
	got, err := m.Get("6")
	assert.Nil(t, err)
	assert.Equal(t, got, &want)
}

func TestAddInvalid(t *testing.T) {
	m := New()
	want := Entry{
		TokenId:     "6",
		Name:        "Name",
		Description: "Description",
		ExternalUrl: "ExternalUrl",
		Image:       "Image",
	}
	err := m.Add("6", want)
	assert.NotNil(t, err)
	got, err := m.Get("6")
	assert.Nil(t, err)
	assert.NotEqual(t, got, &want)
}

func TestValidateEntry(t *testing.T) {
	valid := Entry{
		TokenId:     "6",
		Name:        "Name",
		Description: "Description",
		ExternalUrl: "ExternalUrl",
		Image:       "Image",
		Attributes: []Attribute{
			{TraitType: "TraitType", Value: "Value"},
		},
	}
	invalid := Entry{}
	assert.True(t, ValidateMetadataEntry(valid))
	assert.False(t, ValidateMetadataEntry(invalid))
}
