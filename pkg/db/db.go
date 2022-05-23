package db

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

func Connect() (db *badger.DB, err error) {
	db, err = badger.Open(badger.DefaultOptions("/tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(txn *badger.Txn) error {
		// Your code here…
		return nil
	})

	defer db.Close()

	return
}
