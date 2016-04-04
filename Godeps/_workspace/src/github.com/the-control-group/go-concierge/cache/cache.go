package cache

import (
	"time"
)

// Cacher is ...
type Cacher interface {
	Storer
	Key(prefix string, request []byte) string
	Get(key string) ([]byte, error)
}

// Clearer is an interface for clearing the cache
type Clearer interface {
	Clear(prefix string) error
}

// Storer is ...
type Storer interface {
	GetName() string
	Put(provider, key string, value []byte, options interface{}) error
}

// PutOptions is ...
type PutOptions struct {
	// StoreName overrides provider when storing
	StoreName string
	TTL       time.Duration
}

// Response is ...
type Response struct {
	Name     string
	OK       bool
	Error    error
	Timeout  bool
	Duration time.Duration
	Data     []byte
}
