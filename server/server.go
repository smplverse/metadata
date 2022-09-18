package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smplverse/metadata/data"
)

func Handler(metadata data.Metadata) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buf, err := json.Marshal(metadata)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(buf))
	}
}

func Serve(metadata data.Metadata, port string) error {
	http.HandleFunc("/", Handler(metadata))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		return err
	}
	return nil
}
