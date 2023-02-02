package utils

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),  // no password set
		Password: viper.GetString("redis.password"), // no password set
		DB:       0,                                 // use default DB
	})
}
