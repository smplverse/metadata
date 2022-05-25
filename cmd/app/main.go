package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/piotrostr/metadata/pkg/server"
)

func main() {
	router, err := server.SetupRouter()
	if err != nil {
		log.Fatalln(err)
	}

	port := *flag.String("port", "80", "port to listen on")
	flag.Parse()

	log.Printf("running on :%s\n", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalln(err)
	}
}
