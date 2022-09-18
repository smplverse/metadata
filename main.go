package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	port := 8080
	log.Printf("listening on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
