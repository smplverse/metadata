package server

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/smplverse/metadata/data"
)

func TestServer(t *testing.T) {
	t.Run("serves metadata right", func(t *testing.T) {
		metadata := data.Metadata{
			{
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

		srv := httptest.NewServer(Handler(metadata))
		defer srv.Close()

		res, err := srv.Client().Get(srv.URL + "/")
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != 200 {
			t.Fatalf("expected status code 200, got %d", res.StatusCode)
		}

		if res.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("expected content type application/json, got %s", res.Header.Get("Content-Type"))
		}

		got, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		want, err := json.Marshal(metadata)
		if err != nil {
			t.Fatal(err)
		}

		if string(got) != string(want) {
			t.Fatalf("expected %s, got %s", string(want), string(got))
		}
	})
}
