package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rdb *redis.Client
	ctx context.Context
	ttl int
}

var RClient *RedisClient

func Connect(ctx context.Context, ttl int) {
	username := os.Getenv("redis_username")
	password := os.Getenv("redis_password")
	if len(username) == 0 || len(password) == 0 {
		panic("could not fetch username and password for redis client")
	}

	// connect and test connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-19869.c282.east-us-mz.azure.redns.redis-cloud.com:19869",
		Username: username,
		Password: password,
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	RClient = &RedisClient{
		rdb: rdb,
		ctx: ctx,
		ttl: ttl,
	}
}

func (c *RedisClient) CheckCache(key string) (string, error) {
	return c.rdb.Get(c.ctx, key).Result()
}

func (c *RedisClient) StoreCache(key string, val string) error {
	err := c.rdb.Set(c.ctx, key, val, time.Duration(c.ttl*int(time.Minute))).Err()
	if err != nil {
		return fmt.Errorf("unable to store key-val to cache: %v", err)
	}
	return nil
}
