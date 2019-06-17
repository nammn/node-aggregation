package database

import (
	"fmt"
	"github.com/go-redis/redis"
)

type RedisConnection interface {
	SaveNodeStatValue(key NodeStat) error
}

type RedisClient struct {
	client *redis.Client
}

func NewClient(password string) (*RedisClient, error) {
	client, err := CreateClient(password)
	if err != nil {
		return nil, err
	}
	return &RedisClient{client: client}, nil
}

func CreateClient(password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(pong, err)
	return client, err
}

// responsible for trying to fit the values to a corresponding match in the database to save
func (r RedisClient) SaveNodeStatValue(key NodeStat) error {
	_, _ = fmt.Printf("node obj %+v\n", key)
	return nil
}
