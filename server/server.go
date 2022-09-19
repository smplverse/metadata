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
	router := httprouter.New()
	router.GET("/:tokenID", Handle(metadata))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		return err
	}
	return nil
}
