package storage

import (
	"fmt"

	"github.com/go-redis/redis"
)

var queueClient *redis.Client

func SetupQueue() {
	queueClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := queueClient.Ping().Result()
	fmt.Println(pong, err)
}

func QueueWrite(queue string, value string) {
	err := queueClient.Set(queue, value, 0).Err()

	if err != nil {
		panic(err)
	}
}

func QueueRead(queue string) {
	result, err := queueClient.Get(queue).Result()

	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(queue, result)
	}
}
