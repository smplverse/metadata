package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/piotrostr/metadata/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "ok\n", w.Body.String())
}
