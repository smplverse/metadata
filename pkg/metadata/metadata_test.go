package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUnclaimed(t *testing.T) {
	m := New()
	entry := m.Get("4")
	assert.Equal(t, entry, &BlankEntry)
}

func TestGetClaimed(t *testing.T) {
	m := New()
	want := Entry{}
	m.entries["5"] = want
	got := m.Get("5")
	assert.Equal(t, got, &want)
}

func TestAdd(t *testing.T) {
	m := New()
	entry := Entry{}
	m.Add("6", entry)
	got := m.Get("6")
	assert.Equal(t, got, &entry)
}
