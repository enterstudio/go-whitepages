package cache

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"fmt"
	"time"
)

// Get the provided key from the cache
func Get(cacher Cacher, key string, timeout time.Duration) *Response {
	response := Response{OK: true, Name: cacher.GetName()}
	ch := make(chan Response, 1)
	start := time.Now()

	go func() {
		cached, err := cacher.Get(key)
		if err != nil {
			response.Error = err
			response.OK = false
		} else if len(cached) > 0 {
			response.Data = cached
		}
		response.Duration = time.Since(start)
		ch <- response
	}()

	// Wait for cache get or timeout
	select {
	case response = <-ch:
	case <-time.After(timeout):
		response.OK = false
		response.Timeout = true
	}
	return &response
}

// Put the provided key/value in the store
func Put(storer Storer, provider, key string, data []byte, options PutOptions, timeout time.Duration) *Response {
	response := Response{OK: true, Name: storer.GetName()}
	ch := make(chan Response, 1)
	start := time.Now()

	name := provider
	if options.StoreName != "" {
		name = options.StoreName
	}

	go func() {
		err := storer.Put(name, key, data, options)
		if err != nil {
			response.Error = err
			response.OK = false
		}
		response.Duration = time.Since(start)
		ch <- response
	}()

	// Wait for cache put or timeout
	select {
	case response = <-ch:
	case <-time.After(timeout):
		response.OK = false
		response.Timeout = true
	}

	return &response
}

// Key returns a md5 hash of the request and prepends the provided prefix.
func Key(prefix string, request []byte) string {
	return fmt.Sprintf("%s_%x", prefix, md5.Sum(request))
}

// Compress uses gzip to compress the provided data
func Compress(data []byte) []byte {
	buf := new(bytes.Buffer)
	writer := gzip.NewWriter(buf)
	writer.Write(data)
	writer.Close()

	return buf.Bytes()
}

// Decompress uses gzip to decompress the provided data
func Decompress(data []byte) []byte {
	in := bytes.NewBuffer(data)
	reader, _ := gzip.NewReader(in)

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	return buf.Bytes()
}
