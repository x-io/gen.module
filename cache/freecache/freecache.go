package freecache

import (
	"time"

	"github.com/coocood/freecache"
	"github.com/x-io/gen.module/cache"
)

// CacheFree CacheFree
type CacheFree struct {
	conn *freecache.Cache
}

// Init connects to the database.
func Init(size int) *CacheFree {
	return &CacheFree{
		conn: freecache.NewCache(size),
	}
}

// Set Cache Set
func (t *CacheFree) Set(key string, value []byte, expire time.Duration) error {
	return t.conn.Set([]byte(key), value, int(expire.Seconds()))
}

// Get Cache Get
func (t *CacheFree) Get(key string) ([]byte, error) {
	d, err := t.conn.Get([]byte(key))

	if err != nil {
		if err == freecache.ErrNotFound {
			return nil, cache.NotFound
		}

		return nil, err
	}
	return d, nil
}

// Del Cache Del
func (t *CacheFree) Del(key string) error {
	if t.conn.Del([]byte(key)) {
		return nil
	}
	return cache.NotFound
}
