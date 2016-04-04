package cache

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// SSDBCache is a ssdb-compatible Cacher
type SSDBCache struct {
	pool     redis.Pool
	Host     string
	Database int
}

// NewSSDBCache creates a new SSDBCache with the provided settings
func NewSSDBCache(dsn string) SSDBCache {
	r := SSDBCache{}

	// validate settings, apply defaults
	settings, _ := Parse(dsn)

	r.Host = settings.Host()
	if r.Host == "" {
		r.Host = "127.0.0.1:6379"
	}

	r.Database = settings.GetInt("database", 0)

	poolMaxActive := settings.GetInt("pool_max_active", defaultPoolMaxActive)
	poolMaxIdle := settings.GetInt("pool_max_idle", defaultPoolMaxIdle)

	idleTimeoutStr := settings.GetString("pool_idle_timeout", "")
	idleTimeout, err := time.ParseDuration(idleTimeoutStr)
	if err != nil {
		idleTimeout = defaultPoolIdleTimeout
	}

	r.pool = redis.Pool{
		MaxIdle:     poolMaxIdle,
		MaxActive:   poolMaxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", r.Host)
			if err != nil {
				return nil, err
			}
			if db := r.Database; db > 0 {
				if _, err := conn.Do("SELECT", db); err != nil {
					return nil, err
				}
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return r
}

// GetName returns "ssdb"
func (r SSDBCache) GetName() string {
	return "ssdb"
}

// Key generates a cache key using the provided prefix and request
func (r SSDBCache) Key(prefix string, request []byte) string {
	return Key(prefix, request)
}

// Get fetches the cache data using the provided key
func (r SSDBCache) Get(key string) ([]byte, error) {
	conn := r.pool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))

	// we don't care about nil returned error
	if err != nil && err.Error() == "redigo: nil returned" {
		return data, nil
	}

	return data, err
}

// Put stores the cache data using the provided key and options
func (r SSDBCache) Put(call, key string, value []byte, options interface{}) error {
	var err error

	conn := r.pool.Get()
	defer conn.Close()

	ttl := 0
	if options, ok := options.(PutOptions); ok {
		ttl = int(options.TTL / time.Second)
	}

	if ttl > 0 {
		_, err = conn.Do("SETX", key, value, ttl)
	} else {
		_, err = conn.Do("SET", key, value)
	}

	return err

}
