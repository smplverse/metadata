package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadata_GetUnclaimed(t *testing.T) {
	m := New()
	entry := m.Get("4")
	assert.Equal(t, entry, BlankEntry)
}

func TestMetadata_GetClaimed(t *testing.T) {
	m := New()
	want := Entry{}
	m.entries["5"] = want
	got := m.Get("5")
	assert.Equal(t, got, want)
}
