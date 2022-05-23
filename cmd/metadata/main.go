package main

import (
	"github.com/piotrostr/metadata/pkg/db"
	_ "github.com/piotrostr/metadata/pkg/server"
)

func main() {
	db.Connect()
	// server.Run()
}
