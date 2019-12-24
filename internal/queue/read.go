package queue

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Read(queue string) string {
	result, err := client.Get(queue).Result()

	if err == redis.Nil {
		fmt.Println(queue, "does not exist")
		return ""
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(queue, result)
		return result
	}
}
