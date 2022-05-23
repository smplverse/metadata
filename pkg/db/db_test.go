package db_test

import (
	"testing"

	"github.com/piotrostr/metadata/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestDbConnects(t *testing.T) {
	d, err := db.Connect()
	assert.Nil(t, err)
	assert.NotNil(t, d)
}
