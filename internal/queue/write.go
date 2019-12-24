package queue

func Enqueue(queue string, value string) string {
	err := client.Set(queue, value, 0).Err()

	if err != nil {
		panic(err)
	}

	return "OK"
}
