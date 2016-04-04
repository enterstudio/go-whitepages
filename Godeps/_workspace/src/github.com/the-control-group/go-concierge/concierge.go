package concierge

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/the-control-group/go-concierge/cache"
	"github.com/the-control-group/go-concierge/stats"
)

const (
	cachePrefix         = "cache_"
	defaultTimeout      = "5s"
	defaultCacheTimeout = "200ms"
	defaultStoreTimeout = "200ms"
	defaultCacheTTL     = "12h"
)

// Doer defines how to do a request. The RequestBytes and ResponseBytes are needed for caching purposes.
type Doer interface {
	GetRequest() interface{}
	GetRequestBytes() []byte

	GetResponse() interface{}
	GetResponseBytes() []byte

	ResponseError() error
	AddStats(map[string]interface{})

	Do(time.Duration) error
}

// Options are for caching and stats. Setting these fields to nil effectively disables these features.
type Options struct {
	Cache        cache.Cacher
	Store        cache.Storer
	PutOptions   cache.PutOptions
	Stats        stats.StatHandler
	Timeout      time.Duration
	CacheTimeout time.Duration
	StoreTimeout time.Duration
	Username     string
	Password     string
}

// Request checks the cache if enabled, runs the Doer, puts the response in the cache if enabled, and reports stats if enabled.
func Request(name string, doer Doer, opts Options) ([]byte, bool, error) {
	var cacheKey string
	var err error

	// Set defaults, if needed
	timeout, _ := time.ParseDuration(defaultTimeout)
	if opts.Timeout != 0 {
		timeout = opts.Timeout
	}
	cacheTimeout, _ := time.ParseDuration(defaultCacheTimeout)
	if opts.CacheTimeout != 0 {
		cacheTimeout = opts.CacheTimeout
	}
	storeTimeout, _ := time.ParseDuration(defaultStoreTimeout)
	if opts.StoreTimeout != 0 {
		storeTimeout = opts.StoreTimeout
	}
	cacheTTL, _ := time.ParseDuration(defaultCacheTTL)
	putOptions := cache.PutOptions{TTL: cacheTTL}
	if opts.PutOptions.TTL != 0 {
		putOptions.TTL = opts.PutOptions.TTL
	}
	putOptions.StoreName = opts.PutOptions.StoreName

	// Start timer
	start := time.Now()
	stats := make(map[string]interface{})
	stats["cached"] = false
	stats["ok"] = false
	stats["cache_get_ok"] = true
	stats["cache_put_ok"] = true

	// Defer sending of stats
	defer func() {
		// Send stats, if enabled
		if opts.Stats != nil {
			stats["total_duration"] = float32(time.Since(start))
			doer.AddStats(stats)
			opts.Stats.SendStats(name, stats)
		}
	}()

	// Check cache, if enabled
	if opts.Cache != nil {
		cacheKey = opts.Cache.Key(name, doer.GetRequestBytes())

		cacheGetStart := time.Now()
		cacheGetResponse := cache.Get(opts.Cache, cachePrefix+cacheKey, cacheTimeout)
		stats["cache_get_duration"] = float32(time.Since(cacheGetStart))

		cached, err := cacheGetResponse.Data, cacheGetResponse.Error

		if err != nil {
			stats["cache_get_ok"] = false
			log.Error(fmt.Errorf("concierge.Request(): cache get error for '%s': %s", name, err.Error()))
		}
		if len(cached) > 0 {
			stats["ok"] = true
			stats["cached"] = true
			return cache.Decompress(cached), true, nil
		}
	}

	// Do request
	err = doer.Do(timeout)
	if err != nil {
		return []byte{}, false, fmt.Errorf("concierge.Request(): doer error for '%s': %s", name, err.Error())
	}

	// Check response
	if err := doer.ResponseError(); err != nil {
		return []byte{}, false, fmt.Errorf("concierge.Request(): response not OK for '%s': %s\n\nRequest:\n%s\nResponse:\n%s", name, err.Error(), string(doer.GetRequestBytes()), string(doer.GetResponseBytes()))
	}

	// Put cache, if enabled
	if opts.Cache != nil {
		cachePutStart := time.Now()
		cachePutResponse := cache.Put(opts.Cache, name, cachePrefix+cacheKey, cache.Compress(doer.GetResponseBytes()), putOptions, cacheTimeout)
		stats["cache_put_duration"] = float32(time.Since(cachePutStart))

		if cachePutResponse.Error != nil {
			stats["cache_put_ok"] = false
			log.Error(fmt.Errorf("concierge.Request(): cache put error for '%s': %s", name, cachePutResponse.Error.Error()))
		}
	}

	// Put store, if enabled
	if opts.Store != nil {
		storePutStart := time.Now()
		// FIXME: Should we use a different key when storing?
		storePutResponse := cache.Put(opts.Store, name, cachePrefix+cacheKey, doer.GetResponseBytes(), putOptions, storeTimeout)
		stats["store_put_duration"] = float32(time.Since(storePutStart))

		if storePutResponse.Error != nil {
			stats["store_put_ok"] = false
			log.Error(fmt.Errorf("concierge.Request(): store put error for '%s': %s", name, storePutResponse.Error.Error()))
		}
	}

	stats["ok"] = true
	return doer.GetResponseBytes(), false, nil
}
