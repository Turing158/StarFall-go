package util

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

const redisKey = "StarFall:"

type RedisUtil struct {
}

var RedisClient *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("连接Redis出错：", err)
		return
	}
	fmt.Println("连接成功，Ping结果：", pong)
	RedisClient = client
}

func CloseRedis() {
	err := RedisClient.Close()
	if err != nil {
		return
	}
}

func (RedisUtil) Set(key string, value any) bool {
	err := RedisClient.Set(redisKey+key, value, 24*time.Hour).Err()
	if err != nil {
		return false
	}
	return true
}

func (r RedisUtil) SetObj(key string, value any) bool {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return false
	}
	return r.Set(key, string(jsonData))
}

func (RedisUtil) SetWithExpireTime(key string, value any, exp time.Duration) bool {
	err := RedisClient.Set(redisKey+key, value, exp).Err()
	if err != nil {
		return false
	}
	return true
}

func (r RedisUtil) SetObjWithExpireTime(key string, value any, exp time.Duration) bool {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return false
	}
	return r.SetWithExpireTime(key, string(jsonData), exp)
}

func (RedisUtil) Get(key string) any {
	value, err := RedisClient.Get(redisKey + key).Result()
	if err != nil {
		return nil
	}
	return value
}

func (RedisUtil) GetObj(key string, obj interface{}) any {
	value, err := RedisClient.Get(redisKey + key).Result()
	if err != nil {
		return nil
	}
	objP := &obj
	errJson := json.Unmarshal([]byte(value), objP)
	if errJson != nil {
		return nil
	}
	return *objP
}

func (RedisUtil) Del(key string) error {
	return RedisClient.Del(redisKey + key).Err()
}

func (RedisUtil) Has(key string) bool {
	return RedisClient.Exists(redisKey+key).Val() > 0
}
