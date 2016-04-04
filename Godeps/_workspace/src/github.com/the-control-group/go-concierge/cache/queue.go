package cache

// RedisQueue is a storer
type RedisQueue struct {
	r RedisCache
}

// NewRedisQueue creates a new RedisQueue with the provided DSN
func NewRedisQueue(dsn string) RedisQueue {
	return RedisQueue{r: NewRedisCache(dsn)}
}

// GetName returns "redis_queue"
func (q RedisQueue) GetName() string {
	return "redis_queue"
}

// Put sends value to the queue
func (q RedisQueue) Put(queue, key string, value []byte, options interface{}) error {
	var err error

	conn := q.r.pool.Get()
	defer conn.Close()

	_, err = conn.Do("RPUSH", queue, value)

	return err
}
