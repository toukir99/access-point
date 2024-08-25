package db

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "my-redis",               // Set your Redis password here
		DB:       0,                // Set the Redis database number
	})

	// Test connection
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}

func GetRedisClient() *redis.Client {
	return Rdb
}