package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis - Example")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// We can set 'key' and 'value'
	err = client.Set("name", "Jonathan", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// Getting values
	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

}
