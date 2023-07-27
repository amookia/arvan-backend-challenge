package redis

import (
	"fmt"
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/service/repository"
	"github.com/redis/go-redis/v9"
)

type redisRepo struct {
	conn   *redis.Client
	logger *log.Logger
}

func New(config config.Redis, logs *log.Logger) (repository.Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Host, config.Port),
	})
	return redisRepo{conn: rdb, logger: logs}, nil
}
