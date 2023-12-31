package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis_rate/v10"
)

func (r redisRepo) UserRemaining(username string, rate int) int {
	key := fmt.Sprintf("%s:request", username)
	res, err := r.limiter.Allow(r.ctx, key, redis_rate.PerMinute(rate))
	if err != nil {
		r.logger.Error(err.Error())
	}
	return res.Remaining
}

func (r redisRepo) UserMonthlyUsageUpdate(key string, incr int64) {
	_, err := r.conn.IncrBy(r.ctx, key, incr).Result()
	if err != nil {
		r.logger.Error(err.Error())
	}
}

func (r redisRepo) UserMonthlyUsage(key string) int64 {
	switch r.Exists(key) {
	case true:
		return r.Get(key)
	case false:
		r.Set(key, 0, 30*24*time.Hour)
		return 0
	}
	return 0
}
