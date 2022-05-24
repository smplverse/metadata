package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/piotrostr/metadata/pkg/metadata"
	"github.com/piotrostr/metadata/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestRequiresApiKeyEnvVarSet(t *testing.T) {
	os.Unsetenv("METADATA_API_KEY")
	defer os.Setenv("METADATA_API_KEY", "secret")

	router, err := server.SetupRouter()
	assert.Equal(t, err, server.ErrUnsetApiKey)
	assert.Nil(t, router)
}

func TestPingRoute(t *testing.T) {
	router, err := server.SetupRouter()
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestNonExistingMetadata(t *testing.T) {
	router, err := server.SetupRouter()
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/52", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var entry metadata.Entry
	err = json.Unmarshal(w.Body.Bytes(), &entry)
	assert.Nil(t, err)
	assert.Equal(t, &entry, &metadata.BlankEntry)
}

func Test401ForUnauthorized(t *testing.T) {
	router, err := server.SetupRouter()
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	emptyEntry := &metadata.Entry{}
	jsonEntry, err := json.Marshal(emptyEntry)
	if err != nil {
		t.Fatal(err)
	}

	entryToAdd := bytes.NewBuffer([]byte(jsonEntry))
	req, _ := http.NewRequest("POST", "/32", entryToAdd)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, `{"error":"unauthorized"}`, w.Body.String())
}

func TestAddingToMetadataWorks(t *testing.T) {
	router, err := server.SetupRouter()
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	emptyEntry := &metadata.Entry{}
	jsonEntry, err := json.Marshal(emptyEntry)
	if err != nil {
		t.Fatal(err)
	}

	entryToAdd := bytes.NewBuffer([]byte(jsonEntry))
	req, _ := http.NewRequest("POST", "/32", entryToAdd)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", os.Getenv("METADATA_API_KEY"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/32", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var entry metadata.Entry
	err = json.Unmarshal(w.Body.Bytes(), &entry)
	assert.Nil(t, err)
	assert.Equal(t, &entry, emptyEntry)
}
