package database

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"strconv"
)

const (
	DEFAULTIMESLICE = 60
	NODEKEY         = "nodestat"
)

type RedisConnection interface {
	SaveNodeStatValue(key NodeStat) error
	Ping() error
	GetFromGivenAverage(timeslice float64) error
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
func (r RedisClient) SaveNodeStatValue(nodeStat NodeStat) error {
	if nodeStat == (NodeStat{}) {
		return errors.New("empty parsed ns")
	}
	//TODO: more sophisticated error handling. Not Valid Nodestat etc.
	ns, err := json.Marshal(nodeStat)
	if err != nil {
		return nil
	}
	r.client.ZAdd(NODEKEY, redis.Z{Score: float64(nodeStat.Timestamp), Member: ns})
	_, _ = fmt.Printf("node obj %+v\n", nodeStat)
	return nil
}

// returns the average from zadd given the timeslice
// Uses ZRANGEBYSCORE
// https://redis.io/commands/zrangebyscore
// Uses additionally a peak to get the "newest" value, which should be in O(1)
// because of network one can assume to cache this here instead
func (r RedisClient) GetFromGivenAverage(timeslice float64) error {
	if timeslice == 0 {
		timeslice = DEFAULTIMESLICE
	}
	newestDate, err := strconv.ParseFloat(r.client.ZRevRange(NODEKEY, 0, 0).String(), 64)
	if err != nil {
		return nil
	}
	from := fmt.Sprintf("%f", newestDate-timeslice)
	to := fmt.Sprintf("%f", newestDate)
	vals, err := r.client.ZRangeByScore(NODEKEY, redis.ZRangeBy{Min: from, Max: to}).Result()
	if err != nil {
		return nil
	}
	fmt.Printf("%+v", vals)
	return nil
}

func (r RedisClient) Ping() error {
	err := r.client.Ping().Err()
	if err != nil {
		return err
	}
	return nil
}
