package queue

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Setup() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
