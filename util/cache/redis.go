package cache

import (
	"context"
	"gin-template/config"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var client *redis.Client

type Repo interface {
	i()
	GetClient() *redis.Client
}

type redisRepo struct {
	client *redis.Client
}

func newRedis() (Repo, error) {
	connect, err := redisConnect()
	if err != nil {
		return nil, err
	}
	return &redisRepo{client: connect}, nil
}

func init() {
	repo, _ := newRedis()
	client = repo.GetClient()
}

// redisConnect 连接redis
func redisConnect() (*redis.Client, error) {
	cfg := config.Get().Redis
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Pass,
		DB:           cfg.Db,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConn,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}
	return client, nil
}

func (redis *redisRepo) i() {}

func (redis *redisRepo) GetClient() *redis.Client {
	return redis.client
}
