package redis

import (
	"fmt"
	config "github.com/spf13/viper"
	redis "gopkg.in/redis.v5"
	"sync"
	"time"
)

type Redis struct {
	client *redis.Client
}

var redisConfig = map[string]*Redis {
	"master": nil,
	"slave":  nil,
}

type Storage interface {
	Set(key string, data string, expr time.Duration) error
	Get(key string) string
}

var lock = sync.RWMutex{}
var RedisNil = redis.Nil

func GetConn(str string) *Redis {
	switch str {
	case "master":
		return master()
	case "slave":
		return slave()
	default:
		return master()
	}
}

func master() *Redis {
	if redisConfig["master"] == nil {
		redisConfig["master"] = createRedisConnection("master")
	}

	return PreserveRedisConn("master")
}

func slave() *Redis {
	if redisConfig["slave"] == nil {
		redisConfig["slave"] = createRedisConnection("slave")
	}

	return PreserveRedisConn("slave")
}

func createRedisConnection(types string) *Redis {
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

	return &Redis{conn}
}

func PreserveRedisConn(connType string) *Redis {
	_, err := redisConfig[connType].client.Ping().Result()
	if err != nil {
		redisConfig[connType] = createRedisConnection(connType)
		fmt.Println("connection preserved..")
	}
	return redisConfig[connType]
}

func (c *Redis) Set(key string, data string, expr time.Duration) error {
	if err := c.client.Set(key, data, expr).Err(); err != nil {
		return err
	}

	return nil
}

func (c *Redis) Get(key string) string {
	b, err := c.client.Get(key).Result()
	if err != nil {
		return ""
	}

	return b
}
