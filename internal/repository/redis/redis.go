package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/repository"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type redisRepo struct {
	conn    *redis.Client
	limiter *redis_rate.Limiter
	logger  *log.Logger
	ctx     context.Context
}

func New(config config.Redis, ctx context.Context, logs *log.Logger) (repository.Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Host, config.Port),
	})
	return redisRepo{
		conn: rdb, limiter: redis_rate.NewLimiter(rdb), logger: logs, ctx: ctx,
	}, nil
}
