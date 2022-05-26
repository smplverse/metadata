package db

import (
	"context"
	"testing"

	redis "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func TestClient(t *testing.T) {
	rdb := Client()
	defer rdb.Close()

	err := rdb.Set(ctx, "key", "value", redis.KeepTTL).Err()
	if err != nil {
		t.Fatal(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "value", val)
}
