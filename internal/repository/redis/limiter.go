package redis

import (
	"github.com/go-redis/redis_rate/v10"
)

func (r redisRepo) UserRemaining(username string, rate int) int {
	res, err := r.limiter.Allow(r.ctx, username, redis_rate.PerMinute(rate))
	if err != nil {
		r.logger.Fatal(err)
	}
	return res.Remaining
}
