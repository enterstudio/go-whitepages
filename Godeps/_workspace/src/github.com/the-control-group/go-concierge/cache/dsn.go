package cache

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// DSN is a data source name
type DSN struct {
	url *url.URL
	*DSNValues
}

// Parse DSN string and returns DSN instance
func Parse(dsn string) (*DSN, error) {
	parsed, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}
	d := DSN{
		parsed,
		&DSNValues{parsed.Query()},
	}
	return &d, nil
}

// ParseQuery parses the provided query and returns DSN values
func ParseQuery(query string) (*DSNValues, error) {
	parsed, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	return &DSNValues{parsed}, nil
}

// NewValues returns DSNValues from url.Values
func NewValues(query url.Values) (*DSNValues, error) {
	return &DSNValues{query}, nil
}

// Host returns the DSN URL host
func (d *DSN) Host() string {
	return d.url.Host
}

// Scheme returns the DSN URL scheme
func (d *DSN) Scheme() string {
	return d.url.Scheme
}

// Path returns the DSN URL path
func (d *DSN) Path() string {
	return d.url.Path
}

// User returns the DSN URL user
func (d *DSN) User() *url.Userinfo {
	return d.url.User
}

// DSNValues are URL values
type DSNValues struct {
	url.Values
}

// GetInt returns the integer value referenced by param or def if it's not set
func (d *DSNValues) GetInt(param string, def int) int {
	value := d.Get(param)
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}

	return def
}

// GetString returns the string value referenced by param or def if it's not set
func (d *DSNValues) GetString(param string, def string) string {
	value := d.Get(param)
	if value == "" {
		return def
	}

	return value
}

// GetBool returns the boolean value reference by param or def if it's not set
func (d *DSNValues) GetBool(param string, def bool) bool {
	value := strings.ToLower(d.Get(param))
	if value == "true" || value == "1" {
		return true
	} else if value == "0" || value == "false" {
		return false
	} else {
		return def
	}
}

// GetDuration returns the duration in seconds referenced by param or def if it's not set
func (d *DSNValues) GetDuration(param string, def time.Duration) time.Duration {
	if i, err := time.ParseDuration(d.Get(param)); err == nil {
		return i
	}

	return def
}
