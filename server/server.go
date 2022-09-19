package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/smplverse/metadata/data"
)

func Handle(metadata data.Metadata) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenId := ps.ByName("tokenID")
		metadataEntry, ok := metadata[tokenId]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		buf, err := json.Marshal(metadataEntry)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) // nolint: errcheck
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(buf))
	}
}

func Serve(metadata data.Metadata, port string) error {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "smplverse/metadata")
	})
	router.GET("/healthz", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "im healthy")
	})
	router.GET("/v1/:tokenID", Handle(metadata))

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		return err
	}

	return nil
}
