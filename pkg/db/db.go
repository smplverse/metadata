package db

import (
	"fmt"
	"os"

	redis "github.com/go-redis/redis/v8"
)

func Client() *redis.Client {
	// in the cluster only the pods can connect, hence no need for password
	host := os.Getenv("REDIS_HOST")
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", host),
	})
	return rdb
}
