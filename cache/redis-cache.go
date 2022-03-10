package cache

import (
	"encoding/json"
	"studentList/models"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) StudentCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}
func (cache *redisCache) Set(key string, value *models.Student) {
	student := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	student.Set(key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *models.Student {
	student := cache.getClient()
	val, err := student.Get(key).Result()
	if err != nil {
		return nil
	}
	post := models.Student{}
	json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}
	return &post
}
