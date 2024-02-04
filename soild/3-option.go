package soild

import "time"

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultCache = false
	defaultTime  = 10 * time.Second
)

type Option func(*Connection)

func WithCache(cache bool) Option {
	return func(c *Connection) {
		c.cache = cache
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Connection) {
		c.timeout = timeout
	}
}

func NewConnection(addr string, opts ...Option) *Connection {
	conn := &Connection{
		addr:    addr,
		cache:   defaultCache,
		timeout: defaultTime,
	}
	for _, opt := range opts {
		opt(conn)
	}
	return conn
}
