package middleware

import (
	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/repository"
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/pkg/logger"
)

type middle struct {
	config config.Middleware
	redis  repository.Redis
	logger logger.Logger
}

func New(cfg config.Middleware, rd repository.Redis, logger logger.Logger) service.Middleware {
	return middle{config: cfg, redis: rd, logger: logger}
}

func (m middle) UserQuotaRequest(username string) bool {
	remain := m.redis.UserRemaining(username, m.config.PerMinute)
	return remain == 0
}

func (m middle) UserQuotaTraffic(username string, size int64) bool {
	return m.redis.UserMonthlyUsage(username) >= int64(m.config.Monthly)
}
