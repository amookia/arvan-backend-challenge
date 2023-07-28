package middleware

import (
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/repository"
	"github.com/amookia/arvan-backend-challenge/internal/service"
)

type middle struct {
	config config.Middleware
	redis  repository.Redis
	logger *log.Logger
}

func New(cfg config.Middleware, rd repository.Redis, logger *log.Logger) service.Middleware {
	return middle{config: cfg, redis: rd, logger: logger}
}

func (m middle) IsUserLimited(username string) bool {
	remain := m.redis.UserRemaining(username, m.config.PerMinute)
	return remain == 0
}

func (m middle) IsUserLimitedCapacity(username string, size int64) bool {
	if m.redis.UserMonthlyUsage(username) >= int64(m.config.Monthly) {
		return true
	}
	m.redis.UserMonthlyUsageUpdate(username, size)
	return false
}
