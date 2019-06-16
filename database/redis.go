package database

import (
	"fmt"
	"github.com/go-redis/redis"
)

func CreateClient(password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client, err
}

func SaveValue(key string, value string) error {
	return nil
}
