package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/smplverse/metadata/data"
)

func TestServer(t *testing.T) {
	t.Run("serves metadata right", func(t *testing.T) {
		metadata := data.Metadata{
			"1": {
				TokenID:     "1",
				Name:        "name",
				Description: "description",
				Image:       "image",
				ExternalURL: "external_url",
				Attributes: []data.Attribute{
					{
						TraitType: "trait_type",
						Value:     "value",
					},
				},
			},
		}

		router := httprouter.New()
		router.GET("/:tokenID", Handle(metadata))

		req, err := http.NewRequest("GET", "/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != 200 {
			t.Fatalf("expected status code 200, got %d", rr.Code)
		}

		contentType := rr.Header().Get("Content-Type")
		if contentType != "application/json" {
			t.Fatalf("expected content type application/json, got %s", contentType)
		}

		got, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatal(err)
		}

		want, err := json.Marshal(metadata["1"])
		if err != nil {
			t.Fatal(err)
		}

		if string(got) != string(want) {
			t.Fatalf("expected %s, got %s", string(want), string(got))
		}
	})
}
