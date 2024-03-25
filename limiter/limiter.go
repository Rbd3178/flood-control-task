package limiter

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// Limiter
type Limiter struct {
	client   *redis.Client
	interval int
	limit    int64
}

// New establishes a connection with a Redis database,
// sets interval and amount of calls allowed in that interval
// and returns a pointer to a Limiter instance
func New(cfg *Config) *Limiter {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})
	return &Limiter{
		client:   client,
		interval: cfg.Interval,
		limit:    cfg.Limit,
	}
}

// Check
func (l *Limiter) Check(ctx context.Context, userID int64) (bool, error) {
	userIDKey := strconv.Itoa(int(userID))
	count, err := l.client.Get(ctx, userIDKey).Int64()
	if err != nil && err != redis.Nil {
		return false, err
	}

	if count >= l.limit {
		return false, nil
	}

	p := l.client.TxPipeline()
	p.Incr(ctx, userIDKey)
	p.Expire(ctx, userIDKey, time.Duration(l.interval)*time.Second)
	_, err = p.Exec(ctx)
	if err != nil {
		return false, err
	}

	go func() {
		time.Sleep(time.Duration(l.interval) * time.Second)
		exists, err := l.client.Exists(ctx, userIDKey).Result()
		if exists == 1 && err == nil {
			l.client.Decr(ctx, userIDKey)
		}
	}()
	return true, nil
}
