package redis

import (
	"fmt"
	config "github.com/spf13/viper"
	redis "gopkg.in/redis.v5"
	"sync"
	"time"
)

var redisConfig = map[string]*redis.Client{
	"master": nil,
	"slave":  nil,
}

var lock = sync.RWMutex{}
var RedisNil = redis.Nil

func Master() *redis.Client {
	if redisConfig["master"] == nil {
		redisConfig["master"] = createRedisConnection("master")
	}

	return PreserveRedisConn("master")
}

func Slave() *redis.Client {
	if redisConfig["slave"] == nil {
		redisConfig["slave"] = createRedisConnection("slave")
	}

	return PreserveRedisConn("slave")
}

func createRedisConnection(types string) *redis.Client {
	lock.Lock()
	defer lock.Unlock()

	conn := redis.NewClient(&redis.Options{
		Addr:        config.GetString("redis." + types),
		MaxRetries:  2,
		IdleTimeout: 5 * time.Minute,
	})

	_, err := conn.Ping().Result()
	if err != nil {
		panic(err)
	}

	return conn
}

func PreserveRedisConn(connType string) *redis.Client {
	_, err := redisConfig[connType].Ping().Result()
	if err != nil {
		redisConfig[connType] = createRedisConnection(connType)
		fmt.Println("connection preserved..")
	}
	return redisConfig[connType]
}
