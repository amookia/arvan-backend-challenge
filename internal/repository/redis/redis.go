package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

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

func (r redisRepo) Exists(key string) bool {
	exist, _ := r.conn.Exists(r.ctx, key).Result()
	return exist == 1
}

func (r redisRepo) Get(key string) int64 {
	res, _ := r.conn.Get(r.ctx, key).Result()
	val, _ := strconv.ParseInt(res, 10, 64)
	return val
}

func (r redisRepo) Set(key string, val int64, expire time.Duration) {
	r.conn.Set(r.ctx, key, val, expire)
}

func (r redisRepo) IncrBy(key string, val int64) {
	r.conn.IncrBy(r.ctx, key, val)
}
