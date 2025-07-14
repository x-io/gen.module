package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/x-io/gen.module/cache"
)

// CacheRides CacheRides
type CacheRides struct {
	conn *redis.Client
}

// Init connects to the database.
func CacheInit(c *Config) *CacheRides {
	return &CacheRides{
		conn: redis.NewClient(&redis.Options{
			Addr:     c.Host,
			Password: c.Password, // no password set
			DB:       c.DB,       // use default DB
		}),
	}
}

// Set Cache Set
func (t *CacheRides) Set(key string, value []byte, expiration time.Duration) error {
	state := t.conn.Set(context.Background(), key, value, expiration)
	return state.Err()
}

// Get Cache Get
func (t *CacheRides) Get(key string) ([]byte, error) {
	res := t.conn.Get(context.Background(), key)
	d, err := res.Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, cache.NotFound
		}

		return nil, err
	}

	return d, nil
}

// Del Cache Del
func (t *CacheRides) Del(key string) error {
	res := t.conn.Del(context.Background(), key)

	return res.Err()
}
