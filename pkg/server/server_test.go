package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/piotrostr/metadata/pkg/metadata"
	"github.com/piotrostr/metadata/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestNonExistingMetadata(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/52", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var entry metadata.Entry
	err := json.Unmarshal(w.Body.Bytes(), &entry)
	assert.Nil(t, err)
	assert.Equal(t, &entry, &metadata.BlankEntry)
}
