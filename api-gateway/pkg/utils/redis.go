package utils

import "github.com/redis/go-redis/v9"

var RDB *redis.Client

func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "@999now!CN", // no password set
		DB:       0,            // use default DB
	})
}
